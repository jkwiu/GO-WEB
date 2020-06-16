package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := `C:\Users\sizzf\Downloads\b9e00092a5baebe87f42ec04214a907c.jpg`
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")

	// mime를 사용하기 위해,,
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	// filepath.Base(path) 파일 경로에서 파일만 잘라낸다.
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)

	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	// file info를 확인
	_, err = os.Stat(uploadFilePath)
	assert.NoError(err)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	// 두 파일이 같은지 확인
	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)
}
