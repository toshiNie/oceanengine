package agent

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QueryTransferDetailRequest 转账-查询转账单信息（代理） API Request
type QueryTransferDetailRequest struct {
	// BizRequestNo 请求id，推荐uuid，方便请求链路对齐
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// AgentID 代理商账户id
	AgentID uint64 `json:"agent_id,omitempty"`
	// TransferBizRequestNo 发起转账的幂等id
	TransferBizRequestNo string `json:"transfer_biz_request_no,omitempty"`
	// TransferSerial 转账单号，与transfer_biz_request_no两者传其一即可
	TransferSerial string `json:"transfer_serial,omitempty"`
}

// Encode implements GetRequest interface
func (r QueryTransferDetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("biz_request_no", r.BizRequestNo)
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	values.Set("transfer_biz_request_no", r.TransferBizRequestNo)
	values.Set("transfer_serial", r.TransferSerial)
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QueryTransferDetailResponse 转账-查询转账单信息（代理） API Response
type QueryTransferDetailResponse struct {
	model.BaseResponse
	Data *TransferDetail `json:"data,omitempty"`
}

// TransferDetail 转账详情
type TransferDetail struct {
	// TransferSerial 转账单号
	TransferSerial string `json:"transfer_serial,omitempty"`
	// BizRequestNo 幂等id
	BizRequestNo string `json:"biz_request_no,omitempty"`
	// TransferDirection 转账方向(以目标账户视角确定) 可选值:
	// TRANSFER_IN 转入
	// TRANSFER_OUT 转出
	TransferDirection enum.TransferDirection `json:"transfer_direction,omitempty"`
	// TransferAmount 转账总金额(单位：分)
	TransferAmount int64 `json:"transfer_amount,omitempty"`
	// TransferStatus 转账总状态 可选值:
	// NO_TRANSFER 未转账
	// TRANSFER_FAILED 转账失败(终态)
	// TRANSFER_ING 转账中
	// TRANSFER_PART 转账部分成功(终态)
	// TRANSFER_SUCCESS 转账成功(终态)
	TransferStatus enum.TransferStatus `json:"transfer_status,omitempty"`
	// TransferFinishTime 转账完成时间
	TransferFinishTime string `json:"transfer_finish_time,omitempty"`
	// TransferCreateTime 转账创建时间
	TransferCreateTime string `json:"transfer_create_time,omitempty"`
	// TransferTargetRecordList 账户信息列表
	TransferTargetRecordList []TransferTargetRecord `json:"transfer_target_record_list,omitempty"`
}

// TransferTargetRecord 账户信息
type TransferTargetRecord struct {
	// AccountID 锚定账户id，1:N的1
	AccountID uint64 `json:"account_id,omitempty"`
	// TargetAccountID 目标账户id，1:N的N
	TargetAccountID uint64 `json:"target_account_id,omitempty"`
	// TransferAmount 转账金额(单位：分)
	TransferAmount int64 `json:"transfer_amount,omitempty"`
	// TransferStatus 账户间转账状态 可选值:
	// NO_TRANSFER 未转账
	// TRANSFER_FAILED 转账失败(终态)
	// TRANSFER_ING 转账中
	// TRANSFER_PART 转账部分成功(终态)
	// TRANSFER_SUCCESS 转账成功(终态)
	TransferStatus enum.TransferStatus `json:"transfer_status,omitempty"`
	// TransferCapitalRecordList 转账资金类型列表
	TransferCapitalRecordList []TransferCapitalRecord `json:"transfer_capital_record_list,omitempty"`
}

// TransferCapitalRecord 转账资金类型
type TransferCapitalRecord struct {
	// CapitalType 可转资金类型 可选值:
	// CREDIT_BIDDING 授信竞价
	// CREDIT_BRAND 授信品牌
	// CREDIT_GENERAL 授信通用
	// PREPAY_BIDDING 预付竞价
	// PREPAY_BRAND 预付品牌
	// PREPAY_GENERAL 预付通用
	CapitalType enum.CapitalType `json:"capital_type,omitempty"`
	// TransferBalance 可转资金金额(单位：分)
	TransferBalance int64 `json:"transfer_balance,omitempty"`
	// TransferStatus转账资金状态 可选值:
	// NO_TRANSFER 未转账
	// TRANSFER_FAILED 转账失败(终态)
	// TRANSFER_ING 转账中
	// TRANSFER_PART 转账部分成功(终态)
	// TRANSFER_SUCCESS 转账成功(终态)
	TransferStatus enum.TransferStatus `json:"transfer_status,omitempty"`
	// FailReason 失败原因
	FailReason string `json:"fail_reason,omitempty"`
}
