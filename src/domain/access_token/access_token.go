package access_token

import "time"

const expirationTime = 24

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int    `json:"user_id"`
	ClientId    int    `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() *AccessToken {
	return &AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

// Web frontend - Client-Id: 123
// Android App - Client-Id: 234
