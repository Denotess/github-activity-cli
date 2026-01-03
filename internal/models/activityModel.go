package models

import (
	"time"
)

type Actor struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Payload struct {
	RepositoryId int    `json:"repository_id"`
	PushId       int64  `json:"push_id"`
	Ref          string `json:"ref"`
	Head         string `json:"head"`
	Before       string `json:"before"`
}

type Activity struct {
	Id        string    `json:"id"`
	Type      string    `json:"type"`
	Actor     Actor     `json:"actor"`
	Repo      Repo      `json:"repo"`
	Payload   Payload   `json:"payload"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
}
