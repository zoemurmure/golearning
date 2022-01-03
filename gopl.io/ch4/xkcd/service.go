package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func EstablishDatabase(n int) error {
	file, err := os.OpenFile("./database.json", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	defer func() {
		file.Close()
	}()

	if err != nil {
		return err
	}
	var database = []commicSearchResult{}
	for i := 1; i <= n; i++ {
		var data = commicSearchResult{}
		url := queryURL + strconv.Itoa(i) + "/info.0.json"
		resp, err := http.Get(url)
		defer func() {
			resp.Body.Close()
		}()
		if err != nil {
			return err
		}
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("EstablishDatabase: query %s failed: %d", url, resp.StatusCode)
		}

		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}

		database = append(database, data)
	}

	result, err := json.Marshal(database)
	if err != nil {
		return err
	}
	if _, err := file.Write(result); err != nil {
		return err
	}
	return nil
}

func Query(value string) (string, error) {
	file, err := os.OpenFile("./database.json", os.O_RDONLY, 0644)
	defer func() {
		file.Close()
	}()
	if err != nil {
		return "", err
	}

	var database = []commicSearchResult{}
	if err := json.NewDecoder(file).Decode(&database); err != nil {
		return "", err
	}

	for _, item := range database {
		if value == item.Title {
			return item.Img, nil
		}
	}
	return "", fmt.Errorf("%s does not exist", value)
}
