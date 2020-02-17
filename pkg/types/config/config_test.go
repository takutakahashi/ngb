package config

import (
	"testing"
)

func TestScript(t *testing.T) {
	s := Script{Path: "../../../src/prerequest.sh"}
	rc, err := s.Execute()
	t.Logf("rc: %v", rc)
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	if _, ok := rc.Params["user[email]"]; !ok {
		t.Fatalf("rc: %v", rc)
	}
	if _, ok := rc.Cookies[0]["key"]; !ok {
		t.Fatalf("rc: %v", rc)
	}
}

func TestRequestConfig(t *testing.T) {}
