package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
        "os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
        log.Print("helloworld: received a request")

        cmd := exec.Command("/bin/sh", "script.sh")
        cmd.Stderr = os.Stderr
        out, err := cmd.Output()
        if err != nil {
		log.Fatal(err)
                w.WriteHeader(500)
        }
        fmt.Printf("output logging ",out)
        w.Write(out)
}

func main() {
        log.Print("helloworld: starting server...")

        http.HandleFunc("/", handler)

        port := os.Getenv("PORT")
        if port == "" {
                port = "8080"
        }

        log.Printf("helloworld: listening on %s", port)
        log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
