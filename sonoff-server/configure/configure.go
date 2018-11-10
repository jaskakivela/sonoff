package main

import (
    "net/http"
    "log"
    "os"
)

func SonoffConfigServer(w http.ResponseWriter, req *http.Request) {

    ip := os.Getenv("EXTERNALIP")
    if len(ip) == 0 {
        ip = "192.168.0.16"
    }
    
    port := os.Getenv("SERVERPORT")
    if len(port) == 0 {
        port = "8443"
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write([]byte("{\"error\": 0,\"reason\": \"ok\",\"IP\": \""+ip+"\",\"port\": \""+port+""\"}"))
}

func main() {

    http.HandleFunc("/dispatch/device", SonoffConfigServer)
    port := os.Getenv("CONFPORT")
    if len(port) == 0 {
        port = "8442"
    }


    err := http.ListenAndServeTLS(":"+port, "/certificate.pem/sslcert", "/key.pem/sslkey", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}