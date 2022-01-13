package fuzzer

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	sys "golang.org/x/sys/unix"
)

var (
	bpOffset = [3]uint64{0x30d28, 0x30eac, 0x30ebb}
)

var remoteIOV = []sys.RemoteIovec{}
var localIOV = []sys.Iovec{}

var backupData = make([][]byte, 0)
var backupRegs sys.PtraceRegs

var first, second, third int

var timeRestore int64

func Debug(s string, args []string, data []byte) {
	runtime.LockOSThread()
	defer trace("Debug")()
	var pid int
	var entrypoint uint64
	var bpAddr = make([]uintptr, 3)
	var bpData = make([][]byte, 3)
	//var stdout bytes.Buffer
	var err error
	//var stderr bytes.Buffer

	// open process
	cmd := exec.Command(s, args...)
	//cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &sys.SysProcAttr{Ptrace: true}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	pid = cmd.Process.Pid

	// get rip
	var regs sys.PtraceRegs
	if err := sys.PtraceGetRegs(pid, &regs); err != nil {
		log.Panic(err)
	}
	entrypoint = regs.Rip
	log.Printf("Value of rip: %x\n", regs.Rip)

	// set breakpoint1 before loop
	bpAddr[0] = uintptr(entrypoint + bpOffset[0])
	bpData[0], err = SetBreakPoint(pid, bpAddr[0])
	log.Printf("Set breakpoint at %x\n", bpAddr)
	if err != nil {
		log.Panic(err)
	}

	// Set breakpoint before end
	bpAddr[1] = uintptr(entrypoint + bpOffset[1])
	_, err = SetBreakPoint(pid, bpAddr[1])
	if err != nil {
		log.Panic(err)
	}

	bpAddr[2] = uintptr(entrypoint + bpOffset[2])
	_, err = SetBreakPoint(pid, bpAddr[2])
	if err != nil {
		log.Panic(err)
	}

	// Continue
	if err := Continue(pid); err != nil {
		log.Panic(err)
	}

	log.Println("Hit bp1")

	// Cancel breakpoint
	if err := CancelBreakPoint(pid, bpAddr[0], bpData[0]); err != nil {
		log.Panic(err)
	}

	// Backup
	getMemSecs(pid)
	if err := Backup(pid); err != nil {
		log.Panic(err)
	}
	fmt.Printf("My pid is %d, child pid is %d\n", os.Getpid(), pid)

	for i := 0; i < 1000; i++ {
		mutateData := getData(data, method)
		// Change data
		dataAddr := uintptr(backupRegs.Rax)
		//if _, err := sys.PtracePokeData(pid, dataAddr, mutateData); err != nil {
		//	log.Panic("Change data: ", err)
		//}
		sys.ProcessVMWritev(pid, []sys.Iovec{{&mutateData[0], uint64(len(mutateData))}},
			[]sys.RemoteIovec{{dataAddr, len(mutateData)}}, 0)

		// Continue
		if err = Continue(pid); err != nil {
			log.Panic(err)
		}
		if err = Continue(pid); err != nil {
			log.Panic(err)
		}

		//if err := errorParse(stdout); err != nil {
		//	log.Panic(err)
		//}
		//if first%1000 == 0 {
		//	fmt.Printf("1:%7d 2:%7d 3:%7d\n", first, second, third)
		//}
		// Restore
		Restore(pid)
		//log.Println("After restore: " + checkRIP(pid))
	}
	log.Println(timeRestore)
}

func SetBreakPoint(pid int, addr uintptr) ([]byte, error) {
	defer trace("SetBreakPoint")()
	data := make([]byte, 1)
	if _, err := sys.PtracePeekData(pid, addr, data); err != nil {
		return nil, errors.New("SetBreakPOint: peek data: " + err.Error())
	}
	if _, err := sys.PtracePokeData(pid, addr, []byte{0xcc}); err != nil {
		return nil, errors.New("SetBreakPOint: poke data: " + err.Error())
	}
	return data, nil
}

func CancelBreakPoint(pid int, addr uintptr, data []byte) error {
	if _, err := sys.PtracePokeData(pid, addr, data); err != nil {
		return errors.New("CancelBreakPOint: poke data: " + err.Error())
	}

	sys.PtracePeekData(pid, addr, data)
	log.Printf("Cancel bp at %x, data is %x\n", addr, data)

	var regs sys.PtraceRegs
	if err := sys.PtraceGetRegs(pid, &regs); err != nil {
		return err
	}

	regs.Rip = uint64(addr)

	if err := sys.PtraceSetRegs(pid, &regs); err != nil {
		return err
	}

	return nil
}

func Continue(pid int) error {
	log.Println("Before continue: " + checkRIP(pid))
	var ws sys.WaitStatus
	if err := sys.PtraceCont(pid, 0); err != nil {
		return err
	}
	if _, err := sys.Wait4(pid, &ws, 0, nil); err != nil {
		return err
	}
	log.Println(ws.StopSignal())
	log.Println("After continue: " + checkRIP(pid))
	return nil
}

func Backup(pid int) error {
	defer trace("Backup")()

	if _, err := sys.ProcessVMReadv(pid, localIOV, remoteIOV, 0); err != nil {
		return errors.New("Backup: vm read: " + err.Error())
	}

	if err := sys.PtraceGetRegs(pid, &backupRegs); err != nil {
		return errors.New("Backup: get regs: " + err.Error())
	}

	return nil
}

func Restore(pid int) error {
	//defer trace("Restore")()
	if _, err := sys.ProcessVMWritev(pid, localIOV, remoteIOV, 0); err != nil {
		return errors.New("Restore: vm write: " + err.Error())
	}

	if err := sys.PtraceSetRegs(pid, &backupRegs); err != nil {
		return errors.New("Restore: set regs: " + err.Error())
	}
	return nil
}

func getMemSecs(pid int) error {
	defer trace("getMemSecs")()
	var output bytes.Buffer
	cmd := exec.Command("cat", "/proc/"+strconv.Itoa(pid)+"/maps", "|", "grep", "rw")
	cmd.Stdout = &output
	cmd.Run()

	reg := regexp.MustCompile(`(\S+)-(\S+) rw`)
	result := reg.FindAllStringSubmatch(output.String(), -1)
	for idx, item := range result {
		begin, _ := strconv.ParseInt(item[1], 16, 0)
		end, _ := strconv.ParseInt(item[2], 16, 0)

		len := int(end - begin)
		rIov := sys.RemoteIovec{uintptr(begin), len}
		remoteIOV = append(remoteIOV, rIov)

		data := make([]byte, len)
		backupData = append(backupData, data)
		iov := sys.Iovec{&backupData[idx][0], uint64(len)}
		localIOV = append(localIOV, iov)
	}

	return nil
}

func trace(name string) func() {
	start := time.Now()
	return func() { duration := time.Since(start).Milliseconds(); log.Printf("%s using %d\n", name, duration) }
}

func errorParse(buf bytes.Buffer) error {
	//defer trace("errorParse")()
	content := make([]byte, 18)
	if n, err := buf.Read(content); err != nil {
		log.Printf("Buffer read %d\n", n)
		return errors.New("errorParse: buf read: " + err.Error())
	}
	log.Println("test: " + string(content) + "\n")
	if strings.Contains(string(content), "3") {
		third++
	} else if strings.Contains(string(content), "2") {
		second++
	} else if strings.Contains(string(content), "1") {
		first++
	}
	if strings.Contains(string(content), "succeed") {
		log.Println(string(content))
	}
	buf.Reset()
	return nil
}

func checkRIP(pid int) string {
	var regs sys.PtraceRegs
	if err := sys.PtraceGetRegs(pid, &regs); err != nil {
		return ""
	}

	data := make([]byte, 1)
	if _, err := sys.PtracePeekData(pid, uintptr(regs.Rip), data); err != nil {
		return ""
	}

	return fmt.Sprintf("RIP is: %#x, data at here is %#x", regs.Rip, data)
}
