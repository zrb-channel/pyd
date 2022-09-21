package schema

import json "github.com/json-iterator/go"

type (
	QueryResult struct {
		ID          string `json:"id"`           // 进件记录id
		Status      string `json:"inputstatus"`  // 进件状态码
		StatusText  string `json:"inputname"`    // 进件状态
		CreditMoney string `json:"credit_money"` // 授信额度
		CreditTime  string `json:"credit_time"`  // 授信时间
		DrawMoney   string `json:"draw_money"`   // 提款额度
		DrawTime    string `json:"draw_time"`    // 提款时间
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
