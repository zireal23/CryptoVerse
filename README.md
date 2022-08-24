![GitHub language count](https://img.shields.io/github/languages/count/zireal23/CryptoVerse) ![GitHub top language](https://img.shields.io/github/languages/top/zireal23/CryptoVerse) ![GitHub followers](https://img.shields.io/github/followers/zireal23?style=social)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/zireal23/CryptoVerse)
# CRYPTOVERSE
![icon](https://user-images.githubusercontent.com/62436360/183359750-f6f6c2f3-bafb-427b-a7dc-21a67040c85d.png)

CryptoVerse is a website to keep track of all things crypto.
Keep Track of rising and falling crypto prices as well as daily price changes, valuation and market cap of individual currencies and a bulletin board of your desired currency.


![Currency Page](https://user-images.githubusercontent.com/62436360/183359665-e4bb0355-bcda-45d1-8834-34c7711e89df.png)
![Homepage](https://user-images.githubusercontent.com/62436360/183359703-cfd7edd5-9911-4320-81f4-8d1276bb37d9.png)
![news](https://user-images.githubusercontent.com/62436360/183359731-14dc2faa-a631-41a7-9886-faef17215b80.png)


## Features

- Market Cap
- Total Markets
- Total 24h Volume
- CryptoCurrency Prices
- Daily Change in Prices
- Price Gradients
- Currency Overview
- News
- Aggregate Linear, Geometric and Harmonic Means
- Streaming CryptoCurrency data in real time



## Tech Used
![Screenshot from 2022-08-24 12-42-03](https://user-images.githubusercontent.com/62436360/186360059-d2ce733f-90cc-4d83-9627-e8348e2e5b59.png)
- CryptoCurrency Data fetching is done at intervals of 10 seconds from CoinRanking API through a golang application
   that extracts prices, exhchanges, dips and rises.
- The data is converted to a kafka encoded byte message with the name of the cryptocurrency as the key and the message as the value.
  The byte encoding is done adhering to the protobuffer schema and the message then goes through an additonal schema verification.
- Golang based Consumer application that reads from the offset of the kafka Cluster and converts the byte slice back to the original message.
- Each individual message consists of the cryptocurrency ID, name of the currency and the last price including the timestamp.
- The messages are then processed to get the linear, geomtric and harmonic aggregates. The aggregate calculation is done using rolling window functions with window size of 1 day.
- MongoDB time series collections are used as a sink to persist the aggregates as well as the real time prices.
- NodeJs-Express Backend to serve real time  prices and aggregates through REST endpoints with support for gRPC endpoints as well.
- React App built using ANT Designs and React Charts to display data systematically and meaningfully.


- HTML
- CSS
- JavaScript
- React JS
- Redux
- ANT UI
- CoinRanking API
- Apache KAfka
- Golang
- gRPC
- Protobuffers schema 
- MongoDB
- Docker


## Run Locally
Docker
- Download and install docker-engine to run the containerized apps and docker-desktop to monitor the apps.
- Download and install Docker compose.
- Download the docker-compose.yml file from the repository.
- From the terminal, run
```
docker compose up
```
- Check the docker logs to make sure the apps are running.
- To check the produced messages and topics created, visit kafdrop's homepage `localhost:9000` 
- To check produced messages, connect to the mongoDB container using mongoDB compass with the authentication string `mongodb://zirael:sayan@mongodb:27017/`



Clone the project

```bash
  git clone https://github.com/zireal23/CryptoVerse.git
```

Go to the project directory

```bash
  cd Cryptoverse
```

Install dependencies

```bash
  npm install
```

Start the server

```bash
  npm run start
```

