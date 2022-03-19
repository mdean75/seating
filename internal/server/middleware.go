package server

import (
	"net/http"
)

func HandlePreFlight(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodOptions {
			w.Write([]byte("ok"))
			return
		}

		next.ServeHTTP(w, r)
	}
}

func logRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: MAKE THIS AN INITIALIZATION, MAYBE PART OF CONFIG
		//		location, _ := time.LoadLocation("America/Chicago")
		//if r.Context().Value("log") != "" {
		//	fmt.Println(r.Host, time.Now().Format(time.RFC822Z), r.Method, r.RequestURI, r.Proto, r.RemoteAddr, r.UserAgent())
		//}
		//fmt.Println(r.Host, time.Now().Format(time.RFC822Z), r.Method, r.RequestURI, r.Proto, r.RemoteAddr, r.UserAgent())
		//next.ServeHTTP(w, r)

	}
}
