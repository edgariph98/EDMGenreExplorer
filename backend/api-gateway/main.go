package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// serviceConfig holds the upstream address for a backend service.
type serviceConfig struct {
	target *url.URL
}

// proxyHandler returns a Gin handler that reverse-proxies the request to
// the given upstream target, rewriting the path so that the /api/<prefix>
// segment is stripped before forwarding.
func proxyHandler(cfg serviceConfig, stripPrefix string) gin.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(cfg.target)

	// Customise the director so the upstream only sees the path after the
	// strip prefix (e.g. /api/genres/graph → /genres/graph).
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		// Strip the gateway prefix so the upstream receives only its own path
		// segment (e.g. /api/genres/graph → /genres/graph).
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
	// Read upstream addresses from environment variables, falling back to
	// the default ports documented in the README.
	genreServiceURL := getEnv("GENRE_SERVICE_URL", "http://localhost:3001")
	songServiceURL := getEnv("SONG_SERVICE_URL", "http://localhost:3002")
	artistServiceURL := getEnv("ARTIST_SERVICE_URL", "http://localhost:3003")
	port := getEnv("PORT", "8080")

	genreTarget := mustParseURL(genreServiceURL)
	songTarget := mustParseURL(songServiceURL)
	artistTarget := mustParseURL(artistServiceURL)

	r := gin.Default()

	// Health check – useful for container orchestration readiness probes.
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Route groups: each group is proxied to its respective upstream service.
	// Only standard RESTful HTTP methods are allowed.
	allowedMethods := []string{
		http.MethodGet, http.MethodPost, http.MethodPut,
		http.MethodPatch, http.MethodDelete,
	}
	routes := []struct {
		prefix  string
		handler gin.HandlerFunc
	}{
		{"/api/genres/*path", proxyHandler(serviceConfig{target: genreTarget}, "/api/genres")},
		{"/api/songs/*path", proxyHandler(serviceConfig{target: songTarget}, "/api/songs")},
		{"/api/artists/*path", proxyHandler(serviceConfig{target: artistTarget}, "/api/artists")},
	}
	for _, route := range routes {
		for _, method := range allowedMethods {
			r.Handle(method, route.prefix, route.handler)
		}
	}

	log.Printf("API Gateway listening on :%s", port)
	log.Printf("  /api/genres  → %s", genreServiceURL)
	log.Printf("  /api/songs   → %s", songServiceURL)
	log.Printf("  /api/artists → %s", artistServiceURL)

	if err := r.Run(":" + port); err != nil {
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

// mustParseURL parses rawURL and exits if it is invalid.
func mustParseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		log.Fatalf("invalid URL %q: %v", rawURL, err)
	}
	return u
}
