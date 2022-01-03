package main

import "fmt"

func main() {
	//studentsAge := map[string]int{
	//	"john": 32,
	//	"tom":  31,
	//}

	studentsAge := make(map[string]int)
	//var studentsAge map[string]int //panic: assignment to entry in nil map
	studentsAge["john"] = 31
	studentsAge["tom"] = 32

	fmt.Println(studentsAge)

	age, exist := studentsAge["larry"]
	if exist {
		fmt.Println("Age of larry is", age)
	} else {
		fmt.Println("Name does not exist")
	}

	for name, age := range studentsAge {
		fmt.Printf("%s\t%d\n", name, age)
	}
}
