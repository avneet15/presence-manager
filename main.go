package main

import (
	"fmt"
	"log"
	"net/http"
	"presence-manager/datastore"
	"presence-manager/notifications"
	"time"
)


func processEntryHandler(conn *datastore.Store, notifier notifications.LoggerNotifier) func(w http.ResponseWriter, r *http.Request) {
	if conn == nil {
		panic("Empty datastore connection!")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		entryId := r.URL.Path[15:]
		conn.UpsertEntry(entryId, 15)
		time.AfterFunc(time.Second*30, func() {
			isAlive := conn.IsEntryPresent(entryId)
			if !isAlive {
				notifier.Notify(entryId)
			}
		})
		fmt.Fprintf(w, "%s", entryId)
	}

}

func getEntryStateHandler(conn *datastore.Store) func(w http.ResponseWriter, r *http.Request) {
	if conn == nil {
		panic("Empty datastore connection!")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		entryId := r.URL.Path[17:]
		isAlive := conn.IsEntryPresent(entryId)
		fmt.Fprintf(w, "%t", isAlive)
	}

}

func main(){
	conn := datastore.New()
	mux := http.NewServeMux()
	fmt.Println("Starting service..")

	mux.HandleFunc("/process-entry/", processEntryHandler(conn, notifications.LoggerNotifier{}))
	mux.HandleFunc("/get-entry-state/", getEntryStateHandler(conn))
	log.Fatal(http.ListenAndServe(":8080", mux))

}
