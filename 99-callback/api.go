package httpcallback

import (
	"encoding/json"
	"net/http"
	"time"
)

type server struct {
}

func (this *server) addWork(url string, when time.Time) {
	// todo: add the work and find a way to
	// callback to the url in the request
}

func (this *server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Content-Type", "application/json")

	var request struct {
		Url   string `json:"url"`
		After int64  `json:"after"`
	}

	var decoder = json.NewDecoder(req.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(rw, "invalid json", http.StatusBadRequest)
		return
	}

	when := time.Now().Add(time.Duration(request.After) * time.Millisecond)
	this.addWork(request.Url, when)

	rw.WriteHeader(http.StatusCreated)
}

func createApiServer() (http.Handler, error) {
	server := &server{
	// you can set fields here, like: foo: "bar"
	}

	return server, nil
}
