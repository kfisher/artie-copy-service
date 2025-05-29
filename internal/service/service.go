// Copyright 2025 Kevin Fisher
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
// may be used to endorse or promote products derived from this software without
// specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

// Package service provides the service's server.
package service

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kfisher/artie-copy-service/internal/cfg"
	"github.com/kfisher/artie-copy-service/internal/models"
)

// Run configures the routes and starts the HTTP server.
func Run() error {

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "welcome to the world of endless wonder\n")
	})

	r.HandleFunc("/status", getStatus).Methods("GET")

	r.HandleFunc("/copy/start", startCopy).Methods("POST")
	r.HandleFunc("/copy/cancel", cancelCopy).Methods("POST")

	r.HandleFunc("/reset", reset).Methods("POST")

	r.HandleFunc("/copy-operations", getCopyOperationList).Methods("GET")
	r.HandleFunc("/copy-operations/{id}", getCopyOperation).Methods("GET")

	r.Use(loggingMiddleware)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port)
	return http.ListenAndServe(addr, r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("received request", "method", r.Method, "url", r.URL)
		next.ServeHTTP(w, r)
	})
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	// TODO[HIGH]: Need a proper implementation for this. Still not sure
	// where all this information will be stored or fetched.
	status := models.OpticalDrive{
		Id:           0,
		Name:         cfg.Device.Name,
		Host:         "",
		DeviceName:   "",
		SerialNumber: cfg.Device.Serial,
		State:        models.DriveStateIdle,
		DiscLabel:    "",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(status)
}

func startCopy(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented yet!", 500)
}

func cancelCopy(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented yet!", 500)
}

func reset(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented yet!", 500)
}

func getCopyOperationList(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented yet!", 500)
}

func getCopyOperation(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented yet!", 500)
}
