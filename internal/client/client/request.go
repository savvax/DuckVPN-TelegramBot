package client

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

//result map[string]interface{}

func (c *Client) generalRequest(method string, endpoint string, bodyRequest io.Reader) (*http.Response, error) {
	//if err = c.checkToken(); err != nil {
	//	return
	//}

	baseURL, err := url.Parse(c.ApiURL)
	if err != nil {
		return nil, err
	}

	fullURL := baseURL.ResolveReference(&url.URL{Path: endpoint})

	req, err := http.NewRequest(method, fullURL.String(), bodyRequest)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Bearer "+c.Auth.AccessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
	//defer resp.Body.Close()

}

//err = unmarshalJSONResponse(resp.Body, &result)
//return

func unmarshalJSONResponse(r io.Reader, v interface{}) error {
	if r == nil {
		return errors.New("reader is nil")
	}

	decoder := json.NewDecoder(r)
	return decoder.Decode(v)
}
