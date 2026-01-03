package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github-activity/internal/models"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type Http interface {
	Do(req *http.Request) (*http.Response, error)
}

type GitHubActivityService struct {
	baseURLTemplate string
	client          Http
}

func NewGitHubActivityService(baseURLTemplate string, client Http) (*GitHubActivityService, error) {
	if strings.TrimSpace(baseURLTemplate) == "" {
		return nil, fmt.Errorf("base URL template is empty")
	}

	testURL := strings.Replace(baseURLTemplate, "{NAME}", "test", 1)
	if _, err := url.Parse(testURL); err != nil {
		return nil, fmt.Errorf("invalid base URL template: %w", err)
	}

	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}

	return &GitHubActivityService{
		baseURLTemplate: baseURLTemplate,
		client:          client,
	}, nil
}

func (s *GitHubActivityService) BuildCallUrl(name string) (string, error) {
	rawUrl := os.Getenv("URL")
	rawUrl = strings.Replace(rawUrl, "{NAME}", url.QueryEscape(name), 1)
	parsed, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}
	return parsed.String(), nil
}

func (s *GitHubActivityService) FetchData(ctx context.Context, name string) ([]models.Activity, error) {
	url, err := s.BuildCallUrl(name)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("user not found. please check the username")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data: %d, %s", resp.StatusCode, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var activities []models.Activity
	if err := json.Unmarshal(body, &activities); err != nil {
		return nil, err
	}

	return activities, nil
}
func DescribeActivity(event models.Activity) string {
	repo := event.Repo.Name
	if strings.TrimSpace(repo) == "" {
		repo = "unknown repo"
	}
	switch event.Type {
	case "PushEvent":
		if event.Payload.Ref != "" {
			return fmt.Sprintf("Pushed to %s (%s)", repo, event.Payload.Ref)
		}
		return fmt.Sprintf("Pushed to %s", repo)

	case "CreateEvent":
		if event.Payload.Ref != "" {
			return fmt.Sprintf("Created %s in %s", event.Payload.Ref, repo)
		}
		return fmt.Sprintf("Created something in %s", repo)

	case "ForkEvent":
		return fmt.Sprintf("Forked %s", repo)

	case "WatchEvent":
		return fmt.Sprintf("Starred %s", repo)

	case "IssuesEvent":
		return fmt.Sprintf("Issue activity in %s", repo)

	default:
		return fmt.Sprintf("%s in %s", event.Type, repo)
	}
}
