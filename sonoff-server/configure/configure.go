package main

import (
//    "fmt"
//    "io"
    "net/http"
    "log"
    "os"
    "time"
)

func SonoffConfigServer(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte("{\"error\": 0,\"reason\": \"ok\",\"IP\": \"192.168.0.16\",\"port\": \"8443\"}"))
    // fmt.Fprintf(w, "This is an example server.\n")
    // io.WriteString(w, "This is an example server.\n")
}

func main() {
    http.HandleFunc("/dispatch/device", SonoffConfigServer)
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = "8442"
    }

    Sleep(600)

    err := http.ListenAndServeTLS(":"+port, "certificate.pem", "key.pem", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}