package chai

import "net/http"

func main() {
	// net/http
	mux := http.NewServeMux()
	mux.Handle("/api/get", http.HandlerFunc(hello))
	err := http.ListenAndServe(":8080", mux)
}

type Hello struct {
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func hello(w http.ResponseWriter, r *http.Request) {

}