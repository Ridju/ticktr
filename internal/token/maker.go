package token

import "time"

type Maker interface {
	CreateToken(userID64 int64, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}
