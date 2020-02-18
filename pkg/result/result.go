package result

import (
	"log"

	types "github.com/takutakahashi/ngb/pkg/types/result"
)

func Analyze(br types.BenchmarkResult) int {
	p95, err := br.Percentile(95)
	if err != nil {
		log.Fatal(err)
		return 1
	}
	log.Printf("95 percentile: %dms", p95)
	return 0
}
