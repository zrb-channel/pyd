package pyd

import (
	"context"
	"errors"
	config "github.com/zrb-channel/pyd/config"
	"net/http"
	"strconv"
	"time"

	"github.com/zrb-channel/utils"

	"github.com/zrb-channel/pyd/schema"
	"github.com/zrb-channel/utils/hash"

	"go.uber.org/zap"

	log "github.com/zrb-channel/utils/logger"

	json "github.com/json-iterator/go"
)

func Apply(ctx context.Context, conf *pyd.Config, req *pyd.ApplyRequest) (*pyd.ApplyResponseData, error) {

	var data = map[string]string{
		"mall_id":        conf.AppId,
		"name":           req.Name,
		"mobile":         req.Mobile,
		"idcard":         req.IdCard,
		"company_name":   req.CompanyName,
		"company_reg_no": req.CompanyRegNo,
		"pid":            req.PID,
		"tm":             strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
	}

	signStr := data["mall_id"] + data["name"] + data["mobile"] + data["idcard"] + data["company_name"] + data["company_reg_no"] + data["pid"] + data["tm"] + conf.AppKey
	data["sign"] = hash.MD5String(signStr)

	resp, err := utils.Request(ctx).
		SetFormData(data).
		SetFile("idcard_img_f", req.IdCardImageF).
		SetFile("idcard_img_b", req.IdCardImageB).
		SetHeader("Content-Type", "multipart/form-data").
		Post(config.BaseURL + "/loans/open/setdatapf")

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New(resp.Status())
	}

	f := map[string]any{"body": resp.String(), "data": data, "idCardImage": map[string]string{"idcard_img_f": req.IdCardImageF, "idcard_img_b": req.IdCardImageB}}
	result := &pyd.BaseResponse{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.WithError(err).Error("浦逸贷提交失败", zap.Any("data", f))
		return nil, errors.New("提交失败")
	}

	if result.Code != http.StatusOK {

		return nil, errors.New(result.Msg)
	}

	res := &pyd.ApplyResponseData{}
	if err = json.Unmarshal(result.Data, res); err != nil {
		log.WithError(err).Error("浦逸贷提交失败", zap.Any("data", f))
		return nil, errors.New("提交失败")
	}

	fields := map[string]any{
		"body":    resp.String(),
		"data":    data,
		"status":  res.Status,
		"message": res.Message,
	}

	if res.Status != 1000 {
		log.Error("浦逸贷提交失败2 ", zap.Any("data", fields))
		return nil, errors.New(res.Message)
	}

	return res, nil
}
