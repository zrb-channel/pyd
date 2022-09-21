package pyd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	json "github.com/json-iterator/go"
	"github.com/zrb-channel/pyd/config"
	"github.com/zrb-channel/pyd/schema"
	"github.com/zrb-channel/utils"
	"github.com/zrb-channel/utils/hash"
)

func Query(ctx context.Context) error {
	data := map[string]string{
		"mall_id":      config.AppID,
		"company_name": "株洲瑞特建材销售有限公司",
		"tm":           strconv.FormatInt(time.Now().UnixNano()/1e6, 10),
	}

	data["sign"] = hash.MD5String(data["mall_id"] + data["company_name"] + data["tm"] + config.AppKey)
	resp, err := utils.Request(ctx).
		SetFormData(data).Post(config.BaseURL + "/loans/open/inputstatus")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New(resp.String())
	}

	result := &schema.BaseResponse{}
	if err = json.Unmarshal(resp.Body(), result); err != nil {
		return err
	}

	res := &schema.QueryResponseData{}
	if err = json.Unmarshal(result.Data, res); err != nil {
		return err
	}

	if result.Code != http.StatusOK {
		fmt.Println(result.Msg)
		return errors.New(result.Msg)
	}

	if res.Status != 1000 {
		fmt.Println(res.Message)
		return errors.New(res.Message)
	}

	fmt.Println(res.String())
	return nil
}
