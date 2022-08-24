import mongoose from "mongoose";

// interface cryptoPriceModelInterface {
//   ID: string;
//   Name: string;
//   RealPrice: number;
//   ArithmeticAggregatePrice: number;
//   GeometricAggregatePrice: number;
//   HarmonicAggregatePrice: number;
//   TimeStamp: Date;
// }

const cryptoPriceSchema = new mongoose.Schema({
  ID: { type: String},
  Name: { type: String, required: true },
  RealPrice: { type: Number, required: true },
  ArithmeticAggregatePrice: { type: Number},
  GeometricAggregatePrice: { type: Number},
  HarmonicAggregatePrice: { type: Number},
  timeStamp: { type: Date, required: true },
},
{collection: "cryptoPricesTimeSeries"}
);

export default mongoose.model(
  "CryptoPrice",
  cryptoPriceSchema
);
