package client

import "net/http"

type HttpClient struct {
	client http.Client
}

func (c *HttpClient) Get(url string, header, form map[string]string) {
	// the get request needn't request body
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	// set request header
	for headerItemKey, headerItemValue := range header {
		r.Header.Set(headerItemKey, headerItemValue)
	}
	// for query param
	for formItemKey, formItemValue := range form {
		r.Form.Set(formItemKey, formItemValue)
	}
}

func Post() {}
