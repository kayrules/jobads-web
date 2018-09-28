package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"../config"
	"../helper"
	"../model"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

// Cart controller
func Cart(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil || sess.Values["customer_name"] == nil {
		return c.Redirect(http.StatusMovedPermanently, "/login")
	}

	// fetch pricing rules from api
	apiURL := config.APIURL + config.GetPricingRulesURL + "/" + sess.Values["customer_id"].(string)
	result, err := helper.HTTPGet(apiURL)
	if err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}
	response := new([]model.PricingRules)
	err = json.Unmarshal(result, &response)
	if err != nil {
		return c.HTML(http.StatusBadRequest, err.Error())
	}

	calculateURL := config.APIURL + config.GetCalculateURL + "/" + sess.Values["customer_id"].(string)

	log.Println(apiURL, "response", response)

	params := map[string]interface{}{
		"customer_id":   sess.Values["customer_id"],
		"customer_name": strings.ToUpper(sess.Values["customer_name"].(string)),
		"post_url":      calculateURL,
	}
	return c.Render(http.StatusOK, "cart.html", params)
}
