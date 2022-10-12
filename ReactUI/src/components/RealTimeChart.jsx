import React, { useState, useEffect } from "react";
import {
  BarChart,
  Bar,
  Line,
  LineChart,
  XAxis,
  YAxis,
  Tooltip,
  CartesianGrid,
  Legend,
  ResponsiveContainer,
} from "recharts";

const RealTimeChart = (props) => {
  const coinName = props.coinName;
  const [coinData, setCoinData] = useState([]);
  useEffect(() => {
    const interval = setInterval(() => {
      fetch(
        `https://cryptobackend-zireal23.koyeb.app/latestaggregate/${coinName}`
      )
        .then((res) => res.json())
        .then((result) => {
          const currdate = new Date(result.PriceAt * 1000);
          const newCoinData = {
            timeStamp: currdate.toLocaleTimeString(),
            currentPrice: result.RealPrice,
            linearMean: result.ArithmeticAggregatePrice,
            geometricMean: result.GeometricAggregatePrice * 600,
            harmonicMean: result.HarmonicAggregatePrice / 10,
          };
          setCoinData((currentData) => [...currentData, newCoinData]);
          console.log(newCoinData.currentPrice);
        });
    }, 10000);
    return () => clearInterval(interval);
  }, [coinName]);

  return (
    <div>
      <ResponsiveContainer width="90%" height={500}>
        <LineChart data={coinData}>
          <XAxis dataKey="timeStamp" />
          <YAxis />
          <CartesianGrid stroke="#eee" strokeDasharray="5 5" />
          <Tooltip />
          <Legend verticalAlign="top" height={36} />
          <Line
            name="Price"
            type="monotone"
            dataKey="currentPrice"
            stroke="#8884d8"
          />
          <Line
            name="Linear Mean"
            type="monotone"
            dataKey="linearMean"
            stroke="#82ca9d"
          />
          <Line
            name="Geometric Mean"
            type="monotone"
            dataKey="geometricMean"
            stroke="#853451"
          />
          <Line
            name="Harmonic Mean"
            type="monotone"
            dataKey="harmonicMean"
            stroke="#845321"
          />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
};

export default RealTimeChart;
