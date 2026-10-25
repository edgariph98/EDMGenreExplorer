package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// proxyHandler returns a Gin handler that reverse-proxies the request to
// the given upstream target, rewriting the path so that the /api/<prefix>
// segment is stripped before forwarding.
func proxyHandler(target *url.URL, stripPrefix string) gin.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// Strip the gateway prefix so the upstream receives only its own path
		// segment (e.g. /api/genres/graph → /graph).
		trimmed := strings.TrimPrefix(req.URL.Path, stripPrefix)
		if trimmed == "" {
			trimmed = "/"
		}
		req.URL.Path = trimmed
		req.URL.RawPath = trimmed
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("proxy error for %s: %v", r.URL, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, `{"error":"upstream service unavailable"}`)
	}

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	configPath := flag.String("config", getEnv("GATEWAY_CONFIG", "config.yaml"), "path to the gateway YAML config")
	flag.Parse()

	cfg, err := LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r := gin.Default()

	// Health check – useful for container orchestration readiness probes.
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	log.Printf("API Gateway listening on :%s (config: %s)", cfg.Server.Port, *configPath)
	for _, svc := range cfg.Services {
		handler := proxyHandler(svc.target, svc.Prefix)
		for _, method := range svc.Methods {
			r.Handle(method, svc.Prefix+"/*path", handler)
		}
		log.Printf("  %-14s → %s (%s) %v", svc.Prefix, svc.URL, svc.Name, svc.Methods)
	}

	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

// getEnv returns the value of the environment variable named by key, or
// fallback if the variable is not set / empty.
func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
