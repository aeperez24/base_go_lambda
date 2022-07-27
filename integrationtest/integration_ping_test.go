package integrationtest

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {

	RunTestWithIntegrationServerGin(func(port string) {
		expected := `{"message":"pong with :integration"}`
		url := fmt.Sprintf("http://localhost:%s/ping/", port)
		resp, _ := http.Get(url)
		bodyresp, _ := ioutil.ReadAll(resp.Body)
		assert.Equal(t, expected, string(bodyresp))
	})
}
