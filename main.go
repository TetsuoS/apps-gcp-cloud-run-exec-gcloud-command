package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/getGcloudConfigList", getGcloudConfigListHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

func getGcloudConfigListHandler(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("gcloud", "config", "list").Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Fprint(w, string(out))

	return
}
