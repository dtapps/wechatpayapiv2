package wechatpayapiv2

import (
	"context"
	"encoding/xml"
	"net/http"
)

// PayUnifiedOrderNotifyHttpRequest 小程序支付 - 统一下单 - 回调通知 - 请求参数
type PayUnifiedOrderNotifyHttpRequest struct {
	ReturnCode         string `json:"return_code" xml:"return_code"`                   // 返回状态码
	ReturnMsg          string `json:"return_msg" xml:"return_msg"`                     // 返回信息
	Appid              string `json:"appid" xml:"appid"`                               // 小程序ID
	MchId              string `json:"mch_id" xml:"mch_id"`                             // 商户号
	DeviceInfo         string `json:"device_info" xml:"device_info"`                   // 设备号
	NonceStr           string `json:"nonce_str" xml:"nonce_str"`                       // 随机字符串
	Sign               string `json:"sign" xml:"sign"`                                 // 签名
	SignType           string `json:"sign_type" xml:"sign_type"`                       // 签名类型
	ResultCode         string `json:"result_code" xml:"result_code"`                   // 业务结果
	ErrCode            string `json:"err_code" xml:"err_code"`                         // 错误代码
	ErrCodeDes         string `json:"err_code_des" xml:"err_code_des"`                 // 错误代码描述
	Openid             string `json:"openid" xml:"openid"`                             // 用户标识
	IsSubscribe        string `json:"is_subscribe" xml:"is_subscribe"`                 // 是否关注公众账号
	TradeType          string `json:"trade_type" xml:"trade_type"`                     // 交易类型
	BankType           string `json:"bank_type" xml:"bank_type"`                       // 付款银行
	TotalFee           int    `json:"total_fee" xml:"total_fee"`                       // 订单金额
	SettlementTotalFee int    `json:"settlement_total_fee" xml:"settlement_total_fee"` // 应结订单金额
	FeeType            string `json:"fee_type" xml:"fee_type"`                         // 货币种类
	CashFee            int    `json:"cash_fee" xml:"cash_fee"`                         // 现金支付金额
	CashFeeType        string `json:"cash_fee_type" xml:"cash_fee_type"`               // 现金支付货币类型
	CouponFee          string `json:"coupon_fee" xml:"coupon_fee"`                     // 总代金券金额
	CouponCount        int    `json:"coupon_count" xml:"coupon_count"`                 // 代金券使用数量
	CouponType         string `json:"coupon_type" xml:"coupon_type"`                   // 代金券类型
	CouponId           string `json:"coupon_id" xml:"coupon_id"`                       // 代金券ID
	TransactionId      string `json:"transaction_id" xml:"transaction_id"`             // 微信支付订单号
	OutTradeNo         string `json:"out_trade_no" xml:"out_trade_no"`                 // 商户订单号
	Attach             string `json:"attach" xml:"attach"`                             // 商家数据包
	TimeEnd            string `json:"time_end" xml:"time_end"`                         // 支付完成时间
}

// PayUnifiedOrderNotifyHttp 小程序支付 - 统一下单 - 回调通知
// https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_7&index=8
func (c *Client) PayUnifiedOrderNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateXml PayUnifiedOrderNotifyHttpRequest, err error) {
	err = xml.NewDecoder(r.Body).Decode(&validateXml)
	return validateXml, err
}
