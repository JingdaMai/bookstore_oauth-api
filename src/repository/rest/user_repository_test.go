package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://fake/users/login",
		ReqBody:      `{"email":"email@gamil.com","password":"the-password"}`,
		RespHTTPCode: -1,
		RespBody:     `{}`,
	})

	repository := usersRepository{}
	user, err := repository.LoginUser("email@gamil.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to login user", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost/users/login",
		ReqBody:      `{"email":"email@gamil.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login credentials", "status":"404", "error":"not_found"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gamil.com", "the-password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost/users/login",
		ReqBody:      `{"email":"email@gamil.com","password":"the-password"}`,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message":"invalid login credentials", "status":404, "error":"not_found"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gamil.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "invalid login credentials", err.Message)
}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost/users/login",
		ReqBody:      `{"email":"email@gamil.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": "1","first_name": "jingda","last_name": "mai","email": "a@com","date_created": "2020-08-23T08:02:34Z","status": "active"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gamil.com", "password")

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal users login response", err.Message)
}

func TestLoginUserNoError(t *testing.T) {
	rest.FlushMockups()
	_ = rest.AddMockups(&rest.Mock{
		HTTPMethod:   http.MethodPost,
		URL:          "http://localhost/users/login",
		ReqBody:      `{"email":"email@gamil.com","password":"the-password"}`,
		RespHTTPCode: http.StatusOK,
		RespBody:     `{"id": 1,"first_name": "jingda","last_name": "mai","email": "a@com","date_created": "2020-08-23T08:02:34Z","status": "active"}`,
	})

	repository := usersRepository{}

	user, err := repository.LoginUser("email@gamil.com", "password")

	assert.NotNil(t, user)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, user.Id)
	assert.EqualValues(t, "jingda", user.FirstName)
	assert.EqualValues(t, "mai", user.LastName)
	assert.EqualValues(t, "a@com", user.Email)
}
