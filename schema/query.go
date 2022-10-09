package pyd

import (
	json "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
)

type (
	// QueryResult 返回订单
	// example
	// {
	//	"id": "27998",
	//	"inputstatus": "10",
	//	"inputname": "未完成",
	//	"credit_money": "0.00",
	//	"credit_time": "",
	//	"draw_money": "0.00",
	//	"draw_time": ""
	//}
	QueryResult struct {
		// 进件记录id
		ID string `json:"id"`

		// 进件状态码
		Status string `json:"inputstatus"`

		// 进件状态
		StatusText string `json:"inputname"`

		// 授信额度
		CreditMoney decimal.Decimal `json:"credit_money"`

		// 授信时间 格式：2021-10-27 14:51:47
		CreditTime string `json:"credit_time"`

		// 提款额度
		DrawMoney decimal.Decimal `json:"draw_money"`

		// 提款时间 格式：2021-10-27 14:51:47
		DrawTime string `json:"draw_time"`

		// 拒绝原因
		Reason string `json:"reason"`
	}

	QueryResponseData struct {
		ErrorResponse
		Res QueryResult `json:"res"`
	}
)

func (s *QueryResponseData) String() string {
	v, _ := json.Marshal(s)
	return string(v)
}
