package data

import (
	"time"
)

type Url struct {
	ShortUrl  string    `json:"shorturl"`
	ShortId   string    `json:"shortid"`
	LongUrl   string    `json:"longurl"`
	Clicks    int       `json:"clicks"`
	CreatedAt time.Time `json:"createdat"`
	Owner     string    `json:"owner"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
