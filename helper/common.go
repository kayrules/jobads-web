package helper

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"../model"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// HTTPGet function
func HTTPGet(url string) (rsp []byte, err error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	response, err := netClient.Get(url)
	if err != nil {
		return nil, err

	}
	rsp, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return rsp, err
}

// CreateSession function
func CreateSession(c echo.Context, customer *model.Customer) error {
	sess, err := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	log.Println(customer)
	sess.Values["customer_id"] = string(customer.ID)
	sess.Values["customer_name"] = string(customer.Name)
	sess.Save(c.Request(), c.Response())

	return err
}

// DeleteSession function
func DeleteSession(c echo.Context) error {
	sess, err := session.Get("session", c)
	delete(sess.Values, "customer_id")
	delete(sess.Values, "customer_name")
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	sess.Save(c.Request(), c.Response())
	return err
}
