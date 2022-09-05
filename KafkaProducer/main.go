package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	"github.com/joho/godotenv"
	"github.com/zirael23/CryptoKafkaProducer/coinApi"
	kafkaSchemapb "github.com/zirael23/CryptoKafkaProducer/kafkaSchema"
	"google.golang.org/protobuf/proto"
)


func main(){
	if(os.Getenv("GO_ENV")!="production"){
	err := godotenv.Load();

	if err != nil {
		log.Println("Error while loading env file", err.Error());
	}
}
	
	
	producer, err := initialiseKafkaProducer();
	if err != nil {
		log.Println("Error while initialising producer: ", err.Error());
		os.Exit(1);
	}
	
	queryAPIandPublishMessage(producer);	


}

func queryAPIandPublishMessage(producer sarama.SyncProducer){
	var count int64 = 0;
	for {
		startTime := time.Now();
		coins := coinApi.GetAllCoins();
		fmt.Println(len(coins));
		for _, currentCoin  := range coins{
			kafkaMessage := createMessageFormat(currentCoin);
			publishMessage(kafkaMessage, producer,currentCoin);
			time.Sleep(100*time.Millisecond);
		}
		count += 100;
		log.Println("The entire process took: ",time.Since(startTime).Seconds());
		log.Println("No of messages pushed:", count);

	}



}


func createMessageFormat(coinData coinApi.Coin) []byte {
	coinPrice, err  := strconv.ParseFloat(coinData.PriceUsd, 32);
	if err != nil{
		log.Println("Error while converting price to float", err.Error());
	}
	message := &kafkaSchemapb.CoinData{
		Id: coinData.ID,
		Name: coinData.Name,
		Price: float32(coinPrice),
		Timestamp: time.Now().Unix(),
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
	// kafkaConnectionURL := os.Getenv("kafkaConnectionURL");
	// saslMechanism := os.Getenv("SASLMECHANISM");
	// saslUserName := os.Getenv("SASLUSER");
	// saslPassword := os.Getenv("SASLPASSWORD");
	// clientID := os.Getenv("CLIENTID");

	config := sarama.NewConfig();
	config.Producer.Retry.Max = 5;
	config.Producer.RequiredAcks = sarama.WaitForAll;
	config.Producer.Return.Successes = true;
	// config.Net.SASL.Enable = true;
	// config.Net.TLS.Enable = true;
	// config.Net.SASL.Mechanism = sarama.SASLMechanism(saslMechanism);
	// config.Net.SASL.User = saslUserName;
	// config.Net.SASL.Password = saslPassword;
	// config.ClientID = clientID;
	config.ClientID = "crypto_producer";

	//async producer 
	// prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config);

	//sync producer
	var producer sarama.SyncProducer;
	var err error;
	kafkaConnectionURL := os.Getenv("KAFKA_CONNECTION");
	for{
		producer, err = sarama.NewSyncProducer([]string{kafkaConnectionURL}, config);
		if err == nil {
			break;
		}
		log.Println("Couldnt connect to kafka, Retrying....");
	}

	return producer, err;
}



func publishMessage(message []byte, producer sarama.SyncProducer, coin coinApi.Coin) {
	kafkaTopic := os.Getenv("KAFKA_TOPIC");
	msg := &sarama.ProducerMessage{
		Topic: kafkaTopic,
		Key: sarama.StringEncoder(coin.ID),
		Value: sarama.ByteEncoder(message),
	}

	_, _, err := producer.SendMessage(msg);
	if err != nil {
		log.Println("Error pushing message to Kafka", err.Error());
	}

	// fmt.Println("Message pushed to partition: ", partition);
	// fmt.Println("Message pushed to offset: ", offset);

}