package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadFileHandlerIntegration_Success(t *testing.T) {
	serverURL := "http://localhost:8080"

	fileContents := []byte("Sample file contents")
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("filename", "sample.txt")
	part.Write(fileContents)
	writer.Close()

	req, err := http.NewRequest("POST", serverURL+"/upload", body)
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
	if response.Message == "File uploaded and processed successfully" {
		t.Logf("Expected: %s, Got: %s", "File to be uploaded successfully", response.Message)
	} else {
		t.Errorf("Expected: %s, Got: %s", "File to be uploaded successfully", response.Message)
	}
}

func TestUploadFileHandlerIntegration_EmptyFile(t *testing.T) {
	serverURL := "http://localhost:8080"

	fileContents := []byte("")
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("filename", "sample.txt")
	part.Write(fileContents)
	writer.Close()

	req, err := http.NewRequest("POST", serverURL+"/upload", body)
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

	//assert.Equal(t, "File is empty", response.Message)
	if (response.Message) == "File is empty" {
		t.Logf("Expected: %s, Got: %s", "File is empty", response.Message)
	} else {
		t.Errorf("Expected: %s, Got: %s", "File is empty", response.Message)
	}
}

func TestUploadFileHandlerIntegration_FileTypeValidation(t *testing.T) {
	serverURL := "http://localhost:8080"

	fileContents := []byte("Sample file data")
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("filename", "sample.xyz")
	part.Write(fileContents)
	writer.Close()

	req, err := http.NewRequest("POST", serverURL+"/upload", body)
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
	if (response.Message) == "Unsupported file type" {
		t.Logf("Expected: %s, Got: %s", "Unsupported file type", response.Message)
	} else {
		t.Errorf("Expected: %s, Got: %s", "Unsupported file type", response.Message)
	}
}

func TestGetFileSizeHandlerIntegration_Success(t *testing.T) {
	serverURL := "http://localhost:8080"

	filename := "sample.txt"

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/file/%s", serverURL, filename), nil)
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
	assert.Equal(t, int64(20), response.FileSize)

	if (response.FileName) == "sample.txt" && (response.FileSize) == int64(20) {
		t.Logf("Expected - FileName: %s & FileSize %d, Got - FileName: %s & FileSize %d", "sample.txt", int64(20), response.FileName, response.FileSize)
	} else {
		t.Errorf("Expected - FileName: %s & FileSize %d, Got - FileName: %s & FileSize %d", "sample.txt", int64(20), response.FileName, response.FileSize)
	}
}

func TestGetFileSizeHandlerIntegration_Failure(t *testing.T) {
	serverURL := "http://localhost:8080"

	filename := "sample.xyz"

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/file/%s", serverURL, filename), nil)
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
	assert.Equal(t, int64(0), response.FileSize)

	if (response.FileName) == "sample.xyz" && (response.FileSize) == int64(0) {
		t.Logf("Expected - FileName: %s & FileSize %d, Got - FileName: %s & FileSize %d", "sample.xyz", int64(0), response.FileName, response.FileSize)
	} else {
		t.Errorf("Expected - FileName: %s & FileSize %d, Got - FileName: %s & FileSize %d", "sample.xyz", int64(0), response.FileName, response.FileSize)
	}
}
