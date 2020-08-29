package access_token

import (
	"github.com/JingdaMai/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const expirationTime = 24

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (at *AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid access token ID")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user ID")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client ID")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
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
