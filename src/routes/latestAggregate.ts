import express from "express";
import {fetchAggregatePrices,fetchAggregatePricesForCoin} from "../controllers/cryptoAggregatePricesController";

const router = express.Router();


router.get("/:coinSymbol", fetchAggregatePricesForCoin);
router.get("/", fetchAggregatePrices);



export default router;