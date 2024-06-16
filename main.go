package main

import (
	"log"
	"net/http"
	fileOps "simpleapi/pkg"
	"strconv"
	"time"
)

type api_req struct {
	file_path string
}

func (ar *api_req) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	count := 0
	ar.file_path = "request_log.csv"
	request_data := []string{time.Now().UTC().String(), r.Method}
	fileOps.WriteFile(request_data)
	data := fileOps.GetData(ar.file_path)

	//Count the number of request in last one minute
	for _, v := range data {

		temp_time, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", v[0])
		if temp_time.After(time.Now().UTC().Add(time.Minute * -1)) {
			count = count + 1
		}
	}

	w.Write([]byte("No of request in last 1 min is: " + strconv.Itoa(count)))
}

func main() {
	mux := http.NewServeMux()

	count_Request := api_req{}

	mux.Handle("/", &count_Request)

	log.Println("Listening on port 8080.")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
