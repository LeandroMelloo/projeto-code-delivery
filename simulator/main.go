package main

import (
	"fmt"
	kafka2 "github.com/projetocodedelivery/simulator/application/kafka"
	"github.com/projetocodedelivery/simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume() // trabalhar de forma assícrona
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg) // trabalhar de forma assícrona
	}

}