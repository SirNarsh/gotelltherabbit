package readconf

// GeneralConf used to deserialize general.json Contains enabled features & RabbitMQ Server
type GeneralConf struct {
	RabbitMQServer, ServerBind           string
	EnableRabbit2HTTP, EnableHTTP2Rabbit bool
}
