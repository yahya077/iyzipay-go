package client

import (
	"fmt"
	"github.com/yahya077/iyzipay-go/builder"
	"net/http"
)

type IyzicoClientOptions struct {
	BaseUrl   string
	APIKey    string
	APISecret string
}

type IyzicoClient struct {
	DefaultClient DefaultClient
	Options       IyzicoClientOptions
}

func NewIyzicoClient(options IyzicoClientOptions) IyzicoClient {
	return IyzicoClient{
		DefaultClient: NewDefaultClient(),
		Options:       options,
	}
}

func (c *IyzicoClient) Post(request interface{}, endPoint string) (*http.Response, error) {
	return c.DefaultClient.SetHeader(c.Options.buildHeader(request)).SetRequestSchema(request).Post(c.Options.BaseUrl + endPoint)
}

func (c *IyzicoClientOptions) buildAuthorization(pki string) string {
	return fmt.Sprintf("IYZWS %s:%s", c.APIKey, pki)
}

func (c *IyzicoClientOptions) buildHeader(request interface{}) http.Header {
	h := http.Header{}

	randomString, pki := builder.NewRequestStringBuilder().BuildPki(c.APIKey, c.APISecret, request)

	h.Add("Content-Type", "application/json")
	h.Add("Accept", "application/json")
	h.Add("Authorization", c.buildAuthorization(pki))
	h.Add("x-iyzi-rnd", randomString)
	h.Add("Cache-Control", "no-cache")

	return h
}
