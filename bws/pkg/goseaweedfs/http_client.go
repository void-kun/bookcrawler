package goseaweedfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"strings"
)

type httpClient struct {
	client  *http.Client
	authKey string
}

type downloadCB func(io.Reader) error

func newHttpClient(client *http.Client, opts ...HttpClientOption) *httpClient {
	c := &httpClient{
		client: client,
	}
	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *httpClient) Close() (err error) {
	return
}

func (c *httpClient) get(url string, header map[string]string) (body []byte, statusCode int, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err == nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
		if c.authKey != "" {
			req.Header.Set("Authorization", c.authKey)
		}

		var resp *http.Response
		resp, err = c.client.Do(req)
		if err == nil {
			body, statusCode, err = readAll(resp)
			// empty file check
			if IsFileMarkBytes(body, EmptyMark) {
				body = []byte{}
			}
		}

	}
	return
}

func (c *httpClient) delete(url string) (statusCode int, err error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return
	}

	// auth key
	if c.authKey != "" {
		req.Header.Set("Authorization", c.authKey)
	}

	r, err := c.client.Do(req)
	if err != nil {
		return
	}

	body, statusCode, err := readAll(r)
	if err == nil {
		switch r.StatusCode {
		case http.StatusNoContent, http.StatusNotFound, http.StatusAccepted, http.StatusOK:
			err = nil
			return
		}

		m := make(map[string]interface{})
		if e := json.Unmarshal(body, &m); e == nil {
			if s, ok := m["error"].(string); ok {
				err = fmt.Errorf("Delete %s: %v", url, s)
				return
			}
		}
		err = fmt.Errorf("Delete %s. Got response but can not parse. Body:%s Code:%d", url, string(body), r.StatusCode)
	}
	return
}

func (c *httpClient) download(url string, callback downloadCB) (filename string, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	if c.authKey != "" {
		req.Header.Set("Authorization", c.authKey)
	}

	r, err := c.client.Do(req)
	if err == nil {
		if r.StatusCode != http.StatusOK {
			drainAndClose(r.Body)
			err = fmt.Errorf("Download %s but error. Status:%s", url, r.Status)
			return
		}

		contentDisposition := r.Header["Content-Disposition"]
		if len(contentDisposition) > 0 {
			if strings.HasPrefix(contentDisposition[0], "filename=") {
				filename = contentDisposition[0][len("filename="):]
				filename = strings.Trim(filename, "\"")
			}
		}

		data, err := io.ReadAll(r.Body)
		if err != nil {
			return filename, err
		}

		var readWriter io.ReadWriter
		if IsFileMarkBytes(data, EmptyMark) {
			readWriter = bytes.NewBuffer([]byte{})
		} else {
			readWriter = bytes.NewBuffer(data)
		}

		// execute callback
		_ = callback(readWriter)

		// drain and close body
		drainAndClose(r.Body)
	}

	return
}

func (c *httpClient) upload(url string, filename string, fileReader io.Reader, mtype string) (respBody []byte, statusCode int, err error) {
	r, w := io.Pipe()
	mw := multipart.NewWriter(w)
	result := make(chan error, 1)

	go func() {
		h := make(textproto.MIMEHeader)
		h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="file"; filename="%s"`, normalizeName(filename)))
		if mtype == "" {
			mtype = mime.TypeByExtension(strings.ToLower(filepath.Ext(filename)))
		}
		if mtype != "" {
			h.Set("Content-Type", mtype)
		}

		part, err := mw.CreatePart(h)
		if err == nil {
			_, err = io.Copy(part, fileReader)
		}

		if err == nil {
			if err = mw.Close(); err == nil {
				err = w.Close()
			} else {
				_ = w.Close()
			}
		} else {
			_ = mw.Close()
			_ = w.Close()
		}

		result <- err
	}()

	// request
	req, err := http.NewRequest(http.MethodPost, url, r)
	if err != nil {
		return nil, 0, err
	}

	// request headers
	req.Header.Set("Content-Type", mw.FormDataContentType())
	if c.authKey != "" {
		req.Header.Set("Authorization", c.authKey)
	}

	// perform request
	res, err := c.client.Do(req)

	_ = r.Close()

	if err == nil {
		if respBody, statusCode, err = readAll(res); err == nil {
			err = <-result
		}
	}
	return
}
