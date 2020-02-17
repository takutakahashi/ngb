package load

import (
	"net/url"
	"os"
	"testing"

	"github.com/takutakahashi/ngb/pkg/types/config"
)

func TestRequest(t *testing.T) {
	url, _ := url.Parse("http://localhost:3000/users")
	s := config.ScriptConfig{
		PreRequest: config.Script{Path: os.Getenv("PROJECT_ROOT") + "/src/prerequest.sh"},
	}
	rr, err := Request(url, s)
	if err != nil || rr.StatusCode != 200 {
		t.Log(rr)
		t.Fatal(err)
	}
}

func TestBenchmark(t *testing.T) {
	u, _ := url.Parse("http://localhost:3000/users")
	s := config.ScriptConfig{
		PreRequest: config.Script{Path: os.Getenv("PROJECT_ROOT") + "/src/prerequest.sh"},
	}
	br, _ := Start(u, 2, 10, s)
	t.Fatal(br.Percentile(95))

}
