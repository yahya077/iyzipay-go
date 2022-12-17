package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DefaultClient struct {
	RequestSchema interface{}
	header        http.Header
}

func NewDefaultClient() DefaultClient {
	return DefaultClient{}
}

func (c DefaultClient) SetHeader(header http.Header) DefaultClient {
	c.header = header
	return c
}

func (c DefaultClient) SetRequestSchema(v interface{}) DefaultClient {
	c.RequestSchema = v
	return c
}

func (c DefaultClient) method(url string, method string) (*http.Response, error) {
	data, err := json.Marshal(c.RequestSchema)

	if err != nil {
		fmt.Println("cannot unmarshall request schema")
	}

	r, err := http.NewRequest(method, url, bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("cannot create new request")
	}

	r.Header = c.header

	client := &http.Client{}

	res, err := client.Do(r)

	if err != nil {
		fmt.Println("request error: \n", err)
	}

	return res, nil
}

func (c DefaultClient) Post(url string) (*http.Response, error) {
	return c.method(url, http.MethodPost)
}

func (c DefaultClient) Delete(url string) (*http.Response, error) {
	return c.method(url, http.MethodDelete)
}
