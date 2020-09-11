package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// main function
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		// Load environment variables
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		} else {
			port = os.Getenv("PORT")
		}
	}
	// verify $PORT again
	if port == "" {
		log.Fatalln("$PORT must be set")
	}

	http.HandleFunc("/", HTTPHandler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalln(err)
	}
}

// LogErr to log error
func LogErr(n int, err error) {
	if err != nil {
		log.Printf("Write failed: %v", err)
	}
}

// HTTPHandler function
func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	resp := GetIP(r) + "\n"
	LogErr(w.Write([]byte(resp)))
}

// GetIP gets IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address
func GetIP(r *http.Request) string {
	ipAddress := r.RemoteAddr
	forwardedAddress := r.Header.Get("X-Forwarded-For")

	if forwardedAddress != "" {
		//X-Forwarded-For
		ipAddress = forwardedAddress //Single IP

		//Array of IPs
		ip := strings.Split(forwardedAddress, ", ")
		if len(ip) > 1 {
			ipAddress = ip[0] //First IP
		}
	}
	return ipAddress
}
