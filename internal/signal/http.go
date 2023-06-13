package signal

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// HTTPSDPServer starts a HTTP Server that consumes SDPs
func HTTPSDPServer() chan string {
	port := flag.Int("port", 8080, "http server port")
	flag.Parse()

	sdpChan := make(chan string)

	// // Middleware để xử lý CORS
	// corsMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
	// 	return func(w http.ResponseWriter, r *http.Request) {
	// 		// Thiết lập các header CORS
	// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
	// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 		w.Header().Set("Access-Control-Allow-Credentials", "true")

	// 		// Tiếp tục xử lý yêu cầu tiếp theo
	// 		next(w, r)
	// 	}
	// }

	http.HandleFunc("/sdp", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprintf(w, "done")
		sdpChan <- string(body)
	})

	go func() {
		err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
		if err != nil {
			panic(err)
		}
	}()

	return sdpChan
}