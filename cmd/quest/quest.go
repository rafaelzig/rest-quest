package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rafaelzig/rest-quest/internal/app/quest"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const serverPortEnvKey = "SERVER_PORT"
const defaultServerPort = uint16(8080)

func main() {
	h := createHandler()
	srv := createServer(h)
	shutdownChan := make(chan struct{})
	go handleGracefulShutdown(h, srv, shutdownChan)
	h.Info(fmt.Sprintf("Starting HTTP server on http://localhost%s", srv.Addr))
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		h.Fatal(fmt.Sprintf("HTTP server ListenAndServe: %v", err))
		os.Exit(1)
	}
	<-shutdownChan
	h.Info("HTTP server gracefully Shutdown")
}

func createHandler() *quest.Server {
	h := &quest.Server{
		Router:         mux.NewRouter(),
		LogHandlerFunc: handleLogEvent(),
	}
	h.Routes()
	return h
}

func handleLogEvent() func(v interface{}) {
	logger := log.New(os.Stdout, "", 0)
	return func(v interface{}) {
		res, err := json.Marshal(v)
		if err != nil {
			logger.Printf("handleLogEvent failed: %s\n", err)
			return
		}
		logger.Println(string(res))
	}
}

func createServer(h *quest.Server) *http.Server {
	server := &http.Server{
		Addr:         ":" + strconv.FormatUint(uint64(parseServerPort(os.Getenv(serverPortEnvKey))), 10),
		Handler:      h,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}
	return server
}

func handleGracefulShutdown(h *quest.Server, srv *http.Server, shutdownChan chan struct{}) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGTERM, syscall.SIGKILL)
	<-interruptChan
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	h.Info("Initiating HTTP server Shutdown")
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		h.Error(fmt.Sprintf("HTTP server Shutdown: %v", err))
	}
	close(shutdownChan)
}

func parseServerPort(str string) uint16 {
	if len(str) == 0 {
		return defaultServerPort
	}
	serverPort, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return defaultServerPort
	}

	return uint16(serverPort)
}
