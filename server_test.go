package main

import (
	_ "embed"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//go:embed testdata/gopher.sixel
var testimage1 string

func TestGETImages(t *testing.T) {
	store := StubSixelStore{
		map[string]string{
			"testimage1": testimage1,
		},
	}
	server := NewSixelServer(&store)

	t.Run("return test image 1", func(t *testing.T) {
		request := newGetImageRequest("testimage1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertImageServed(t, response.Body.String(), testimage1)
	})

	t.Run("return 404 on invalid image", func(t *testing.T) {
		request := newGetImageRequest("doesnotexist")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
		assertResponseBody(t, response.Body.String(), "Requested image not found.")
	})
}

func TestStoreImages(t *testing.T) {
	store := StubSixelStore{map[string]string{}}
	server := NewSixelServer(&store)

	t.Run("saves and then stores POSTed image", func(t *testing.T) {
		request := newPostImageRequest("testimage1", testimage1)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		request = newGetImageRequest("testimage1")
		response = httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertImageServed(t, response.Body.String(), testimage1)
	})

	// TODO
	t.Run("rejects uploads are not valid sixel strings", func(t *testing.T) {

	})
}

type StubSixelStore struct {
	images map[string]string
}

func (s *StubSixelStore) GetSixelImage(id string) string {
	image := s.images[id]
	return image
}

func (s *StubSixelStore) StoreSixelImage(id, image string) {
	s.images[id] = image
}

func newGetImageRequest(id string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/image/%s", id), nil)
	return request
}

func newPostImageRequest(id, image string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/upload/%s", id), strings.NewReader(image))
	return request
}

func assertImageServed(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Error("did not get expected sixel image string after upload")
	}
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body did not match, got %s, want %s", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
