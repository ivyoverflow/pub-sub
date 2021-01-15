// Package handler contains all application handlers.
package handler

import (
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
	svc *service.Manager
	log *logger.Logger
}

// NewBook returns a new configured Book object.
func NewBook(svc *service.Manager, log *logger.Logger) *Book {
	return &Book{svc, log}
}

// Add calls Add service method and process POST requests.
func (handler *Book) Add(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	request := model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handler.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusBadRequest, err.Error())

		return
	}

	createdBook, err := handler.svc.Book.Add(&request)
	if err != nil {
		handler.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	if err = json.NewEncoder(rw).Encode(createdBook); err != nil {
		handler.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	handler.log.Debug(fmt.Sprintf("Book <<< %s >>> added", createdBook.Name))
}

// Get calls Get service method and process GET requests.
func (handler *Book) Get(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	bookID := vars["id"]
	book, err := handler.svc.Book.Get(bookID)
	if err != nil {
		handler.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	if err = json.NewEncoder(rw).Encode(&book); err != nil {
		handler.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	handler.log.Debug(fmt.Sprintf("Book <<< %s >>> sent", book.Name))
}

// Update calls Update service method and process UPDATE requests.
func (handler *Book) Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	bookID := vars["id"]
	request := model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		handler.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusBadRequest, err.Error())

		return
	}

	updatedBook, err := handler.svc.Book.Update(bookID, &request)
	if err != nil {
		handler.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	if err = json.NewEncoder(rw).Encode(&updatedBook); err != nil {
		handler.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	handler.log.Debug(fmt.Sprintf("Book <<< %s >>> updated", updatedBook.Name))
}

// Delete calls Delete service method and process DELETE requests.
func (handler *Book) Delete(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	bookID := vars["id"]
	deletedBook, err := handler.svc.Book.Delete(bookID)
	if err != nil {
		handler.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	if err = json.NewEncoder(rw).Encode(&deletedBook); err != nil {
		handler.log.Error(err.Error())
		fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, http.StatusInternalServerError, err.Error())

		return
	}

	handler.log.Debug(fmt.Sprintf("Book <<< %s >>> deleted", deletedBook.Name))
}
