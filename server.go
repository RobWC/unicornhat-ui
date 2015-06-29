package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type UnicornServer struct {
}

func (us *UnicornServer) Listen(port int) error {
	r := mux.NewRouter()
	r.HandleFunc("/", us.homeHandler)
	r.HandleFunc("/api/v1/pixel/{id}/{r}/{b}/{g}", us.pixelHandler)
	return nil
}

func (us *UnicornServer) homeHandler(w http.ResponseWriter, r *http.Request) {

}

func (us *UnicornServer) pixelHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
}
