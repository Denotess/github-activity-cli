package cmd

import (
	"context"
	"fmt"
	"github-activity/internal/services"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type options struct {
	username string
}

func NewRootCmd() *cobra.Command {
	opts := options{}
	cmd := &cobra.Command{
		Use:           "github-activity",
		Short:         "Github User Activity is a CLI tool for fetching user activity",
		Long:          "Github User Activity fetches and prints public activity for a GitHub user.",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDisplayActivityCmd(cmd.Context(), opts.username)
		},
	}

	cmd.Flags().StringVarP(&opts.username, "username", "u", "", "GitHub username whose activity should be displayed")
	_ = cmd.MarkFlagRequired("username")

	return cmd
}

func RunDisplayActivityCmd(ctx context.Context, username string) error {
	username = strings.TrimSpace(username)
	if username == "" {
		return fmt.Errorf("please provide a username with --username")
	}

	baseURL := strings.TrimSpace(os.Getenv("URL"))
	if baseURL == "" {
		return fmt.Errorf("URL environment variable is required")
	}

	service, err := services.NewGitHubActivityService(baseURL, http.DefaultClient)
	if err != nil {
		return err
	}

	if ctx == nil {
		ctx = context.Background()
	}
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	activities, err := service.FetchData(ctx, username)
	if err != nil {
		return err
	}

<<<<<<< HEAD
	for _, event := range activities {
		fmt.Println(services.DescribeActivity(event))
	}

=======
	pushTotals := make(map[string]int)
	pushOrder := make([]string, 0)

	for _, event := range activities {
		if event.Type == "PushEvent" {
			repo := strings.TrimSpace(event.Repo.Name)
			if repo == "" {
				repo = "unknown repo"
			}

			count := services.CommitCount(event)
			if count == 0 {
				count = 1
			}

			if _, exists := pushTotals[repo]; !exists {
				pushOrder = append(pushOrder, repo)
			}
			pushTotals[repo] += count
			continue
		}

		fmt.Println(services.DescribeActivity(event))
	}

	for _, repo := range pushOrder {
		count := pushTotals[repo]
		if count == 1 {
			fmt.Printf("Pushed 1 commit to %s\n", repo)
			continue
		}
		fmt.Printf("Pushed %d commits to %s\n", count, repo)
	}

>>>>>>> c1e49f9 (feat: added grouping of events)
	return nil
}
