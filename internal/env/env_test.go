// gonotes/internal/env/env_test.go

package env

import (
	"testing"
)

func TestEnv(t *testing.T) {
	id := 1

	if Env == nil {
		t.Fatalf("tests[%d] - Env wrong, expected=(non nil), got=(nil)", id)
	}
}

func TestHome(t *testing.T) {
	id := 10
	home := Home()
	if home == "" {
		t.Fatalf("tests[%d] - Home() returns empty", id)
	}
}

func TestConfigHome(t *testing.T) {
	id := 20
	confdir := ConfigHome()
	if confdir == "" {
		t.Fatalf("tests[%d] - ConfigHome() returns empty", id)
	}
}
