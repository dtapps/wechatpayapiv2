package wechatpayapiv2

import (
	"go.dtapp.net/dorm"
	"go.dtapp.net/golog"
	"go.dtapp.net/gorequest"
)

// client *dorm.GormClient
type gormClientFun func() *dorm.GormClient

// client *dorm.MongoClient
// databaseName string
type mongoClientFun func() (*dorm.MongoClient, string)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId          string // 小程序或者公众号唯一凭证
	AppSecret      string // 小程序或者公众号唯一凭证密钥
	MchId          string // 微信支付的商户id
	MchKey         string // 私钥
	CertString     string
	KeyString      string
	GormClientFun  gormClientFun  // 日志配置
	MongoClientFun mongoClientFun // 日志配置
	Debug          bool           // 日志开关
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		appId      string // 小程序或者公众号唯一凭证
		appSecret  string // 小程序或者公众号唯一凭证密钥
		mchId      string // 微信支付的商户id
		mchKey     string // 私钥
		certString string
		keyString  string
	}
	log struct {
		gorm           bool              // 日志开关
		gormClient     *dorm.GormClient  // 日志数据库
		logGormClient  *golog.ApiClient  // 日志服务
		mongo          bool              // 日志开关
		mongoClient    *dorm.MongoClient // 日志数据库
		logMongoClient *golog.ApiClient  // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	var err error
	c := &Client{}

	c.config.appId = config.AppId
	c.config.appSecret = config.AppSecret
	c.config.mchId = config.MchId
	c.config.mchKey = config.MchKey
	c.config.certString = config.CertString
	c.config.keyString = config.KeyString

	c.requestClient = gorequest.NewHttp()

	gormClient := config.GormClientFun()
	if gormClient != nil && gormClient.Db != nil {
		c.log.logGormClient, err = golog.NewApiGormClient(&golog.ApiGormClientConfig{
			GormClientFun: func() (*dorm.GormClient, string) {
				return gormClient, logTable
			},
			Debug: config.Debug,
		})
		if err != nil {
			return nil, err
		}
		c.log.gorm = true
		c.log.gormClient = gormClient
	}

	mongoClient, databaseName := config.MongoClientFun()
	if mongoClient != nil && mongoClient.Db != nil {
		c.log.logMongoClient, err = golog.NewApiMongoClient(&golog.ApiMongoClientConfig{
			MongoClientFun: func() (*dorm.MongoClient, string, string) {
				return mongoClient, databaseName, logTable
			},
			Debug: config.Debug,
		})
		if err != nil {
			return nil, err
		}
		c.log.mongo = true
		c.log.mongoClient = mongoClient
	}

	return c, nil
}
