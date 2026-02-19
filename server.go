package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
	"github.com/fatih/color"
)

var totalRequests uint64

func startServer() {

	http.HandleFunc("/", handleForward)

	fmt.Println("\nServer running on http://localhost:2024\n")

	go liveMetrics()

	http.ListenAndServe(":2024", nil)
}

func handleForward(w http.ResponseWriter, r *http.Request) {

	start := time.Now()
	atomic.AddUint64(&totalRequests, 1)

	xURL := r.Header.Get("x-url")
	xProxy := r.Header.Get("x-proxy")
	xIdentifier := r.Header.Get("x-identifier")
	xSession := r.Header.Get("x-session-id")

	if xURL == "" {
		http.Error(w, "Missing x-url header", 400)
		return
	}

	client := buildClient(xProxy, xIdentifier)

	var body io.Reader
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		data, _ := io.ReadAll(r.Body)
		body = strings.NewReader(string(data))
	}

	req, err := http.NewRequest(r.Method, xURL, body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Forward non x- headers
	for key, values := range r.Header {
		if !strings.HasPrefix(strings.ToLower(key), "x-") {
			req.Header[key] = values
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), 500)
		logRequest(r, 0, xURL)
		return
	}
	defer resp.Body.Close()

	responseBody, _ := io.ReadAll(resp.Body)

	w.WriteHeader(resp.StatusCode)
	w.Write(responseBody)

	elapsed := time.Since(start)

	logRequest(r, resp.StatusCode, xURL)

	fmt.Printf("   ↳ Time: %v | Session: %s\n", elapsed, xSession)
}

func logRequest(r *http.Request, status int, target string) {

	ip, _, _ := net.SplitHostPort(r.RemoteAddr)

	statusStr := fmt.Sprintf("%d", status)

	switch {
	case status >= 200 && status < 300:
		color.New(color.FgHiGreen, color.Bold).Printf(
			"[REGISTERED] %-6s %s - from %-15s - %s\n",
			r.Method,
			statusStr,
			ip,
			target,
		)

	case status >= 400 && status < 500:
		color.New(color.FgHiYellow, color.Bold).Printf(
			"[REGISTERED] %-6s %s - from %-15s - %s\n",
			r.Method,
			statusStr,
			ip,
			target,
		)

	case status >= 500:
		color.New(color.FgHiRed, color.Bold).Printf(
			"[REGISTERED] %-6s %s - from %-15s - %s\n",
			r.Method,
			statusStr,
			ip,
			target,
		)

	default:
		fmt.Printf(
			"[REGISTERED] %-6s %s - from %-15s - %s\n",
			r.Method,
			statusStr,
			ip,
			target,
		)
	}
}

func liveMetrics() {
	for {
		time.Sleep(5 * time.Second)
		fmt.Printf("\n[METRICS] Total Requests: %d\n\n",
			atomic.LoadUint64(&totalRequests))
	}
}
