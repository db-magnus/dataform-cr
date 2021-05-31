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

        cmd := exec.Command("/bin/bash", "script.sh")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
       	 log.Fatalf("script failed with ",err)
    } 

	w.Write([]byte("script finished\n"))
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
