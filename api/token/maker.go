package token

import "time"

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, error)

	VertifyToken(token string) (*Payload, error)
}
