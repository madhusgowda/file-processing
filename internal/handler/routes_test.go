package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestUploadFileHandlerIntegration_Success(t *testing.T) {
	r := mux.NewRouter()

	server := httptest.NewServer(r)
	defer server.Close()

	server.URL = "http://localhost:8080"

	fileContents := []byte("Sample file contents")
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("filename", "sample.txt")
	part.Write(fileContents)
	writer.Close()

	req, err := http.NewRequest("POST", server.URL+"/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response struct {
		Message string `json:"message"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "File uploaded and processed successfully", response.Message)
}

func TestUploadFileHandlerIntegration_EmptyFile(t *testing.T) {
	r := mux.NewRouter()

	server := httptest.NewServer(r)
	defer server.Close()

	server.URL = "http://localhost:8080"

	fileContents := []byte("")
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("filename", "sample.txt")
	part.Write(fileContents)
	writer.Close()

	req, err := http.NewRequest("POST", server.URL+"/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var response struct {
		Message string `json:"message"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "File is empty", response.Message)
}

func TestUploadFileHandlerIntegration_FileTypeValidation(t *testing.T) {
	r := mux.NewRouter()

	server := httptest.NewServer(r)
	defer server.Close()

	server.URL = "http://localhost:8080"

	fileContents := []byte("Sample file data")
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("filename", "sample.xyz")
	part.Write(fileContents)
	writer.Close()

	req, err := http.NewRequest("POST", server.URL+"/upload", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	var response struct {
		Message string `json:"message"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Unsupported file type", response.Message)
}

func TestGetFileSizeHandlerIntegration_Success(t *testing.T) {

	r := mux.NewRouter()

	server := httptest.NewServer(r)
	defer server.Close()

	server.URL = "http://localhost:8080"

	filename := "sample.txt"

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/file/%s", server.URL, filename), nil)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var response struct {
		FileName string `json:"filename"`
		FileSize int64  `json:"file_size"`
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, filename, response.FileName)
	assert.Equal(t, int64(121), response.FileSize)
}
