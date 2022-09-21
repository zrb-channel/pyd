package schema

import (
	json "github.com/json-iterator/go"
)

type (
	CreateResponseData struct {
		ErrorResponse
		ID  json.Number `json:"id"`
		Url string      `json:"url"`
	}

	CreateRequest struct {
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

func (s *CreateRequest) String() string {
	v, _ := json.Marshal(s)
	return string(v)
}
