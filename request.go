package wechatpayapiv2

import (
	"context"
	"crypto/tls"
	"go.dtapp.net/gorequest"
)

func (c *Client) request(ctx context.Context, url string, params map[string]interface{}, cert *tls.Certificate) (gorequest.Response, error) {

	// 创建请求
	client := c.client

	// 设置请求地址
	client.SetUri(url)

	// 设置格式
	client.SetContentTypeXml()

	// 设置参数
	client.SetParams(params)

	// 设置证书
	client.SetP12Cert(cert)

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.config.PgsqlDb != nil {
		go c.log.GormMiddlewareXml(ctx, request, Version)
	}
	if c.config.MongoDb != nil {
		go c.log.MongoMiddlewareXml(ctx, request, Version)
	}

	return request, err
}
