package sentry

import go_types "github.com/leapforce-libraries/go_types"

type User struct {
	Id        go_types.Int64String `json:"id"`
	Username  *string              `json:"username"`
	Name      *string              `json:"name"`
	IpAddress *string              `json:"ip_address"`
	Email     *string              `json:"email"`
	Data      *struct {
		IsStaff bool `json:"isStaff"`
	} `json:"data"`
}
