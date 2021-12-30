package model

type AuthenRequestBody struct {
	Token string `json:"token"`
}
type ThingRoleBody struct {
	Thingid string `json:"thingid"`
}
type UserSSO struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Userid       string `json:"userid"`
	Usermail     string `json:"usermail"`
	Userstatus   string `json:"userstatus"`
	Userparentid string `json:"userparentid"`
	Usertype     string `json:"usertype"`
}
type UserRoleBody struct {
	Token   string `json:"token"`   //for check user status and assign thing (userid)
	Thingid string `json:"thingid"` //for check assign thing
	Role    string `json:"role"`    //check role
}

//request token
type TokenRequestBodyv2 struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	Clientid     string `json:"clientid"`
	Clientsecret string `json:"clientsecret"`
}
