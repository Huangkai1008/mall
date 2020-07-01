package ecode

import "github.com/tidwall/gjson"

type MallError struct {
	Message string `json:"message"`
}

func GetMsg(body string) string {
	msg := gjson.Get(body, "message")
	return msg.String()
}
