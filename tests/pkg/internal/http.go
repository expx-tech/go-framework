package system

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func GetHttp(url string, result *http.Response) error {
	var err error
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	*result = *resp

	return nil
}

func PostHttp(url string, contentType string, body []byte, headers map[string]string, result *http.Response) error {
	var err error
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", contentType)
	for hn, hv := range headers {
		req.Header.Add(hn, hv)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	*result = *resp

	return nil
}

func DeleteHttp(url string, headers map[string]string, result *http.Response) error {
	var err error
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	for hn, hv := range headers {
		req.Header.Add(hn, hv)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	*result = *resp

	return nil
}

func HttpResponseBody(resp *http.Response, result *[]byte) (err error) {
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		if errBody := resp.Body.Close(); errBody != nil {
			errors.Join(err, errBody)
		}
	}()
	if err != nil {
		return
	}
	*result = body

	return
}
