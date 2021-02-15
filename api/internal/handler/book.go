// Package handler contains all application handlers.
package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"

	"github.com/ivyoverflow/pub-sub/api/internal/lib/types"
	"github.com/ivyoverflow/pub-sub/api/internal/model"
	"github.com/ivyoverflow/pub-sub/api/internal/service"
	"github.com/ivyoverflow/pub-sub/platform/logger"
)

// BookController contains all handlers for book.
type BookController struct {
	ctx context.Context
	svc service.Booker
	log *logger.Logger
}

// NewBookController returns a new configured Book object.
func NewBookController(ctx context.Context, svc service.Booker, log *logger.Logger) *BookController {
	return &BookController{ctx, svc, log}
}

// AbortWithError sends a error response on error.
func AbortWithError(rw http.ResponseWriter, statusCode int, err error) {
	rw.WriteHeader(statusCode)
	fmt.Fprintf(rw, `{"error": {"statusCode": %d, "message": "%s"}}`, statusCode, err.Error())
}

// Insert calls Insert service method and process POST requests.
func (h *BookController) Insert(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	request := model.Book{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.log.Error(err.Error())
		AbortWithError(rw, http.StatusBadRequest, types.ErrorBadRequest)

		return
	}

	insertedBook, err := h.svc.Insert(r.Context(), &request)
	if err != nil {
		h.log.Error(err.Error())
		switch err {
		case types.ErrorDuplicateValue:
			AbortWithError(rw, http.StatusConflict, types.ErrorDuplicateValue)

			return
		case types.ErrorValidation:
			AbortWithError(rw, http.StatusBadRequest, types.ErrorValidation)

			return
		default:
			AbortWithError(rw, http.StatusInternalServerError, types.ErrorInternalServerError)

			return
		}
	}

	rw.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(rw).Encode(insertedBook); err != nil {
		h.log.Error(err.Error())
		AbortWithError(rw, http.StatusInternalServerError, types.ErrorInternalServerError)

		return
	}

	h.log.Debug(fmt.Sprintf("Book <<< %s >>> added", insertedBook.Name))
}

// Get calls Get service method and process GET requests.
func (h *BookController) Get(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	bookID, err := uuid.Parse(vars["id"])
	if err != nil {
		AbortWithError(rw, http.StatusInternalServerError, types.ErrorInternalServerError)

		return
	}

	book, err := h.svc.Get(r.Context(), bookID)
	if err != nil {
		h.log.Error(err.Error())
		switch err {
		case types.ErrorNotFound:
			AbortWithError(rw, http.StatusNotFound, types.ErrorNotFound)

			return
		default:
			AbortWithError(rw, http.StatusInternalServerError, types.ErrorInternalServerError)

			return
		}
	}

	if err = json.NewEncoder(rw).Encode(&book); err != nil {
		h.log.Error(err.Error())
		AbortWithError(rw, http.StatusInternalServerError, types.ErrorInternalServerError)

		return
	}

	h.log.Debug(fmt.Sprintf("Book <<< %s >>> sent", book.Name))
}

// Update calls Update service method and process UPDATE requests.
func (h *BookController) Update(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	bookID, err := uuid.Parse(vars["id"])
	if err != nil {
		AbortWithError(rw, http.StatusInternalServerError, types.ErrorInternalServerError)

		return
	}

	request := model.Book{}
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.log.Error(err.Error())
		AbortWithError(rw, http.StatusBadRequest, types.ErrorBadRequest)

		return
	}

	updatedBook, err := h.svc.Update(r.Context(), bookID, &request)
	if err != nil {
		h.log.Error(err.Error())
		switch err {
		case types.ErrorNotFound:
			AbortWithError(rw, http.StatusNotFound, types.ErrorNotFound)

			return
		case types.ErrorDuplicateValue:
			AbortWithError(rw, http.StatusConflict, types.ErrorDuplicateValue)

			return
		case types.ErrorValidation:
			AbortWithError(rw, http.StatusBadRequest, types.ErrorValidation)

			return
		default:
			AbortWithError(rw, http.StatusInternalServerError, types.ErrorInternalServerError)

			return
		}
	}

	rw.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(rw).Encode(&updatedBook); err != nil {
		h.log.Error(err.Error())
		AbortWithError(rw, http.StatusInternalServerError, types.ErrorInternalServerError)

		return
	}

	h.log.Debug(fmt.Sprintf("Book <<< %s >>> updated", updatedBook.Name))
}

// Delete calls Delete service method and process DELETE requests.
func (h *BookController) Delete(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("content-type", "application/json")
	vars := mux.Vars(r)
	bookID, err := uuid.Parse(vars["id"])
	if err != nil {
		AbortWithError(rw, http.StatusInternalServerError, types.ErrorInternalServerError)

		return
	}

	deletedBook, err := h.svc.Delete(r.Context(), bookID)
	if err != nil {
		h.log.Error(err.Error())
		switch {
		case errors.Cause(err) == types.ErrorNotFound:
			AbortWithError(rw, http.StatusNotFound, types.ErrorNotFound)

			return
		default:
			AbortWithError(rw, http.StatusInternalServerError, types.ErrorInternalServerError)

			return
		}
	}

	rw.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(rw).Encode(&deletedBook); err != nil {
		h.log.Error(err.Error())
		AbortWithError(rw, http.StatusInternalServerError, types.ErrorInternalServerError)

		return
	}

	h.log.Debug(fmt.Sprintf("Book <<< %s >>> deleted", deletedBook.Name))
}
