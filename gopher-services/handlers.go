package main

import (
	"net/http"
	"html/template"
	"gophers-network/util"
	"time"
	"github.com/oklog/ulid"
	"math/rand"
	"gophers-network/schema"
	"gophers-network/db"
	"log"
	"strconv"
)

func generateULID() (ulid.ULID, error) {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return ulid.New(ulid.Timestamp(t), entropy)
}

func createGopherHandler(writer http.ResponseWriter, r *http.Request) {
	type response struct {
		ID string `json: "id"`
	}

	ctx := r.Context()

	body := template.HTMLEscapeString(r.FormValue("body"))
	if len(body) < 1 || len(body) > 140 {
		util.ResponseError(writer, http.StatusBadRequest, "Invalid body")
	}

	// Generating ULID
	id, err := generateULID()
	if err != nil {
		util.ResponseError(writer, http.StatusInternalServerError, "Failed to create ULID")
		return
	}

	// Creating Gopher
	gopher := schema.Gopher{
		ID:        id.String(),
		Body:      body,
		CreatedAt: time.Now().UTC(),
	}

	if err := db.InsertGopher(ctx, gopher); err != nil {
		util.ResponseError(writer, http.StatusInternalServerError, "Failed to create gopher")
		return
	}

	util.ResponseOk(writer, response{ID: gopher.ID})
}

func listGophersHandler(writer http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error

	offset := uint64(0)
	offsetStr := r.FormValue("offset")
	limit := uint64(100)
	limitStr := r.FormValue("limit")
	if len(offsetStr) != 0 {
		offset, err = strconv.ParseUint(offsetStr, 10, 64)
		if err != nil {
			util.ResponseError(writer, http.StatusBadRequest, "Invalid offset parameter")
			return
		}
	}
	if len(limitStr) != 0 {
		limit, err = strconv.ParseUint(limitStr, 10, 64)
		if err != nil {
			util.ResponseError(writer, http.StatusBadRequest, "Invalid limit parameter")
			return
		}
	}

	gophers, err := db.ListGophers(ctx, offset, limit)
	if err != nil {
		log.Println(err)
		util.ResponseError(writer, http.StatusInternalServerError, "Could not fetch gophers")
		return
	}

	util.ResponseOk(writer, gophers)
}