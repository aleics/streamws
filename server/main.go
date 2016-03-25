package main

import (
    "flag"
    "log"
    "net/http"
    
    "github.com/aleics/streamws/server/video"
    "github.com/aleics/streamws/server/image"
)

// address flag
var addr = flag.String("addr", "localhost:8080", "http service address")

// main function
func main() {
    flag.Parse()
    log.SetFlags(0)
    http.HandleFunc("/streamwsvideo", video.HandlerVideo)
    
    http.HandleFunc("/streamwsimage", image.HandlerImage)
    http.HandleFunc("/streamwsimage/freq", image.Freq)
    log.Fatal(http.ListenAndServe(*addr, nil))
}