package form

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFileData(t *testing.T) {
	fileContent := []byte("sample file content")
	fileName := "testfile.txt"
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		t.Fatalf("failed to create form file: %v", err)
	}

	_, err = part.Write(fileContent)
	if err != nil {
		t.Fatalf("failed to write file content: %v", err)
	}

	writer.Close()

	r := httptest.NewRequest(http.MethodPost, "/upload", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())

	filename, data, err := GetFileData(r, "file")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if filename != fileName {
		t.Errorf("expected filename %q, got %q", fileName, filename)
	}

	if !bytes.Equal(data, fileContent) {
		t.Errorf("expected file content %q, got %q", fileContent, data)
	}
}
