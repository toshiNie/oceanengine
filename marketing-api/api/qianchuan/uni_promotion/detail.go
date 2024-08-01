package unipromotion

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	unipromotion "github.com/bububa/oceanengine/marketing-api/model/qianchuan/uni_promotion"
)

// Detail 获取全域推广计划详情
func Detail(clt *core.SDKClient, accessToken string, req *unipromotion.DetailRequest) (*unipromotion.AdDetail, error) {
	var resp unipromotion.DetailResponse
	if err := clt.Get("v1.0/qianchuan/uni_promotion/ad/detail/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
