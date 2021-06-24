package main

import (
	"fmt"
	"presence-manager/datastore"
)

func main(){
	fmt.Println("hello")
	ds := datastore.New()
	//ds.UpsertEntry("A", 5)
	//
	//time.Sleep(time.Duration(3)*time.Second)
	//ds.UpsertEntry("A", 30)
	//time.Sleep(time.Duration(25)*time.Second)
	res := ds.IsEntryPresent("A")
	fmt.Println(res)
}
