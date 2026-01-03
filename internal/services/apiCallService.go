package services

import (
	"encoding/json"
	"fmt"
	"github-activity/internal/models"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func BuildCallUrl(name string) (string, error) {
	rawUrl := os.Getenv("URL")
	rawUrl = strings.Replace(rawUrl, "{NAME}", url.QueryEscape(name), 1)
	parsed, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	return parsed.String(), nil
}

func FetchData(name string) ([]models.Activity, error) {
	url, err := BuildCallUrl(name)
	if err != nil {
		return nil, err
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 404 {
		return nil, fmt.Errorf("user not found. please check the username")
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: %d", response.StatusCode)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var activities []models.Activity
	if err := json.Unmarshal(body, &activities); err != nil {
		return nil, err
	}

	return activities, nil
}
