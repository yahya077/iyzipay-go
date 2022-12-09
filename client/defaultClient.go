package client

import (
	"bytes"
	"encoding/json"
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

func (c DefaultClient) Post(url string) (*http.Response, error) {
	data, err := json.Marshal(c.RequestSchema)

	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	if err != nil {
		return nil, err
	}

	r.Header = c.header

	client := &http.Client{}

	res, err := client.Do(r)

	if err != nil {
		return res, err
	}

	return res, nil
}
