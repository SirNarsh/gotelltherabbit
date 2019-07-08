
# GTTR; GO TELL THE RABBIT


## <u>Work in progress, not ready for usage</u>

<!-- ABOUT THE PROJECT -->
## About The Project

Go tell the rabbit (RabbitMQ) is a minimalist HTTP to RabbitMQ & RabbitMQ to HTTP bridge/middleware written in golang

- No Auth
- No message manipulation, checking, verification
- Limited to synchronis actions only: HTTP 200 means message is delivered to RabbitMQ, ACK messages in RabbitMQ means HTTP was forwarded and resulted in 200.


## Planned features

### A. HTTP to RabbitMQ  (Configurable in config/http2rabbit.json)
* Listen to :8080 forward requests with pattern /exchange/{exchange_name} to RabbitMQ {exchange_name}
* Ether whitelist exchanges user is allowed to post to or add "*" to allow sending to all exchanges

(No auth or https handling please use nginx proxy for if needed.)

### B. RabbitMQ to HTTP (Configurable in config/rabbit2http.json)
* Bind all exchange in config to 'gotelltherabbit' queue
* Whenever a message is sent to any of defined exchanges in config, GTTR will forwarded to respective HTTP URL from config along with defined custom headers (eg: 'Auth: Bearer xyz')

## Example usecase
* You have an external microservice that's outside of your network, instead of exposing RabbitMQ you want to allow external microservice to send messages to your RabbitMQ server by just sending an HTTP request, then 'gotelltherabbit' will forward this message to RabbitMQ exchange, maybe you have a nginx proxy in the middle that also handles authentication.
* You are bridging between 3rd party that uses only REST or Webhooks and your microservices that only use RabbitMQ
* You are moving microservices projects one by one from RPI communication pattern to messaging communication pattern and need an interm solution


### Built With
* [Golang](https://golang.org/)
* [Go AMQP](https://github.com/streadway/amqp)



### License
Released under the [MIT License](LICENSE).
