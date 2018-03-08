package streamblast_api

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"encoding/json"
	"errors"
	"reflect"
)

func TestClient_buildEpisodeUrl(t *testing.T) {
	client := Client{
		BaseURI: "http://test.com",
		DreamsContentID: 100,
	}

	res := client.buildEpisodeURL("h23kkdfg")
	expected := "http://test.com/video/h23kkdfg/manifest_mp4.json?dreams_content_id=100"

	if res != expected {
		t.Errorf("I expected to get url \"%s\" but got \"%s\"", expected, res)
	}
}

func TestClient_GetLinksBadAnswer(t *testing.T) {
	handler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	client := Client{
		BaseURI: server.URL,
		DreamsContentID: 100,
	}

	_, err := client.GetLinks("hhh")
	if err == nil {
		t.Errorf("I expected to get error, but got nothing")
	}
}

func TestClient_GetLinksErrorCode(t *testing.T) {
	handler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.Encode(map[string]int{
			"error_code": 105,
		})
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	client := Client{
		BaseURI: server.URL,
		DreamsContentID: 100,
	}

	_, err := client.GetLinks("hhh")
	expectedError := errors.New("error code: 105")
	if err.Error() != expectedError.Error() {
		t.Errorf(
			"I expected to get error \"%s\", but got \"%s\"",
			expectedError,
			err,
		)
	}
}

func TestClient_GetLinksErrorMessage(t *testing.T) {
	handler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.Encode(map[string]string{
			"error": "Hello world",
		})
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	client := Client{
		BaseURI: server.URL,
		DreamsContentID: 100,
	}

	_, err := client.GetLinks("hhh")
	expectedError := errors.New("error: Hello world")
	if err.Error() != expectedError.Error() {
		t.Errorf(
			"I expected to get error \"%s\", but got \"%s\"",
			expectedError,
			err,
		)
	}
}

func TestClient_GetLinks(t *testing.T) {
	expected := map[string]string{
		"105": "http://example.com",
	}

	handler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		encoder := json.NewEncoder(w)
		encoder.Encode(&expected)
	}

	server := httptest.NewServer(http.HandlerFunc(handler))

	client := Client{
		BaseURI: server.URL,
		DreamsContentID: 100,
	}

	result, err := client.GetLinks("hhh")
	if err != nil {
		t.Errorf("I got unexpected error: %s", err)
	}

	equal := reflect.DeepEqual(expected, result)
	if !equal {
		t.Errorf("I expected to get %+v but got %+v", expected, result)
	}
}
