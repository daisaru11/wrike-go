package wrike

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

func (client *Client) requestGet(path string, result interface{}) error {
	httpReq, err := client.createHTTPRequest("GET", path, nil)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to build HTTP Request: %s", path))
	}

	// TODO: retry requests
	resp, err := client.httpClient.Do(httpReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failure on GET request: %s", path))
	}
	defer resp.Body.Close()

	err = client.handleHTTPResponse(resp, result)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to handle resonse: %s", path))
	}

	return nil
}

func (client *Client) requestPost(path string, request, result interface{}) error {
	httpReq, err := client.createHTTPRequest("POST", path, request)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to build HTTP Request: %s", path))
	}

	resp, err := client.httpClient.Do(httpReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failure on POST request: %s", path))
	}
	defer resp.Body.Close()

	err = client.handleHTTPResponse(resp, result)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to handle resonse: %s", path))
	}

	return nil
}

func (client *Client) requestPut(path string, request, result interface{}) error {
	httpReq, err := client.createHTTPRequest("PUT", path, request)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to build HTTP Request: %s", path))
	}

	resp, err := client.httpClient.Do(httpReq)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failure on PUT request: %s", path))
	}
	defer resp.Body.Close()

	err = client.handleHTTPResponse(resp, result)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to handle resonse: %s", path))
	}

	return nil
}

func (client *Client) createHTTPRequest(method, path string, body interface{}) (*http.Request, error) {
	var bodyReader io.Reader

	if method != "GET" && body != nil {
		json, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(json)
	}

	url := client.createURL(path)
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("bearer %s", client.oauth2Token))

	if bodyReader != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	return req, nil
}

func (client *Client) createURL(path string) string {
	return fmt.Sprintf("%s/%s",
		strings.TrimRight(client.baseURL, "/"),
		strings.TrimLeft(path, "/"),
	)
}

func (client *Client) handleHTTPResponse(resp *http.Response, result interface{}) error {
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("Got error response %s: %s", resp.Status, body)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	return nil
}
