package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
	kafkaSchemapb "github.com/zirael23/CryptoStreams/kafkaSchema"
	"github.com/zirael23/CryptoStreams/util"
	"google.golang.org/protobuf/proto"
)

const (
    zookeeperConn = "zookeeper:2181"
    consumerGroup = "cryptoConsumer"
    topic = "cryptoTopic"
)

func main() {
    // setup sarama log to stdout
    // sarama.Logger = log.New(os.Stdout, "", log.Ltime)

    // // init consumer
    // consumerGroup, err := initConsumer()
    // if err != nil {
    //     fmt.Println("Error consumer group: ", err.Error())
    //     os.Exit(1)
    // }
    // defer consumerGroup.Close()

    // // run consumer
    // consume(consumerGroup)
    dbResources, err := util.OpenDatabaseConnection();
    if err != nil {
        log.Println("No DB Connection");
    }
    log.Println(dbResources);
 
    coinData := kafkaSchemapb.CoinData{
        Id: "BTC",
        Name: "Bitcoin",
        Price: "1232325",
        Timestamp: time.Time.Unix(time.Now()),
    }
    start := time.Now();
    for i := 0; i < 10000000; i++{
        log.Println(i);
        util.InsertCoinPrices(dbResources,&coinData);
    }
    log.Println("It took", time.Since(start).Seconds());
    util.GetCoinPrices("Bitcoin", dbResources);
    util.CloseDatabaseConnection(dbResources);
}

func initConsumer()(*consumergroup.ConsumerGroup, error) {
    // consumer config
    config := consumergroup.NewConfig()
    config.Offsets.Initial = sarama.OffsetOldest
    config.Offsets.ProcessingTimeout = 1 * time.Second

    // join to consumer group
    consumerGroup, err := consumergroup.JoinConsumerGroup(consumerGroup, []string{topic}, []string{zookeeperConn}, config)
    if err != nil {
        return nil, err
    }

    return consumerGroup, err
}

func consume(consumerGroup *consumergroup.ConsumerGroup) {
	var count int64 = 0;
    for {
        select {
        case msg := <-consumerGroup.Messages():
            // messages coming through chanel
            // only take messages from subscribed topic
	    if msg.Topic != topic {
                continue
            }
			count++;
            kafkaResponseMessage := msg.Value;
			var cryptoMessage kafkaSchemapb.CoinData;
            // coinID := string(msg.Key);
			proto.Unmarshal(kafkaResponseMessage,&cryptoMessage);
			fmt.Println(cryptoMessage.Name,msg.Key);

            // commit to zookeeper that message is read
            // this prevent read message multiple times after restart
            err := consumerGroup.CommitUpto(msg)
            if err != nil {
                fmt.Println("Error commit zookeeper: ", err.Error())
            }
        }
    }
}