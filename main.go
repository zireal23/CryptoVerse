package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	"github.com/zirael23/CryptoKafkaProducer/coinApi"
	kafkaSchemapb "github.com/zirael23/CryptoKafkaProducer/kafkaSchema"
	"google.golang.org/protobuf/proto"
)

const (
	kafkaConn = "localhost:9092"
	topic = "crypto"
)

func main(){
	
	coins := coinApi.GetAllCoins();
	
	//create a new producer
	
	producer, err := initialiseKafkaProducer();
	if err != nil {
		log.Println("Error while initialising producer: ", err.Error());
		os.Exit(1);
	}
	kafkaMessage := createMessageFormat(coins[0]);

	publishMessage(kafkaMessage, producer);

}


func createMessageFormat(coinData coinApi.Coin) []byte {
	message := &kafkaSchemapb.CoinData{
		Uuid: coinData.UUID,
		Symbol: coinData.Symbol,
		Name: coinData.Name,
		Marketcap: coinData.MarketCap,
		Price: coinData.Price,
		Change: coinData.Change,
		The24Hvolume: coinData.The24HVolume,
		Btcprice: coinData.BtcPrice,
	}
	kafkaMessage, err := proto.Marshal(message);
	if err != nil {
		log.Println("Error while serializing message", err.Error());
	} 
	return kafkaMessage;
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



func publishMessage(message []byte, producer sarama.SyncProducer) {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg);
	if err != nil {
		log.Fatalln("Error pushing message to Kafka", err.Error());
	}

	fmt.Println("Message pushed to partition: ", partition);
	fmt.Println("Message pushed to offset: ", offset);

}