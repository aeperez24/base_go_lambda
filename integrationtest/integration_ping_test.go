package integrationtest

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {

	RunTestWithIntegrationServerGin(func(port string) {
		expected := `{"message":"pong with :integration"}`
		url := BuildUrl(port, "/ping/")
		resp, _ := http.Get(url)
		bodyresp, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, expected, string(bodyresp))
	})
}

func TestHello(t *testing.T) {

	RunTestWithIntegrationServerGin(func(port string) {
		request := `{"name":"MyName"}`
		expected := `{"message":"hello :MyName"}`
		url := BuildUrl(port, "/hello/")
		bodyresp, _, _ := ExecuteHttpPostCallWithStringBody(url, request, nil)
		assert.Equal(t, expected, string(bodyresp))
	})
}
