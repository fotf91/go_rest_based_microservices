package domain

import (
	"banking/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct {
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {

	u := buildVerifyURL(token, routeName, vars)

	/**
	example of the value of u:
	u = http://localhost:8181/auth/verify?customer_id=2001&routeName=GetCustomer&token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjdXN0b21lcl9pZCI6IjIwMDEiLCJhY2NvdW50cyI6WyI5NTQ3MiIsIjk1NDczIiwiOTU0NzQiXSwidXNlcm5hbWUiOiIyMDAxIiwicm9sZSI6InVzZXIiLCJleHAiOjE2NDA5NzQ2NzN9.ywugzrYcQQXT3MuFH2mx5ZFNNey1sn9teucpC9npkNg
	---
	when visiting this endpoint I get {"isAuthorized":true}
	*/

	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server:" + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

/*
  This will generate a url for token verification in the below format

  /auth/verify?token={token string}
              &routeName={current route name}
              &customer_id={customer id from the current route}
              &account_id={account id from current route if available}

  Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&customer_id=2000&account_id=95470
*/
func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	address := os.Getenv("AUTH_MIDDLEWARE_SERVER_ADDRESS")
	port := os.Getenv("AUTH_MIDDLEWARE_SERVER_PORT")
	u := url.URL{Host: address + ":" + port, Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}

func NewAuthRepository() RemoteAuthRepository {
	return RemoteAuthRepository{}
}
