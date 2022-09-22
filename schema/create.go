package pyd

import (
	json "github.com/json-iterator/go"
)

type (
	ApplyResponseData struct {
		ErrorResponse
		ID  json.Number `json:"id"`
		Url string      `json:"url"`
	}

	ApplyRequest struct {
		IdCardImageF string
		IdCardImageB string
		Name         string
		Mobile       string
		IdCard       string
		CompanyName  string
		CompanyRegNo string
		PID          string
	}
)

func (s *ApplyRequest) String() string {
	v, _ := json.Marshal(s)
	return string(v)
}
