package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	kafkaSchemapb "github.com/zirael23/CryptoStreams/kafkaSchema"
	"github.com/zirael23/CryptoStreams/lib"
	"github.com/zirael23/CryptoStreams/util"
	"google.golang.org/protobuf/proto"
)

var (
    topic = []string{
        "cryptoTopic",
    }
    MIN_COMMIT_COUNT = 100;
)

func main() {

     kafkaConsumer := initConsumer();
    log.Println("The kafka consumer was successfully initialised");

    dbResources, err := util.OpenDatabaseConnection();
    if err != nil {
        log.Println("No DB Connection");
    }
    log.Println("Successfully Connected to DB");
    lib.InitMap();
    consumeKafkaMessages(kafkaConsumer, dbResources);

    util.CloseDatabaseConnection(dbResources);
}





func initConsumer() *kafka.Consumer {
    configMap := kafka.ConfigMap{
        "bootstrap.servers": "kafka:9092",
        "group.id":          "kafkaStreamer",
        "auto.offset.reset": "smallest",
        "enable.auto.commit": "false",
    }
    kafkaConsumer, err := kafka.NewConsumer(&configMap);

    if err != nil {
        log.Println("The consumer failed to initialise", err.Error());
    }
    return kafkaConsumer;
}

func consumeKafkaMessages(kafkaConsumer *kafka.Consumer, dbResources util.DBResources) {
    err := kafkaConsumer.SubscribeTopics(topic,nil);

    if err != nil{
        log.Println("Couldnt subscribe to kafka topic", err.Error());
    }

    numberOfMessagesRead := 0;
    run := true;

    for run {
        kafkaevent := kafkaConsumer.Poll(100);
        
        switch event := kafkaevent.(type){
        case *kafka.Message:
            numberOfMessagesRead += 1;
            if numberOfMessagesRead % 100 == 0{
                kafkaConsumer.Commit();
            }
            log.Println(string(event.Key));
            var coinDataResponse kafkaSchemapb.CoinData;
            eventResponse := event.Value;
            proto.Unmarshal(eventResponse,&coinDataResponse);
            coinDataforDB := util.CoinPriceData{
                ID: coinDataResponse.GetId(),
                Name: coinDataResponse.GetName(),
                RealPrice: coinDataResponse.GetPrice(),
                Timestamp: time.Unix(coinDataResponse.GetTimestamp(),0),
                ArithmeticAggregatePrice: lib.CalulateCurrentArithmeticMean(coinDataResponse.GetPrice(),coinDataResponse.GetId()),
            }
            log.Println("The number of messages read is:", numberOfMessagesRead);
            util.InsertCoinPricesToDB(dbResources,coinDataforDB);
        case kafka.PartitionEOF:
            log.Printf("%% Reached %v\n", event);
        case kafka.Error:
            fmt.Fprintf(os.Stderr, "%% Error: %v\n",event);
            run = false;
            util.CloseDatabaseConnection(dbResources);
            kafkaConsumer.Close();
        default:
            log.Printf("Ignored %v\n", event);
        }

    }

}