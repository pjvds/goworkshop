package httpcallback

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

var (
	_ = Suite(&ClientScenarioTests{})
)

type ClientScenarioTests struct {
	server *httptest.Server
}

// send http request to the api: /callback
func (this *ClientScenarioTests) mustRequestCallbacks(document map[string]interface{}) *http.Response {
	json, _ := json.Marshal(document)
	resp, err := http.Post(this.server.URL, "application/json", bytes.NewBuffer(json))

	if err != nil {
		panic(err)
	}

	return resp
}

func (this *ClientScenarioTests) SetUpSuite(c *C) {
	this.server = mustGetApiTestServer()
}

func (this *ClientScenarioTests) TearDownSuite(c *C) {
	if this.server != nil {
		this.server.Close()
	}
}

func (this *ClientScenarioTests) TestCreateCallback(c *C) {
	server := mustGetApiTestServer()
	defer server.Close()

	resp := this.mustRequestCallbacks(map[string]interface{}{
		// the url to receive a callback
		"url":   "http://localhost:8081/callback_target",
		"after": 5,
	})

	c.Assert(resp.StatusCode, Equals, http.StatusCreated)
}

func (this *ClientScenarioTests) TestServerCallsbackAfterInterval(c *C) {
	server := mustGetApiTestServer()
	defer server.Close()

	request := make(chan *http.Request)

	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		request <- req
	})

	target := httptest.NewServer(handler)
	defer target.Close()

	resp := this.mustRequestCallbacks(map[string]interface{}{
		"url":   target.URL,
		"after": 5,
	})

	c.Assert(resp.StatusCode, Equals, http.StatusCreated)

	timeout := 1 * time.Second
	select {
	case <-request:
		// as expected!
		// todo: assert interval
	case <-time.After(timeout):
		c.Fatalf("timeout: callback server did not callback within %v", timeout.String())
	}
}

/*func (this *ClientScenarioTests) TestCreateMultipleCallbackShouldAllCallback(c *C) {
	server := mustGetApiTestServer()
	defer server.Close()

	numOfRequest := 255 // the number of requests we will do
	numOfCallback := 0  // the number of callbacks received, this should eventually be equal to requests

	request := make(chan *http.Request, numOfRequest)

	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		request <- req
	})

	target := httptest.NewServer(handler)
	defer target.Close()

	for i := 0; i < numOfRequest; i++ {
		go func() {
			resp := this.mustRequestCallbacks(map[string]interface{}{
				"url":   target.URL,
				"after": 5,
			})
			c.Assert(resp.Status, Equals, http.StatusCreated)
		}()
	}

	timeout := 1 * time.Second

	for {
		select {
		case <-request:
			numOfCallback++
			// as expected!
		case <-time.After(timeout):
			c.Fatal("timeout: callback server did not callback within %v", timeout)
		}
	}

	c.Assert(numOfCallback, Equals, numOfRequest)
}*/

// returns an httptest server instance that is running the API
func mustGetApiTestServer() *httptest.Server {
	apiEndpoint, err := createApiServer()
	if err != nil {
		panic(err)
	}

	return httptest.NewServer(apiEndpoint)
}
