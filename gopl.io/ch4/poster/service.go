package poster

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Query(title string) error {
	url := queryURL + title
	resp, err := http.Get(url)
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("query %s failed: %d", title, resp.StatusCode)
	}
	var data = Movie{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	return downloadPoster(title, data.Poster)
}

func downloadPoster(title, url string) error {
	resp, err := http.Get(url)
	defer func() {
		resp.Body.Close()
	}()
	if err != nil {
		return err
	}

	file, err := os.Create(title + ".jpg")
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	return err
}
