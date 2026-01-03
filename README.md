# GitHub Activity CLI

CLI tool that fetches a user’s public GitHub events and prints a readable summary. Push events are grouped per repo (e.g., `Pushed 3 commits to owner/repo`) instead of showing each commit separately.

## Requirements
- Go 1.20+
- GitHub’s public Events API access (unauthenticated; subject to IP rate limits)

## Setup
1) Clone the repo.
2) Copy `.env.example` to `.env` and set the base URL (default works for public events):
   ```
   URL=https://api.github.com/users/{NAME}/events
   ```
3) Install deps:
   ```
   go mod download
   ```

## Usage
- Run from source:
  ```
  go run . github-activity --username <github_user>
  ```
- Example:
  ```
  go run . github-activity --username torvalds
  ```

## Behavior
- Push events are aggregated per repository and show total commits pushed.
- Other events (create, fork, watch/star, issues, etc.) are printed as they appear.

[Project Idea Link](https://roadmap.sh/projects/github-user-activity)