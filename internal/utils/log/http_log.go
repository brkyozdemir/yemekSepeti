package log

import (
	"fmt"
	"net/http"
	"time"
)

type HTTPReqInfo struct {
	method   string
	uri      string
	duration interface{}
	proto    string
	code     int
	ipAddr   string
}

type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &StatusRecorder{
			ResponseWriter: w,
			Status:         200,
		}

		begin := time.Now().UnixNano()
		next.ServeHTTP(w, r)
		end := time.Now().UnixNano()
		ns := end - begin

		info := HTTPReqInfo{
			method:   r.Method,
			uri:      r.RequestURI,
			duration: time.Duration(ns),
			proto:    r.Proto,
			code:     recorder.Status,
			ipAddr:   r.RemoteAddr,
		}

		fmt.Println(colorMap[info.method], info.method, info.uri, info.duration, info.proto, info.ipAddr, info.code)
	})
}

var colorMap = map[string]string{
	"POST": "\033[33m",
	"GET":  "\033[34m",
}
