package main

import (
	"fmt"
	_ "github.com/kelseyhightower/envconfig"
	"thismod/config"
)

func main() {
	conf, err := config.Load()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Port: ", conf.Port)
		fmt.Println("Jaeger url: ", conf.JaegerUrl)
		fmt.Println("Sentry url: ", conf.SentryUrl)
		fmt.Println("Kafka broker: ", conf.KafkaBroker)
		fmt.Println("App id: ", conf.AppId)
		fmt.Println("App key: ", conf.AppKey)
	}
}
