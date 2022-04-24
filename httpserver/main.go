package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/renbozhou/golang/httpserver/metrics"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "v1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)

	for k, v := range r.Header {
		fmt.Println(k, v)
		for _, vv := range v {
			w.Header().Set(k, vv)
		}

	}

	clientIP := getCurrentIP(r)
	log.Printf("client ip %s\n", clientIP)
}

func getCurrentIP(r *http.Request) string {
	// ip: port
	fmt.Println(r.RemoteAddr)
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		fmt.Println(r.RemoteAddr)
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func healthz(w http.ResponseWriter, r *http.Request) {

	a := []byte("server up 200")
	w.Write(a)
}
func images(w http.ResponseWriter, r *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	w.Write([]byte(fmt.Sprintln("<h1>%d</h1>", randInt)))
}
func main() {
	fmt.Println("main")
	server := http.NewServeMux()

	server.HandleFunc("/", index)
	server.HandleFunc("/images", images)
	server.Handle("/metrics", promhttp.Handler())
	server.HandleFunc("/healthz", healthz)

	if err := http.ListenAndServe(":18080", server); err != nil {
		log.Fatalf("http failed,err:%s\n ", err.Error())
	}
}
