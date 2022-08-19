package lib





func CalulateCurrentArithmeticMean(currentPrice float32, cryptoSymbol string) float32 {
	numberOfElements := NumberOfElements;
	oldArithmeticMean := CryptoAggregatePrices[cryptoSymbol].ArithmeticMean;
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