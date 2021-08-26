package main

/*
	Author : Iordanis Paschalidis
	Date   : 26/08/2021
*/

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Service object that contains the Port and Router of the application
type Service struct {
	Port   string
	Router *mux.Router
}

func Middleware(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	jsonFile, _ := os.Open("targets.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//var data interface{}
	//json.Unmarshal(byteValue, &data)
	w.WriteHeader(400)
	w.Write(byteValue)

}

/*
   Running the service in port 8000
       Endpoints:
		POST:
			/v1/rec
*/
func (s Service) run() {
	fmt.Println("server running at port " + s.Port)
	s.Router.HandleFunc("/v1/rec", Middleware).Methods("GET")

	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	handler := c.Handler(s.Router)
	http.ListenAndServe(s.Port, handler)
}

func main() {
	port := ":8000"

	service := Service{Port: port, Router: mux.NewRouter().StrictSlash(true)}
	service.run()
}
