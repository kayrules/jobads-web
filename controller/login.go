package controller

import (
	"encoding/json"
	"net/http"

	"../config"
	"../helper"
	"../model"
	"github.com/labstack/echo"
)

// Home function
func Home(c echo.Context) (err error) {
	return c.Redirect(http.StatusMovedPermanently, "/login")
}

// Login function
func Login(c echo.Context) error {

	apiURL := config.APIURL + config.GetCustomersURL
	result, err := helper.HTTPGet(apiURL)
	if err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}

	response := new([]model.Customer)
	err = json.Unmarshal(result, &response)
	if err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}

	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"customers": response,
	})
}

// LoginPost controller
func LoginPost(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}

	// if not login
	if u.Username == "" {
		params := map[string]interface{}{
			"username": "",
			"message":  "Invalid credential",
		}
		return c.Render(http.StatusOK, "login.html", params)
	}

	// if logged in
	apiURL := config.APIURL + config.GetCustomersURL + "?name=" + u.Username
	result, err := helper.HTTPGet(apiURL)
	if err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}

	customer := new([]model.Customer)
	err = json.Unmarshal(result, &customer)
	if err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}

	tmpCustomer := &model.Customer{
		ID:   (*customer)[0].ID,
		Name: (*customer)[0].Name,
	}

	err = helper.CreateSession(c, tmpCustomer)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, "/cart")
}

// Logout function
func Logout(c echo.Context) error {

	err := helper.DeleteSession(c)
	if err != nil {
		return err
	}

	return c.Redirect(http.StatusMovedPermanently, "/login")
}
