package project

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
)

// BudgetGroupList 获取预算组列表
func BudgetGroupList(clt *core.SDKClient, accessToken string, req *project.BudgetGroupListRequest) (*project.BudgetGroupListResult, error) {
	var resp project.BudgetGroupListResponse
	if err := clt.GetAPI("v3.0/budget_group/list/", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
