package httpUtils

import (
	"net/http"
	"os"
)

func CorsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			// log.Print("preflight detected: ", r.Header)
			w.Header().Add("Connection", "keep-alive")
			w.Header().Add("Access-Control-Allow-Origin", os.Getenv("ALLOWED_CORS_DOMAINS"))
			w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
			w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Add("Access-Control-Allow-Headers", "Token")
			w.Header().Add("Access-Control-Allow-Headers", "Authorization")
			w.Header().Add("Access-Control-Max-Age", "86400")
		} else {
			w.Header().Add("Access-Control-Allow-Origin", os.Getenv("ALLOWED_CORS_DOMAINS"))
			w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET")
			h.ServeHTTP(w, r)
		}
	}
}
