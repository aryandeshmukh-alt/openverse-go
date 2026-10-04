package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

type WebsiteRequest struct {
	Websites []string `json:"websites"`
}

type StatusChecker interface {
	Check(ctx context.Context, name string) (bool, error)
}

type httpChecker struct {
	client *http.Client
}

func (h httpChecker) Check(ctx context.Context, name string) (bool, error) {
	url := name

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)

	if err != nil {
		return false, err
	}

	response, err := h.client.Do(request)

	if err != nil {
		return false, err
	}

	defer response.Body.Close()

	return response.StatusCode == http.StatusOK, nil
}

type WebsiteStatus struct {
	Status string `json:"status"`
}

var websites = make(map[string]string)
var websitesMutex sync.RWMutex
var checker StatusChecker = httpChecker{
	client: &http.Client{
		Timeout: 10 * time.Second,
	},
}

func addWebsites(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request WebsiteRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	websitesMutex.Lock()

	for _, website := range request.Websites {
		website = strings.TrimSpace(website)
		if website != "" {
			websites[website] = "DOWN"
		}
	}

	websitesMutex.Unlock()

	w.WriteHeader(http.StatusOK)
}

func getWebsites(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")

	w.Header().Set("Content-Type", "application/json")

	if name != "" {

		websitesMutex.RLock()

		status, exists := websites[name]

		websitesMutex.RUnlock()

		if !exists {
			http.Error(w, "Website not found", http.StatusNotFound)
			return
		}

		response := map[string]string{
			name: status,
		}

		json.NewEncoder(w).Encode(response)

		return
	}

	websitesMutex.RLock()

	result := make(map[string]string)

	for website, status := range websites {
		result[website] = status
	}

	websitesMutex.RUnlock()

	json.NewEncoder(w).Encode(result)
}

func checkWebsites() {

	websitesMutex.RLock()

	websiteList := make([]string, 0, len(websites))

	for website := range websites {
		websiteList = append(websiteList, website)
	}

	websitesMutex.RUnlock()

	var wg sync.WaitGroup

	for _, website := range websiteList {

		wg.Add(1)

		go func(name string) {

			defer wg.Done()

			ctx, cancel := context.WithTimeout(
				context.Background(),
				10*time.Second,
			)

			defer cancel()

			isUp, err := checker.Check(ctx, name)

			status := "DOWN"

			if err == nil && isUp {
				status = "UP"
			}

			websitesMutex.Lock()

			if _, exists := websites[name]; exists {
				websites[name] = status
			}

			websitesMutex.Unlock()

		}(website)
	}

	wg.Wait()
}

func monitorWebsites() {

	for {

		fmt.Println("Checking website statuses...")

		checkWebsites()

		fmt.Println("Website status check completed.")

		time.Sleep(1 * time.Minute)
	}
}

func main() {

	go monitorWebsites()

	http.HandleFunc("/websites", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			addWebsites(w, r)

		case http.MethodGet:
			getWebsites(w, r)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server started on port 3000")

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Server error:", err)
	}
}