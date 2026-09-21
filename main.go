package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"leadphone-validator/internal/handlers"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3007"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.HandleHealth)
	mux.HandleFunc("/v1/health", handlers.HandleHealth)
	mux.HandleFunc("/validate", handlers.HandleValidate)
	mux.HandleFunc("/v1/validate", handlers.HandleValidate)
	mux.HandleFunc("/validate/batch", handlers.HandleBatchValidate)
	mux.HandleFunc("/v1/validate/batch", handlers.HandleBatchValidate)

	mux.HandleFunc("/ddd", handlers.HandleDDD)
	mux.HandleFunc("/v1/ddd", handlers.HandleDDD)
	mux.HandleFunc("/ddd/", handlers.HandleDDD)
	mux.HandleFunc("/v1/ddd/", handlers.HandleDDD)

	mux.HandleFunc("/ddi", handlers.HandleDDI)
	mux.HandleFunc("/v1/ddi", handlers.HandleDDI)
	mux.HandleFunc("/ddi/", handlers.HandleDDI)
	mux.HandleFunc("/v1/ddi/", handlers.HandleDDI)

	mux.HandleFunc("/sanitize/ninth-digit", handlers.HandleNinthDigit)
	mux.HandleFunc("/v1/sanitize/ninth-digit", handlers.HandleNinthDigit)

	mux.HandleFunc("/whatsapp/link", handlers.HandleWhatsAppLink)
	mux.HandleFunc("/v1/whatsapp/link", handlers.HandleWhatsAppLink)

	mux.HandleFunc("/check/patterns", handlers.HandlePatternCheck)
	mux.HandleFunc("/v1/check/patterns", handlers.HandlePatternCheck)

	mux.HandleFunc("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "openapi.yaml")
	})

	mux.HandleFunc("/documentation.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "documentation.yaml")
	})

	webDist := filepath.Join("web", "dist")
	if info, err := os.Stat(webDist); err == nil && info.IsDir() {
		fs := http.FileServer(http.Dir(webDist))
		mux.Handle("/", fs)
	} else {
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/" {
				handlers.HandleHealth(w, r)
				return
			}
			http.NotFound(w, r)
		})
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      corsMiddleware(mux),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Printf("LeadPhone Validator listening on port %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-stop
	log.Println("Shutting down LeadPhone Validator gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced shutdown: %v", err)
	}

	log.Println("LeadPhone Validator exited cleanly")
}
