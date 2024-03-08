package wechatpayapiv2

import (
	"context"
	"encoding/xml"
	"net/http"
)

// SecApiPayRefundNotifyHttpRequest 小程序支付 - 申请退款 - 回调通知 - 请求参数
type SecApiPayRefundNotifyHttpRequest struct {
	ReturnCode string `json:"return_code" xml:"return_code"` // 返回状态码
	ReturnMsg  string `json:"return_msg" xml:"return_msg"`   // 返回信息
	Appid      string `json:"appid" xml:"appid"`             // 公众账号ID
	MchId      string `json:"mch_id" xml:"mch_id"`           // 退款的商户号
	NonceStr   string `json:"nonce_str" xml:"nonce_str"`     // 随机字符串
	ReqInfo    string `json:"req_info" xml:"req_info"`       // 加密信息
}

// SecApiPayRefundNotifyHttp 小程序支付 - 申请退款 - 回调通知
// https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_16&index=10
func (c *Client) SecApiPayRefundNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateXml SecApiPayRefundNotifyHttpRequest, err error) {
	err = xml.NewDecoder(r.Body).Decode(&validateXml)
	return validateXml, err
}
