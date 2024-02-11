package wails

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	TypeDir  = "fs/dir"
	TypeHtml = "text/html"
	TypeZip  = "application/zip"
)

func bundleType(path string) (filetype string, err error) {
	info, err := os.Stat(path)
	if err != nil {
		return
	}

	switch {
	case info.IsDir():
		filetype = TypeDir

	case info.Mode().IsRegular():
		if filetype, err = fileType(path); err == nil {
			switch filetype {
			case TypeHtml:
			case TypeZip:
			default:
				err = errors.New("unsupported file type: " + filetype)
			}
		}

	default:
		err = errors.New("app must be a file or a directory")
	}
	return
}

func fileType(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", err
	}

	contentType := strings.Split(http.DetectContentType(buffer), ";")[0]

	return contentType, nil
}
