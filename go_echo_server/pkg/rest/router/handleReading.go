package router

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func HandleReading(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now().Unix()
	path := chi.URLParam(r, "path")
	ctx := r.Context()
	type dataRequest struct {
		Path string `json:"path"`
		Body string `json:"body"`
		Headers string `json:"headers"`
		Query string `json:"query"`
	}
	var data dataRequest
	if err := json.NewDecoder(r.Body).Decode(&data.Body); err != nil { log.Println(err) }
	headers := r.Header.Get("headers")
	query := r.URL.Query().Get("query")
	log.Println("request received on: " + path)
	log.Println("request body: " + data.Body)
	log.Println("request headers: " + headers)
	log.Println("request query: " + query)
	
	select {
	case <-ctx.Done():
		processTime := time.Duration(time.Now().Unix() - timeStart).String()
		log.Println("request timed out after " + processTime)
		return
	default:
		processTime := time.Duration(time.Now().Unix() - timeStart).String()
		// render.PlainText(w, r, "Hello from GET on: " + path + " after " + processTime)
		render.JSON(w, r, "Hello from GET on: " + path + " after " + processTime)
		// render.XML(w, r, "Hello from GET on: " + path + " after " + processTime)
		// render.HTML(w, r, "Hello from GET on: " + path + " after " + processTime)
	}
	
}
