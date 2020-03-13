package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    promFilePath, ok := os.LookupEnv("PROM_FILEPATH")
    if !ok {
        promFilePath = "web/prom.txt"
    }

    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request){
        log.Println("/metrics")
        http.ServeFile(w, r, promFilePath)
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}
