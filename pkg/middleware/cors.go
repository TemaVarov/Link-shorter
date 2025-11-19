package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			next.ServeHTTP(w, r)
			return
		}
		header := w.Header()
		header.Set("Acess-Control-Allow-Origin", origin)
		header.Set("Acess-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			header.Set("Acess-Control-Allow-Methods", "GET, PUT, POST, DELETE, HEAD, PATCH")
			header.Set("Acess-Control-Allow-Headers", "autorisation, content-type, content-length")
			header.Set("Acess-Control-Max-Age", "86400")
		}
		next.ServeHTTP(w, r)
	})
}
