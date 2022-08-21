package lib


/**
* This function calculates the current airthmetic mean of the prices of a particular crypto currency.
* The function takes the current price and the crypto symbol as parameters.
* The function first checks if the map has been initialized or not.
* If not then it initializes the map.
* The function then calculates the current arithmetic mean by adding the current price to the old arithmetic mean and dividing it by the number of elements.
* We also need to check if the number of elements is less than the limit or not.
* If not then we need to subtract the oldest price from the array from the current arithmetic mean.
* The function then updates the arithmetic mean in the map and returns the current arithmetic mean.
**/


func CalulateCurrentArithmeticMean(currentPrice float32, cryptoSymbol string) float32 {

	CheckAndInitCurrencyMap(cryptoSymbol);
		
	oldArithmeticMean := CryptoAggregatePrices[cryptoSymbol].ArithmeticMean;
	numberOfElements := CryptoAggregatePrices[cryptoSymbol].NumberOfElements;
	currentArithmeticMean := oldArithmeticMean + currentPrice;
	
	//If the number of elements in the array is less than limit, then we just need to add the new price and increment the number of elements
	//If not then we need to add the new price and subtract the last price in the array as well
	if numberOfElements<LimitOfArrayElements {
		numberOfElements++;
	}else{
		currentArithmeticMean -= (CryptoAggregatePrices[cryptoSymbol].CryptoPricesArray.Front().Value.(float32));
	}
	

	//Check for debugging errors later
	currentArithmeticMean /= float32(numberOfElements);
	CryptoAggregatePrices[cryptoSymbol].ArithmeticMean = currentArithmeticMean;

	return currentArithmeticMean;
}

// func CalculateCurrentGeometricMean(currentPrice float32, cryptoSymbol string){

// }