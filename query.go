package pyd

import (
	"context"
	"errors"
	log "github.com/zrb-channel/utils/logger"
	"net/http"
	"strconv"
	"time"

	json "github.com/json-iterator/go"
	config "github.com/zrb-channel/pyd/config"
	schema "github.com/zrb-channel/pyd/schema"
	"github.com/zrb-channel/utils"
	"github.com/zrb-channel/utils/hash"
)

// Query
// @param ctx
// @param conf
// @date 2022-09-24 01:25:49
func Query(ctx context.Context, conf *schema.Config, companyName string) (*schema.QueryResult, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	data := map[string]string{
		"mall_id":      conf.AppId,
		"company_name": companyName,
		"tm":           strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
	}

	data["sign"] = hash.MD5String(data["mall_id"] + data["company_name"] + data["tm"] + conf.AppKey)
	resp, err := utils.Request(ctx).
		SetFormData(data).Post(config.BaseURL + "/loans/open/inputstatus")
	if err != nil {
		log.WithError(err).Error("[浦逸贷]-订单查询请求失败", log.Fields(map[string]any{"order": data}))
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		log.Error("[浦逸贷]-订单查询，返回statusCode有误", log.Fields(map[string]any{"order": data}))
		return nil, errors.New(resp.String())
	}

	result := &schema.BaseResponse{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		log.Error("[浦逸贷]-订单查询，返回数据解析失败", log.Fields(map[string]any{"order": data, "body": resp.String()}))
		return nil, err
	}

	res := &schema.QueryResponseData{}
	if err = json.Unmarshal(result.Data, res); err != nil {
		log.Error("[浦逸贷]-订单查询，返回数据解析失败2", log.Fields(map[string]any{"order": data, "body": resp.String()}))
		return nil, err
	}

	if result.Code != http.StatusOK {
		log.Error("[浦逸贷]-订单查询，返回数据code有误", log.Fields(map[string]any{"order": data, "body": resp.String(), "result": result}))
		return nil, errors.New(result.Msg)
	}

	if res.Status != 1000 {
		log.Error("[浦逸贷]-订单查询，返回数据status有误", log.Fields(map[string]any{"order": data, "body": resp.String(), "result": res}))

		return nil, errors.New(res.Message)
	}

	return &res.Res, nil
}
