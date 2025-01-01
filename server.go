package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type SixelServer struct {
	store SixelStore
	http.Handler
}

type SixelStore interface {
	GetSixelImage(id string) string
	StoreSixelImage(id, image string)
}

func NewSixelServer(store SixelStore) *SixelServer {
	s := new(SixelServer)

	s.store = store

	router := http.NewServeMux()
	router.Handle("/image/", http.HandlerFunc(s.getImageHandler))
	router.Handle("/upload/", http.HandlerFunc(s.postImageHandler))

	s.Handler = router
	return s
}

func (s *SixelServer) getImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := strings.TrimPrefix(r.URL.Path, "/image/")

	image := s.store.GetSixelImage(id)

	if image == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Requested image not found.")
		return
	}

	fmt.Fprint(w, image)
}

func (s *SixelServer) postImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := strings.TrimPrefix(r.URL.Path, "/upload/")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	s.store.StoreSixelImage(id, string(body))
}
