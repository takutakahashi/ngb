package result

import "testing"

func TestPercentile(t *testing.T) {
	arr := []int64{}
	for i := int64(0); i < 500; i++ {
		arr = append(arr, i)
	}
	br := BenchmarkResult{responseTimes: arr}
	expected := 50
	if p, _ := br.Percentile(expected); p != 250 {
		t.Log(br.Percentile(expected))
		t.Fatal("miss")
	}
	expected = 95
	if p, _ := br.Percentile(expected); p != 475 {
		t.Log(br.Percentile(expected))
		t.Fatal("miss")
	}
	expected = 100
	if p, _ := br.Percentile(expected); p != 500 {
		t.Log(p)
		t.Fatal("miss")
	}
	expected = 101
	if p, err := br.Percentile(expected); err == nil {
		t.Log(p)
		t.Fatal("miss")
	}
}
