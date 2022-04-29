package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/kujilabo/cocotola-tatoeba-api/src/app/config"
)

var timeoutImportMin = 30

func main() {
	cfg, err := config.LoadConfig("local")
	if err != nil {
		panic(err)
	}

	url := "http://localhost:8280/v1/admin/sentence/import"
	fieldname := "file"
	filename := "eng_sentences_detailed.tsv"

	file, err := os.Open("../cocotola-data/datasource/tatoeba/" + filename)
	if err != nil {
		panic(err)
	}

	body := bytes.Buffer{}

	mw := multipart.NewWriter(&body)

	fw, err := mw.CreateFormFile(fieldname, filename)
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(fw, file); err != nil {
		panic(err)
	}

	if err = mw.Close(); err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, &body)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(cfg.Auth.Username, cfg.Auth.Password)
	req.Header.Set("Content-Type", mw.FormDataContentType())

	client := http.Client{
		Timeout: time.Duration(timeoutImportMin) * time.Minute,
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("status: %d\n", resp.StatusCode)
	fmt.Printf("body: %s\n", string(respBody))
}
