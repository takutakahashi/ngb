package load

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/takutakahashi/ngb/pkg/types/config"
	"github.com/takutakahashi/ngb/pkg/types/result"
)

func Start(u *url.URL, c, n int, scripts config.ScriptConfig) (result.BenchmarkResult, error) {
	br, _ := result.NewBenchmarkResult(c, n)
	wg := sync.WaitGroup{}
	wg.Add(c)
	for i := 0; i < c; i++ {
		go func() {
			arr := benchmark(u, n, scripts)
			for _, rtime := range arr {
				br.AddResultResponseTime(rtime)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return br, nil
}
func benchmark(u *url.URL, n int, scripts config.ScriptConfig) []int64 {
	arr := []int64{}
	for i := 0; i < n; i++ {
		rr, _ := Request(u, scripts)
		arr = append(arr, rr.ResponseTime)
	}
	return arr
}

// Request is per benchmark
func Request(u *url.URL, scripts config.ScriptConfig) (result.RequestResult, error) {
	var rc config.RequestConfig
	rc, err := scripts.PreRequest.Execute()
	if err != nil {
		rc = config.RequestConfig{}
	}
	lr, err := request(u, rc)
	if err != nil {
		return lr, err
	}
	_, err = scripts.PostRequest.Execute()

	return lr, err
}

// per request
func request(u *url.URL, rc config.RequestConfig) (result.RequestResult, error) {
	client := &http.Client{}
	req, err := rc.BuildRequest(u.String())
	if err != nil {
		return result.RequestResult{}, errors.Wrap(err, "Build Request error")
	}
	// TODO: use httptrace
	before := time.Now()
	res, err := client.Do(req)
	after := time.Now()
	delta := after.Sub(before)
	if err != nil {
		return result.RequestResult{}, errors.Wrap(err, "Request error")
	}
	defer res.Body.Close()
	rr, err := result.NewRequestResult()
	rr.ResponseTime = delta.Milliseconds()
	rr.StatusCode = res.StatusCode
	return rr, err
}
