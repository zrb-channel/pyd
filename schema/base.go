package pyd

import json "github.com/json-iterator/go"

type (
	Config struct {
		AppId  string
		AppKey string
	}

	BaseResponse struct {
		Code  int             `json:"code"`
		Msg   string          `json:"msg"`
		Data  json.RawMessage `json:"data"`
		ReqId string          `json:"req_id"`
	}

	ErrorResponse struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}
)
