package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)


func health(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	ts := createtimestamb()

	jts, err := json.Marshal(ts)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jts)
}

func createtimestamb() map[string]interface{} {

	timeS := make(map[string]interface{})

	var layout = "Mon, 02 Jan 2006 15:04:05 GMT"

	timeS["unix"] = strconv.Itoa(int(time.Now().Unix()))
	timeS["utf"] = time.Now().Format(layout)

	return timeS
}

func getide(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ts := createtimestamb()

	var valueid = params.ByName("date")

	if _, ok := ts[valueid]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	val, _ := json.Marshal(ts[valueid])

	w.Header().Set("Content-Type", "application/json")

	w.Write(val)
	fmt.Println()
}

func main() {
	r := httprouter.New()

	r.GET("/api", health)
	r.GET("/api/:date", getide)
	
	serve :=&http.Server{
		Addr:         ":8000",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(serve.ListenAndServe())

}
