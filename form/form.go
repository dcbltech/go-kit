package form

import (
	"fmt"
	"io"
	"net/http"
)

func GetFileData(r *http.Request, field string) (filename string, data []byte, err error) {
	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		return "", nil, fmt.Errorf("kit/form: failed to parse multipart form: %w", err)
	}

	file, header, err := r.FormFile(field)
	if err != nil {
		return "", nil, fmt.Errorf("kit/form: failed to retrieve file from form field '%s': %w", field, err)
	}

	filename = header.Filename

	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("kit/form: warning: failed to close file: %v\n", closeErr)
		}
	}()

	data, err = io.ReadAll(file)
	if err != nil {
		return "", nil, fmt.Errorf("kit/form: failed to read file data: %w", err)
	}

	return
}
