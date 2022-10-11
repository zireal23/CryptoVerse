import React from 'react'
import { Coin, CoinData } from "../proto/cryptoData_pb";
import { FetchCryptoCoinDataClient } from "../proto/cryptoData_grpc_web_pb";


const LineChart = (props) => {
  const currentCoinName = props.coinName;
  const client = new FetchCryptoCoinDataClient('http://localhost:8000'); 
  const streamRequest = new Coin();
  streamRequest.setCoinname(currentCoinName);

  const stream = client.getCoinData(streamRequest, {});
  stream.on("data", (response) => {
    console.log(response.getName());
  });
  stream.on("error", (err) => {
    console.log(err);
  });

  return (
    <div>LineChart</div>
  )
}

export default LineChart