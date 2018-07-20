package main
import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
func GetSomething(response http.ResponseWriter, request *http.Request) {
	log.Print("Got a request for something from " + request.Host)
	json.NewEncoder(response).Encode("Something")
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getSomething", GetSomething).Methods("GET")
	http.ListenAndServe(":8080", router)
}
