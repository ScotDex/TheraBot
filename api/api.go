package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Health Check
type ApiHealth struct {
	Version string `json:"api_version"`
}

func CheckApiHealth() (*ApiHealth, error) {
	url := "https://api.eve-scout.com/v2/health"
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "WingspanTherabot")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var healthStatus ApiHealth
	if err := json.Unmarshal(body, &healthStatus); err != nil {
		return nil, err
	}

	return &healthStatus, nil
}

// Route represents a route fetched from the EVE Scout API.
type Route struct {
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	CreatedByID     int       `json:"created_by_id"`
	CreatedByName   string    `json:"created_by_name"`
	UpdatedAt       time.Time `json:"updated_at"`
	UpdatedByID     int       `json:"updated_by_id"`
	UpdatedByName   string    `json:"updated_by_name"`
	CompletedAt     time.Time `json:"completed_at"`
	CompletedByID   int       `json:"completed_by_id"`
	CompletedByName string    `json:"completed_by_name"`
	Completed       bool      `json:"completed"`
	WhExitsOutward  bool      `json:"wh_exits_outward"`
	WhType          string    `json:"wh_type"`
	MaxShipSize     string    `json:"max_ship_size"`
	ExpiresAt       time.Time `json:"expires_at"`
	RemainingHours  int       `json:"remaining_hours"`
	SignatureType   string    `json:"signature_type"`
	OutSystemID     int       `json:"out_system_id"`
	OutSystemName   string    `json:"out_system_name"`
	OutSignature    string    `json:"out_signature"`
	InSystemID      int       `json:"in_system_id"`
	InSystemClass   string    `json:"in_system_class"`
	InSystemName    string    `json:"in_system_name"`
	InRegionID      int       `json:"in_region_id"`
	InRegionName    string    `json:"in_region_name"`
	InSignature     string    `json:"in_signature"`
	Comment         string    `json:"comment"`
}

// API Calls to fetch routes from specific systems

func TurnurRoutes() ([]Route, error) {
	resp, err := http.Get("https://api.eve-scout.com/v2/public/signatures?system_name=turnur")
	if err != nil {
		return nil, fmt.Errorf("failed to make request to EVE Scout: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("EVE Scout API returned non-200 status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var routes []Route
	if err := json.Unmarshal(body, &routes); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return routes, nil
}

func TheraRoutes() ([]Route, error) {
	resp, err := http.Get("https://api.eve-scout.com/v2/public/signatures?system_name=thera")
	if err != nil {
		return nil, fmt.Errorf("failed to make request to EVE Scout: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("EVE Scout API returned non-200 status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var routes []Route
	if err := json.Unmarshal(body, &routes); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return routes, nil
}

func AllRoutes() ([]Route, error) {
	resp, err := http.Get("https://api.eve-scout.com/v2/public/signatures")
	if err != nil {
		return nil, fmt.Errorf("failed to make request to EVE Scout: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("EVE Scout API returned non-200 status: %d", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var routes []Route
	if err := json.Unmarshal(body, &routes); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return routes, nil
}
