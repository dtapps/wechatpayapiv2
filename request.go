package wechatpayapiv2

import (
	"context"
	"crypto/tls"
	"go.dtapp.net/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param gorequest.Params, certStatus bool, cert *tls.Certificate) (gorequest.Response, error) {

	// 创建请求
	client := gorequest.NewHttp()

	// 设置请求地址
	client.SetUri(url)

	// 设置格式
	client.SetContentTypeXml()

	// 设置用户代理
	client.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 设置参数
	client.SetParams(param)

	// 设置证书
	if certStatus {
		client.SetP12Cert(cert)
	}

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 记录日志
	if c.gormLog.status {
		go c.gormLog.client.MiddlewareXml(ctx, request)
	}
	if c.mongoLog.status {
		go c.mongoLog.client.MiddlewareXml(ctx, request)
	}

	return request, err
}
