package main

import (
    "fmt"
    "net/http"
    "strconv"
    "time"
    "flag"
)

func handler(w http.ResponseWriter, r *http.Request) {
    params           := r.URL.Query()
    id               := params["id"][0]
    sleep_seconds, _ := strconv.Atoi( params["t"][0] )
    status, _        := strconv.Atoi( params["s"][0] )
    price            := 0
    if params["p"] != nil {
        price, _ = strconv.Atoi( params["p"][0] )
    }

    if sleep_seconds > 0 {
        time.Sleep( time.Second * time.Duration( sleep_seconds ) ) // milliseconds -> seconds
    }

    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "{\"id\":\"%s\",\"status\":%d,\"price\":%d}", id, status, price)
}

func main() {
    port := flag.Int("port", 8080, "PORT")
    flag.Parse()

    http.HandleFunc("/ad", handler)
    http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
