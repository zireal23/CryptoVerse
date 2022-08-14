package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

const (
	kafkaConn = "localhost:9092"
	topic = "myTopic"
)

func main(){
	//create a new producer

	producer, err := initialiseKafkaProducer();
	if err != nil {
		log.Println("Error while initialising producer: ", err.Error());
		os.Exit(1);
	}


	//read from command line

	reader := bufio.NewReader(os.Stdin);
	for{
		fmt.Println("Enter a message: ");
		msg, err := reader.ReadString('\n');
		if err != nil {
			log.Fatalln("Couldnt read input: ", err.Error());
		}

		publishMessage(msg,producer);
	}

}



func initialiseKafkaProducer() (sarama.SyncProducer, error){
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime);

	//producer config

	config := sarama.NewConfig();
	config.Producer.Retry.Max = 5;
	config.Producer.RequiredAcks = sarama.WaitForAll;
	config.Producer.Return.Successes = true;

	//async producer 
	// prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config);

	//sync producer

	producer, err := sarama.NewSyncProducer([]string{kafkaConn}, config);

	return producer, err;
}



func publishMessage(message string, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg);
	if err != nil {
		log.Fatalln("Error pushing message to Kafka", err.Error());
	}

	fmt.Println("Message pushed to partition: ", partition);
	fmt.Println("Message pushed to offset: ", offset);

}