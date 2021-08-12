package common

import (
	"archive/tar"
	"bufio"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var FILE_NAME = "archive.tar"

// Download the config from the airbyte API
func DownloadFile() (err error) {
	var DOWNLOAD_CONFIG = "/api/v1/deployment/export"

	url := GetFullApiURL(DOWNLOAD_CONFIG)
	dirname, err := os.UserHomeDir()

	// Create the file
	out, err := os.Create(dirname + "/.airbyte/" + FILE_NAME)
	if err != nil {
		return err
	}
	defer out.Close()

	// Post the data
	body, _ := json.Marshal(map[string]string{})
	requestBody := bytes.NewBuffer(body)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	req, _ := http.NewRequest("POST", url, requestBody)

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// Untar takes a destination path and a reader; a tar reader loops over the tarfile
// creating the file structure at 'dst' along the way, and writing any files
func Untar() error {
	dirname, _ := os.UserHomeDir()

	dst := dirname + "/.airbyte/"

	file, _ := os.Open(dirname + "/.airbyte/archive.tar")
	r := bufio.NewReader(file)

	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return nil

		// return any other error
		case err != nil:
			return err

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		target := filepath.Join(dst, header.Name)

		// Take only airbyte_config folder
		if strings.Contains(header.Name, "airbyte_config") {
			if err := os.MkdirAll(dirname+"/.airbyte/airbyte_config", 0755); err != nil {
				return err
			}

			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}
}

func CopyConfigToTarget(targetFolder string) error {
	files := [3]string{"SOURCE_CONNECTION.yaml", "STANDARD_SYNC.yaml", "DESTINATION_CONNECTION.yaml"}

	dirname, _ := os.UserHomeDir()

	dst := dirname + "/.airbyte/airbyte_config/"

	for _, file := range files {
		err := CopyFile(dst+file, targetFolder+"/"+file)
		if err != nil {
			return err
		}
	}

	return nil
}

func CopyFile(from string, to string) error {
	sourceFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create new file
	newFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer newFile.Close()

	_, err = io.Copy(newFile, sourceFile)

	if err != nil {
		return err
	}

	return nil
}

func CleanUp() error {
	dirname, _ := os.UserHomeDir()

	dst := dirname + "/.airbyte/"

	e := os.Remove(dst + FILE_NAME)
	if e != nil {
		return e
	}

	e = os.RemoveAll(dst + "airbyte_config")
	if e != nil {
		return e
	}

	return nil
}
