import express from "express";
import helmet from "helmet";
import cors from "cors";
import bodyParser from "body-parser";
import mongoose from "mongoose";
import apiRouter from "./routes/latestAggregate";

const PORT = process.env.PORT || 5000;
const mongoURI = `mongodb://asif23:sayan@localhost:27017/cryptoDataDB`;


mongoose.connect(mongoURI);

mongoose.connection.on("open", function () {
  console.log("🔗 Connected to MongoDB database.");
});

const app = express();
app.use(helmet());
app.use(
  cors({
    origin: "*",
  })
);
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

app.listen(PORT, function () {
  console.log(`Express app listening on port ${PORT}`);
});


app.get("/", function (req, res) {
    res.status(200).send("hello world!!");
});

app.use("/latestAggregate", apiRouter);