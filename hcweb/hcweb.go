package hcweb

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andrewsjg/healthchecker/api"
	"github.com/andrewsjg/healthchecker/frontend"
	"github.com/andrewsjg/healthchecker/healthchecks"
)

func StartWebServer(chkConfig healthchecks.Healthchecks, testmode bool) {
	log.Println("Starting Web Server")

	// Serve the frontend
	frontendFS := frontend.NewFrontend()

	httpFS := http.FileServer(http.FS(frontendFS))
	http.Handle("/", httpFS)

	http.Handle("/api/v1/getConfig", http.HandlerFunc(api.GetConfig(chkConfig)))
	http.Handle("/api/v1/setConfig", http.HandlerFunc(api.SetConfig))

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", 8474), nil))

}
