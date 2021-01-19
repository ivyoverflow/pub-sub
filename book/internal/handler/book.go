// Package handler contains all application handlers.
package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/ivyoverflow/pub-sub/book/internal/logger"
	"github.com/ivyoverflow/pub-sub/book/internal/model"
	"github.com/ivyoverflow/pub-sub/book/internal/service"
)

// Book contains all handlers for book.
type Book struct {
	ctx context.Context
	svc service.BookI
	log *logger.Logger
}

// NewBook returns a new configured Book object.
func NewBook(ctx context.Context, svc service.BookI, log *logger.Logger) *Book {
	return &Book{ctx, svc, log}
}

// Insert calls Insert service method and process POST requests.
func (handl *Book) Insert(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	request := model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handl.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusBadRequest, err.Error())

		return
	}

	createdBook, err := handl.svc.Insert(r.Context(), &request)
	if err != nil {
		handl.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	if err = json.NewEncoder(rw).Encode(createdBook); err != nil {
		handl.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	handl.log.Debug(fmt.Sprintf("Book <<< %s >>> added", createdBook.Name))
}

// Get calls Get service method and process GET requests.
func (handl *Book) Get(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	bookID := vars["id"]
	book, err := handl.svc.Get(r.Context(), bookID)
	if err != nil {
		handl.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	if err = json.NewEncoder(rw).Encode(&book); err != nil {
		handl.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	handl.log.Debug(fmt.Sprintf("Book <<< %s >>> sent", book.Name))
}

// Update calls Update service method and process UPDATE requests.
func (handl *Book) Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	bookID := vars["id"]
	request := model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handl.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusBadRequest, err.Error())

		return
	}

	updatedBook, err := handl.svc.Update(r.Context(), bookID, &request)
	if err != nil {
		handl.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	if err = json.NewEncoder(rw).Encode(&updatedBook); err != nil {
		handl.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	handl.log.Debug(fmt.Sprintf("Book <<< %s >>> updated", updatedBook.Name))
}

// Delete calls Delete service method and process DELETE requests.
func (handl *Book) Delete(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	bookID := vars["id"]
	deletedBook, err := handl.svc.Delete(r.Context(), bookID)
	if err != nil {
		handl.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	if err = json.NewEncoder(rw).Encode(&deletedBook); err != nil {
		handl.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	handl.log.Debug(fmt.Sprintf("Book <<< %s >>> deleted", deletedBook.Name))
}
