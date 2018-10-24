package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//"gopkg.in/mgo.v2/bson"
//Todo struct to todo
type Todo struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Created   time.Time `json:"createdon"`
}

func main() {
	router := NewRouter() // this func is in router.go // related to Session in handlers.go
	log.Fatal(http.ListenAndServe(":8080", router))
}

func JSONResponse(w http.ResponseWriter, r *http.Request, start time.Time, response []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	log.Printf("%s\t%s\t%s\t%s\t%d\t%d\t%s",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
		code,
		len(response),
		time.Since(start),
	)
	if string(response) != "" {
		w.Write(response)
	}
}

//NotFound responses to routes not defined
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s\t%s\t%d\t%d\t%d",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
		http.StatusNotFound,
		0,
		0,
	)
	w.WriteHeader(http.StatusNotFound)
}

//NewRouter creates the router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/todos", TodoIndex).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(NotFound)
	return r
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var todo Todo
	todo.ID = "toto"
	response, err := json.MarshalIndent(todo, "", "    ")
	if err != nil {
		panic(err)
	}
	JSONResponse(w, r, start, response, http.StatusOK)
}
