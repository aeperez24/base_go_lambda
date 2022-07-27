package integrationtest

import (
	"aeperez24/goLambda/config"
	"aeperez24/goLambda/handlers"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var appPort string

func init() {
	rand.Seed(time.Now().UnixNano())
	appPort = "111" + fmt.Sprintf("%02d", rand.Intn(99))
	log.Default().Printf("port is %s", appPort)
}
func RunTestWithIntegrationServerGin(testFunc func(integrationServerPort string)) {
	config.LoadViperConfig("../envs/", "isolation")
	server, port := createTestServerGin()
	go start(server)
	testFunc(port)
	server.Shutdown(context.Background())

}
func getHost(port string) string {
	return fmt.Sprintf("http://localhost:%s", port)
}

func BuildUrl(port string, path string) string {
	return getHost(port) + path
}
func start(server http.Server) {
	err := server.ListenAndServe()
	if err != nil {
		println(err)
	}
}
func createTestServerGin() (http.Server, string) {
	return handlers.BuildGinServer(appPort), appPort

}

func ExecuteHttpPostCallWithInterfaceBody(url string, bodyInterface interface{}, headers map[string]string) ([]byte, *http.Response, error) {
	body, _ := json.Marshal(bodyInterface)
	return executeHttpPostCall(url, body, headers)

}

func ExecuteHttpPostCallWithStringBody(url string, bodyString string, headers map[string]string) ([]byte, *http.Response, error) {
	body := []byte(bodyString)
	return executeHttpPostCall(url, body, headers)
}

func executeHttpPostCall(url string, body []byte, headers map[string]string) ([]byte, *http.Response, error) {
	postBuffer := bytes.NewBuffer(body)

	req, _ := http.NewRequest("POST", url, postBuffer)
	for name, value := range headers {
		req.Header.Add(name, value)
	}

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, _ := client.Do((req))
	bodyresp, err := ioutil.ReadAll(resp.Body)
	return bodyresp, resp, err
}
