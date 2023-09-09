package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

type RequestData struct {
	Method  string
	URL     string
	Headers http.Header
	Body    string
}

type ResponseData struct {
	StatusCode int
	Headers    http.Header
	Body       string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.ListenAndServe(":8080", loggingMiddleware(http.DefaultServeMux))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract data from the request
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		requestData := RequestData{
			Method:  r.Method,
			URL:     r.URL.String(),
			Headers: r.Header,
			Body:    string(bodyBytes),
		}
		log.Printf("Request Data: %+v\n", requestData)

		// Create a response writer to capture the response
		recorder := httptest.NewRecorder()

		// Call the next handler
		next.ServeHTTP(recorder, r)

		// Extract data from the response
		responseData := ResponseData{
			StatusCode: recorder.Code,
			Headers:    recorder.Header(),
			Body:       recorder.Body.String(),
		}
		log.Printf("Response Data: %+v\n", responseData)

		// Copy the recorded response to the actual response writer
		for k, v := range recorder.Header() {
			w.Header()[k] = v
		}
		w.WriteHeader(recorder.Code)
		w.Write(recorder.Body.Bytes())
	})
}
