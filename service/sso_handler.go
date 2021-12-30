package service

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"
	"github.com/thanhlam/home-sso-service/model"
)

func ParseSSOToken(c echo.Context) error {
	authenRequestBody := new(model.AuthenRequestBody)
	err := c.Bind(authenRequestBody)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]interface{}{"code": "6", "message": "Body is Invalid", "data": nil})
	}
	token := authenRequestBody.Token
	authResponse, err := BasicAuth(token)
	//
	if err != nil {
		return c.JSON(200, map[string]interface{}{"code": "10", "message": "Connection refused", "data": nil})
	}
	//
	var result map[string]interface{}
	json.Unmarshal([]byte(authResponse), &result)

	if result["error"] != nil {
		if result["message"] != nil {
			return c.JSON(200, map[string]interface{}{"code": "7", "message": (result["message"]).(string), "data": nil})
		}
		return c.JSON(200, map[string]interface{}{"code": "5", "message": (result["error"]).(string), "data": nil})
	}
	//test
	attributes := result["attributes"].(map[string]interface{})
	attributes["username"] = result["id"]
	//return c.JSON(200, map[string]interface{}{"code": "0", "message": "success", "data": map[string]interface{}{"info": attributes}})
	return c.JSON(200, map[string]interface{}{"code": "0", "message": "success", "data": attributes})
}
func RequestSSOTokenv2(c echo.Context) error {
	tokenRequestBodyv2 := new(model.TokenRequestBodyv2)
	err := c.Bind(tokenRequestBodyv2)
	if err != nil {
		fmt.Println(err)
		return c.JSON(400, map[string]interface{}{"code": "6", "message": "Body is Invalid", "data": map[string]interface{}{"token": nil, "refresh_token": nil}})
	}
	username := tokenRequestBodyv2.Username
	password := tokenRequestBodyv2.Password
	clientID := tokenRequestBodyv2.Clientid
	clientSecret := tokenRequestBodyv2.Clientsecret
	_, userstatus := CheckUserStatus(username)
	if userstatus != "ACTIVE" {
		return c.JSON(200, map[string]interface{}{"code": "8", "message": "USER IS DISABLE", "data": map[string]interface{}{"token": nil, "refresh_token": nil}})
	}
	token, refreshtoken := RequestTokenv2(clientID, clientSecret, username, password)
	if token == "" {
		return c.JSON(200, map[string]interface{}{"code": "8", "message": "Authen Failed", "data": map[string]interface{}{"token": nil, "refresh_token": nil}})
	}
	return c.JSON(200, map[string]interface{}{"code": "0", "message": "success", "data": map[string]interface{}{"token": token, "refresh_token": refreshtoken}})
}
