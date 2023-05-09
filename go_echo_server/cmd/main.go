package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dmilojkovic76/echo_server/go_echo_server/greetings"
	"github.com/dmilojkovic76/echo_server/go_echo_server/pkg/rest/router"
	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	hn, _ := os.Hostname()
	timestampFormat := "2006/01/02 15:04:05" // becouse chi Logger middleware logs: 2023/05/05 14:16:02 [dusanm-probook450/eUv71vwpEa-000001]
	fmt.Println(hn)
	log.SetPrefix(time.Now().Format(timestampFormat) + " [" + hn + "] ")
	// log.SetPrefix(time.Now().Format(time.RFC3339) + " - echo_server: ")
	log.SetFlags(0)

	// Output the StartUp message
	log.Print(greetings.StartUp())

    // loads values from .env into the system
	// First check if .env file exists
	if file_info, ef_err := os.Stat(".env"); ef_err != nil {
		log.Print("No '.env' file found! Using defaults!")
	} else {
		log.Print("Trying to load environment variables from '.env' file: " + file_info.Name() + "!\n")
		// If .env file exists, load it
		if err := godotenv.Load(); err != nil {
			log.Print("Error loading'.env' file. Proceding with defaults!")
		}
	}
}

func main() {
	http_port, http_port_exists := os.LookupEnv("HTTP_PORT")
	https_port, https_port_exists := os.LookupEnv("HTTPS_PORT")
	https_cert, https_cert_exists := os.LookupEnv("HTTPS_CERT")
	https_key, https_key_exists := os.LookupEnv("HTTPS_KEY")

	router := router.InitializeHandlers()

	// Run HTTPS server if the TLS certificate and key are present
	if https_cert_exists && https_key_exists && len(https_cert) > 0 && len(https_key) > 0 {
		if https_port_exists { // Run the server with the custom port
			log.Print("Server starting... Listening on port: " + https_port + " with TLS")
			err := http.ListenAndServeTLS(":" + https_port, https_cert, https_key, router)
			if err != nil {
				log.Fatal(err)
			}
		} else { // Run the server with the default port
			log.Print("Server starting... Listening on default port: 8443 with TLS")
			err := http.ListenAndServeTLS(":8443", https_cert, https_key, router)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else { // Run the server without TLS
		if http_port_exists { // Run the server with the custom port
			log.Print("Server starting... Listening on port: " + http_port)
			err := http.ListenAndServe(":" + http_port, router)
			if err != nil {
				log.Fatal(err)
			}
		} else { // Run the server with the default port
			log.Print("Server starting... Listening on default port: 8080")
			err := http.ListenAndServe(":8080", router)
			if err != nil {
				log.Fatal(err)
			}
			log.Print(err)
		}
	}
}
