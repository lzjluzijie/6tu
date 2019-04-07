package onedrive

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type UploadResponse struct {
	ID string
}

type CreateSessionResponse struct {
	UploadUrl string
}

func Upload(size int64, r io.Reader) (id string, err error) {
	url := fmt.Sprintf("https://graph.microsoft.com/v1.0/me/drive/root:/yitu/%s/%d:/createUploadSession", date, rand.Uint64())

	req, err := NewRequest("POST", url, nil)
	if err != nil {
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Println(string(data))

	createSessionResponse := &CreateSessionResponse{}
	err = json.Unmarshal(data, createSessionResponse)
	if err != nil {
		return
	}

	uploadURL := createSessionResponse.UploadUrl

	req, err = NewRequest("PUT", uploadURL, r)
	if err != nil {
		return
	}

	req.Header.Add("Content-Length", fmt.Sprintf("%d", size))
	req.Header.Add("Content-Range", fmt.Sprintf("bytes 0-%d/%d", size-1, size))
	req.ContentLength = size

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	//fmt.Println(string(data))

	uploadResponse := &UploadResponse{}
	err = json.Unmarshal(data, uploadResponse)
	if err != nil {
		return
	}

	id = uploadResponse.ID
	return
}
