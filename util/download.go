package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Download downloads an audio file from the given URL
func Download(url string, filename string) error {

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if strings.Contains(res.Header.Get("Content-Type"), "audio/") {
		return fmt.Errorf("url does not point to an audio file")
	}
	defer res.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}

	return nil
}
