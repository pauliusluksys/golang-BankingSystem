package domain

import (
	"bankingV2/logger"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"net/url"
)

type AuthRepositoryDb struct {
	client *sqlx.DB
}
type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
	FindBy(username, password string) (*Login, error)
}
type RemoteAuthRepository struct {
	client *sqlx.DB
}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	u := buildVerifyURL(token, routeName, vars)

	if response, err := http.Get(u); err != nil {
		fmt.Println("Error while sending..." + err.Error())
		return false
	} else {
		//m := make(map[string]bool)
		var isAuthorized bool
		if err = json.NewDecoder(response.Body).Decode(&isAuthorized); err != nil {
			logger.Error("Error while decoding response from auth server:" + err.Error())
			return false
		}
		log.Println(isAuthorized)

		return isAuthorized
	}
}

func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8181", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	log.Println(u.String())
	return u.String()
}
func (d RemoteAuthRepository) FindBy(username, password string) (*Login, error) {
	var login Login
	sqlVerify := `SELECT username, u.customer_id, role, group_concat(a.account_id) as account_numbers from users u
					LEFT JOIN accounts a ON a.customer_id = u.customer_id
					WHERE username = ? and password = ?
					group by u.customer_id;`
	err := d.client.Get(&login, sqlVerify, username, password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		} else {
			log.Println("Error while verifying login request from database: " + err.Error())
			return nil, errors.New("unexpected database error")
		}
	}
	return &login, nil
}

func NewAuthRepository(dbClient *sqlx.DB) RemoteAuthRepository {

	return RemoteAuthRepository{dbClient}
}
