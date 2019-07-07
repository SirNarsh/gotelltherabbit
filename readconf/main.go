package readconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var rabbit2http Rabbit2http
var r2HLoaded = false
var http2rabbit HTTP2rabbit
var h2RLoaded = false
var generalConf GeneralConf
var confLoaded = false

func load(filePath string, configPointer interface{}) {
	log.Printf("Loading config file %s", filePath)
	jsonData, err := ioutil.ReadFile(filePath)
	failOnError(err, fmt.Sprintf("Error loading config file %s", filePath))
	err2 := json.Unmarshal(jsonData, &configPointer)
	failOnError(err2, fmt.Sprintf("Error deserializing %s", filePath))
	log.Printf("Loaded config file %s", filePath)
}

// GetGeneral return GeneralConf shared by all components
func GetGeneral() GeneralConf {
	if !confLoaded {
		load("./config/general.json", &generalConf)
		confLoaded = true
	}
	return generalConf
}

// GetR2H Return RabbitMQ to HTTP Config
func GetR2H() Rabbit2http {
	if !r2HLoaded {
		load("./config/rabbit2http.json", &rabbit2http)
		r2HLoaded = true
	}
	return rabbit2http
}

// GetH2R Return Http to RabbitMQ Config
func GetH2R() HTTP2rabbit {
	if !h2RLoaded {
		load("./config/http2rabbit.json", &http2rabbit)
		h2RLoaded = true
	}
	return http2rabbit
}
