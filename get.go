package wechatpayapiv2

import (
	"crypto/tls"
	"go.dtapp.net/golog"
)

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) GetMchId() string {
	return c.config.mchId
}

func (c *Client) GetMchKey() string {
	return c.config.mchKey
}

func (c *Client) GetCertString() string {
	return c.config.certString
}

func (c *Client) GetKeyString() string {
	return c.config.keyString
}

func (c *Client) P12ToPem() (*tls.Certificate, error) {
	pemCert, err := tls.X509KeyPair([]byte(c.GetCertString()), []byte(c.GetKeyString()))
	return &pemCert, err
}

func (c *Client) GetLogGorm() *golog.ApiClient {
	return c.log.logGormClient
}

func (c *Client) GetLogMongo() *golog.ApiClient {
	return c.log.logMongoClient
}
