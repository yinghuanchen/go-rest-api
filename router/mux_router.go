package router

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type MuxRouter struct {
	*mux.Router
}

func NewMuxRouter() Router {
	router := mux.NewRouter()
	return &MuxRouter{Router: router,}
}

func (m *MuxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.Router.HandleFunc(uri, f).Methods("GET")
}
func (m *MuxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	m.Router.HandleFunc(uri, f).Methods("POST")
}
func (m *MuxRouter) SERVE(port string) {
	log.Println("Server listening on port ", port)
	if err := http.ListenAndServe(port, m.Router); err != nil {
		log.Fatal(err)
	}
}