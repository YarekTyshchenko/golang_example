package handler

import (
	"testing"
	"log"
)

type mockStorage struct {}
func (s mockStorage) Store(r Request) (string, error) {
	return "uuid", nil
}

func TestHandle(t *testing.T) {
	s := mockStorage{}
	h := &Handler{storage: s}

	r := Request{
		Name: "foo",
		Company: "bar",
	}
	res, err := h.Handle(r)
	if err != nil {
		log.Fatal("Handler returned error", err)
	}
	if res.Id != "uuid" {
		log.Fatal("Handler doesn't return correct UUID")
	}
}