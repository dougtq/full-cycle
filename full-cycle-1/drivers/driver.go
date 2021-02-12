package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func loadDrivers(file string) []byte {
	jsonFile, err := os.Open(file)

	if err != nil {
		panic(err.Error())
	}

	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		panic(err.Error())
	}

	return data
}

func listDrivers(w http.ResponseWriter, r *http.Request) {
	drivers := loadDrivers("drivers.json")
	w.Write([]byte(drivers))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/drivers", listDrivers)
	http.ListenAndServe(":8081", r)
}
