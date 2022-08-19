package util

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type CoinPriceData struct{
	ID string
	Name string
	RealPrice float32
	ArithmeticAggregatePrice float32
	GeometricAggregatePrice float32
	HarmonicAggregatePrice float32
	Timestamp time.Time
}

type DBResources struct {
	client *mongo.Client
	ctx context.Context
	cancel context.CancelFunc
	selectedCollection *mongo.Collection
}


func initTimeSeriesCollection(client *mongo.Client) error {
	//creating the database
	cryptoDataDB := client.Database("cryptoDataDB");

	// setting time series data options
	timeSeriesOptions := options.TimeSeries().SetTimeField("timestamp").SetGranularity("seconds");

	// setting mongodb collection options
	collectionOptions := options.CreateCollection().SetTimeSeriesOptions(timeSeriesOptions).SetExpireAfterSeconds(604800);

	//Creating the time series collection
	err := cryptoDataDB.CreateCollection(context.TODO(),"cryptoPrices",collectionOptions);

	if(err != nil){
		return err;
	}
	return nil;
}


func OpenDatabaseConnection() (DBResources, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://zirael:sayan@localhost:27017/"));
	if err != nil {
		log.Printf("Couldnt create mongoDB client due to: %v", err);
		return DBResources{}, err;
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);

	err = client.Connect(ctx);

	var dbResources DBResources;
	
	if(err != nil){
		log.Println("Couldnt connect to mongodb instance");
		cancel();
		return dbResources, err;
	}
	
	//Creating and initialising the time series collection
	err = initTimeSeriesCollection(client);

	if err != nil {
		log.Println("Couldnt create the time series collection");
	}
	
	selectedCollection := client.Database("cryptoDataDB").Collection("cryptoPrices");

	log.Println("Successfully connected to mongo and initialised a time series collection!");

	dbResources = DBResources{
		client: client,
		ctx: ctx,
		cancel: cancel,
		selectedCollection: selectedCollection,

	}

	// //creating a sample document 
	// btcPrice := bson.D{{Key: "Name", Value: "BTC"}, {Key: "Price", Value: "1"}};

	// //Inserting into DB
	// result, err := cryptoPrices.InsertOne(context.TODO(),btcPrice);

	return dbResources, nil;
}


func InsertCoinPricesToDB(dbResources DBResources,coinData CoinPriceData){
	selectedCollection := dbResources.selectedCollection;

	insertCoinPriceQuery := bson.D{
		{Key: "Name", Value: coinData.Name},
		{Key: "RealPrice", Value: coinData.RealPrice},
		{Key: "ArithmeticAggregatePrice", Value: coinData.ArithmeticAggregatePrice},
		{Key: "Timestamp", Value: primitive.NewDateTimeFromTime(coinData.Timestamp)},
	}
	//filter := bson.M{"ID": coinData.GetId()};

	_, err := selectedCollection.InsertOne(context.TODO(),insertCoinPriceQuery);

	if err != nil{
		log.Println("Couldnt insert data into DB", err.Error());
	}
	//log.Println(result);
}


func GetCoinPrices(coin string, dbResources DBResources){
	filter := bson.M{	
		"Name": coin,
	};
	var result bson.M;
	opts := options.FindOne();
	err := dbResources.selectedCollection.FindOne(dbResources.ctx,filter,opts).Decode(&result);
	if err != nil {
		log.Println("Couldnt find coin price");
	}
	queryResult, err := bson.Marshal(result);
	if err != nil {
		log.Println("Couldnt marhshal the query result");
	}
	var coinData CoinPriceData;
	err = bson.Unmarshal(queryResult,&coinData);
	if err != nil {
		log.Println("Couldnt unmarshal the query result into the struct");
	}
	log.Println(coinData.RealPrice, coinData.Name);
}



func CloseDatabaseConnection(dbResources DBResources){
	dbResources.client.Disconnect(dbResources.ctx);
	dbResources.cancel();
	log.Println("Successfully closed connection to mongoDB");
}