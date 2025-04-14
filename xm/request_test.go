package xm

import (
	"testing"
	"time"
)

func Test_GetRequest(t *testing.T) {
	url := "https://api.juejin.cn/user_api/v1/user/profile_id"
	type GetRequest struct {
		Aid    int    `json:"aid"`
		UUID   string `json:"uuid"`
		Spider int    `json:"spider"`
		WebId  int64  `json:"web_id"`
	}
	getReq := GetRequest{
		Aid:    2608,
		UUID:   "7419192284367259177",
		Spider: 0,
		WebId:  7419192284367259177,
	}

	jsonStr, err := StructToQueryParams(getReq, false)
	if err != nil {
		t.Error("StructToQueryParams error:" + err.Error())
		return
	}
	t.Log("URL:" + url + "?" + jsonStr.Encode())
	respBody, err := GetRequestByOrigin(url, nil, jsonStr.Encode(), 30*time.Second)
	if err != nil {
		t.Error("GetRequestByOrigin error:" + err.Error())
		return
	}
	t.Log("Response:\n" + string(respBody))
}
