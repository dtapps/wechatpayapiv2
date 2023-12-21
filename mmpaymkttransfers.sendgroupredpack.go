package wechatpayapiv2

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/gorandom"
	"go.dtapp.net/gorequest"
)

type MmpaymkttransfersSendgroupredpackResponse struct {
	ReturnCode string `json:"return_code" xml:"return_code"`                   // 返回状态码
	ReturnMsg  string `json:"return_msg,omitempty" xml:"return_msg,omitempty"` // 返回信息

	ResultCode string `json:"result_code" xml:"result_code"`                       // 业务结果
	ErrCode    string `json:"err_code,omitempty" xml:"err_code,omitempty"`         // 错误代码
	ErrCodeDes string `json:"err_code_des,omitempty" xml:"err_code_des,omitempty"` // 错误代码描述

	MchBillno   string `json:"mch_billno" xml:"mch_billno"`     // 商户订单号
	MchId       string `json:"mch_id" xml:"mch_id"`             // 商户号
	Wxappid     string `json:"wxappid" xml:"wxappid"`           // 公众账号appid
	ReOpenid    string `json:"re_openid" xml:"re_openid"`       // 用户openid
	TotalAmount int64  `json:"total_amount" xml:"total_amount"` // 付款金额
	SendListid  string `json:"send_listid" xml:"send_listid"`   // 微信单号
}

type MmpaymkttransfersSendgroupredpackResult struct {
	Result MmpaymkttransfersSendgroupredpackResponse // 结果
	Body   []byte                                    // 内容
	Http   gorequest.Response                        // 请求
}

func newMmpaymkttransfersSendgroupredpackResult(result MmpaymkttransfersSendgroupredpackResponse, body []byte, http gorequest.Response) *MmpaymkttransfersSendgroupredpackResult {
	return &MmpaymkttransfersSendgroupredpackResult{Result: result, Body: body, Http: http}
}

// MmpaymkttransfersSendgroupredpack
// 现金红包 - 发放裂变红包
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon_sl.php?chapter=13_5&index=4
func (c *Client) MmpaymkttransfersSendgroupredpack(ctx context.Context, notMustParams ...gorequest.Params) (*MmpaymkttransfersSendgroupredpackResult, error) {
	cert, err := c.P12ToPem()
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("nonce_str", gorandom.Alphanumeric(32)) // 随机字符串
	// 签名
	params.Set("sign", c.getMd5Sign(params))
	// 	请求
	request, err := c.request(ctx, apiUrl+"/mmpaymkttransfers/sendgroupredpack", params, true, cert)
	if err != nil {
		return newMmpaymkttransfersSendgroupredpackResult(MmpaymkttransfersSendgroupredpackResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response MmpaymkttransfersSendgroupredpackResponse
	err = xml.Unmarshal(request.ResponseBody, &response)
	return newMmpaymkttransfersSendgroupredpackResult(response, request.ResponseBody, request), err
}
