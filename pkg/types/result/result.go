package result

import (
	"errors"
	"log"
	"net/http"
	"sort"
)

type RequestResult struct {
	ResponseTime int64
	StatusCode   int
	RawResponse  *http.Response
	RawRequest   *http.Request
}

type RawResults []RequestResult

type BenchmarkResult struct {
	MaxCount      int
	RawResults    []RequestResult
	responseTimes []int64
	ResultCh      chan RequestResult
	FinishCh      chan struct{}
}

func NewBenchmarkResult(c, n int) (BenchmarkResult, error) {
	return BenchmarkResult{
		MaxCount:      c * n,
		RawResults:    []RequestResult{},
		responseTimes: []int64{},
		ResultCh:      make(chan RequestResult),
	}, nil

}

func NewRequestResult() (RequestResult, error) {
	return RequestResult{}, nil
}

func (br *BenchmarkResult) AddResult(rr RequestResult) {
	br.responseTimes = append(br.responseTimes, rr.ResponseTime)
}

func (br *BenchmarkResult) AddResultResponseTime(rtime int64) {
	br.responseTimes = append(br.responseTimes, rtime)
}

func (br BenchmarkResult) Count() int {
	return len(br.responseTimes)
}

func (br BenchmarkResult) Finish() bool {
	return br.Count() == br.MaxCount-1
}

func (br BenchmarkResult) Percentile(unit int) (int64, error) {
	if unit > 100 {
		return 0, errors.New("percentile is up to 100")
	}
	log.Print(br.responseTimes)
	sort.Slice(br.responseTimes, func(i, j int) bool {
		return br.responseTimes[i] < br.responseTimes[j]
	})
	length := len(br.responseTimes)
	i := length * unit / 100
	return br.responseTimes[i-1], nil
}
