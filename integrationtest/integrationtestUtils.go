package integrationtest

import (
	"aeperez24/goLambda/config"
	"aeperez24/goLambda/handlers"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func RunTestWithIntegrationServerGin(testFunc func(port string)) {
	config.LoadViperConfig("../envs/", "isolation")
	server, port := createTestServerGin()
	go start(server)
	testFunc(port)
	server.Shutdown(context.Background())

}
func start(server http.Server) {
	err := server.ListenAndServe()
	if err != nil {
		println(err)
	}
}
func createTestServerGin() (http.Server, string) {
	port := "11082"
	return handlers.BuildGinServer(":" + port), port

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
