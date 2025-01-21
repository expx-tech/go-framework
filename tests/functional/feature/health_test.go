//go:build functional

package feature

import (
	"net/http"
	"testing"

	lib "github.com/expx-tech/go-framework/tests/pkg"
)

func TestHealthz(t *testing.T) {
	api := lib.NewApi()
	resp := new(http.Response)
	err := api.GetHttp("/healthz", resp)
	if err != nil {
		t.Fatalf("unexpeted error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got: %v", resp.StatusCode)
	}
}
