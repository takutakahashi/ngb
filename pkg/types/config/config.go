package config

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

type ScriptConfig struct {
	PreRequest  Script `yaml: prerequest`
	PostRequest Script `yaml: postrequest`
}

type RequestConfig struct {
	URL     url.URL
	Headers map[string]string   `json:"headers"`
	Cookies []map[string]string `json:"cookies"`
	Params  map[string]string   `json:"params"`
}

func (rc RequestConfig) buildBody() (string, error) {
	v := url.Values{}
	for key, val := range rc.Params {
		v.Add(key, val)
	}
	return v.Encode(), nil
}
func (rc RequestConfig) buildValues() (url.Values, error) {
	v := url.Values{}
	for key, val := range rc.Params {
		v.Add(key, val)
	}
	return v, nil
}

func (rc RequestConfig) BuildRequest(u string) (*http.Request, error) {
	cookies, err := rc.buildCookie()
	header, err := rc.buildHeader()
	values, err := rc.buildValues()
	req, err := http.NewRequest(
		"POST",
		u,
		strings.NewReader(values.Encode()),
	)
	req.Header = header
	for _, c := range cookies {
		req.AddCookie(c)
	}
	return req, err
}

func (rc RequestConfig) buildHeader() (http.Header, error) {
	header := http.Header{}
	for k, v := range rc.Headers {
		header.Add(k, v)
	}
	header.Add("Content-Type", "application/x-www-form-urlencoded")
	return header, nil

}
func (rc RequestConfig) buildCookie() ([]*http.Cookie, error) {
	cookies := []*http.Cookie{}
	for _, cmap := range rc.Cookies {
		key, ok := cmap["key"]
		if !ok {
			return cookies, errors.New("cookie error: no key")
		}
		value, ok := cmap["value"]
		if !ok {
			return cookies, errors.New("cookie error: no value")
		}
		c := &http.Cookie{
			Name:       key,
			Value:      value,
			Path:       cmap["path"],
			Domain:     cmap["domain"],
			RawExpires: cmap["expires"],
		}
		cookies = append(cookies, c)
	}
	return cookies, nil
}

func loadConfig(input []byte) (RequestConfig, error) {
	rq := RequestConfig{}
	err := json.Unmarshal(input, &rq)
	return rq, err
}

type Script struct {
	Path string
}

func (s Script) Execute() (RequestConfig, error) {
	if s.Path == "" {
		return RequestConfig{}, nil
	}
	out, err := exec.Command(s.Path).Output()
	if err != nil {
		return RequestConfig{}, err
	}
	return loadConfig(out)
}
