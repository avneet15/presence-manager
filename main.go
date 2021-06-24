package main

import (
	"fmt"
	"log"
	"net/http"
	"presence-manager/datastore"
)

var conn = datastore.New()

func processEntryHandler(w http.ResponseWriter, r *http.Request) {
	entryId := r.URL.Path[15:]
	conn.UpsertEntry(entryId, 15)
	fmt.Fprintf(w, "%s", entryId)
}

func getEntryStateHandler(w http.ResponseWriter, r *http.Request) {
	entryId := r.URL.Path[17:]
	isAlive := conn.IsEntryPresent(entryId)
	fmt.Fprintf(w, "%t", isAlive)
}

func main(){
	fmt.Println("Starting service..")

	http.HandleFunc("/process-entry/", processEntryHandler)
	http.HandleFunc("/get-entry-state/", getEntryStateHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
