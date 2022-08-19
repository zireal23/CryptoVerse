package lib

import (
	"container/list"
	"log"
)

/**
* For every coin I am storing the prices of every 10 second interval for a 24 hour period.
* So there will be 8640 records for every coin.
* Now as soon as there are more than 8640 records, I remove the last record and store the most recent one.
* This way I can keep maintaining my rolling array in O(1) order and the averages as well.
* So I have a struct that will have an array for holding the prices, an integer to keep track of the number
* of the elements in the array and three floats to keep track of the current linear,
* geomteric and harmonic means.
**/


/**
* The cryptoPricesArray is a FIFO structure that supports pushing elements to the end of the list
* as well as removing the first element of the list that is also the oldest.
* The only overhead is allocating memory and deallocating memory during insertion and removal respectively.
**/

type RollingMeans struct {
	CryptoPricesArray *list.List
	ArithmeticMean float32
	GeomtericMean float32
	HarmonicMean float32
}
var NumberOfElements int32 = 0;
var LimitOfArrayElements int32 = 8640;


	/**
		The map stores the references to structs and not the struct itself
		So when we access a map value what we get is the copy of that struct 
		For reading purposes that is alright
		But for modifying the values, we cant modify the copy since the origin value is unmutated
		Whats best would be to have a map of struct pointers, that way we access the struct stored directly
	**/
var CryptoAggregatePrices map[string]*RollingMeans;



func InitMap(){
	CryptoAggregatePrices = make(map[string]*RollingMeans);
}




func UpdateCryptoStructs(cryptoSymbol string, currentPrice float32){
	if(NumberOfElements+1<=LimitOfArrayElements){
		NumberOfElements++;
	}else{
		/*
		* We are removing the oldest value
		* We hold the original array in a variable, remove the last element and assign it again to the struct
		*/
		log.Println("Reslicing the prices holder");
		oldestCurrencyPrice := CryptoAggregatePrices[cryptoSymbol].CryptoPricesArray.Front();
		CryptoAggregatePrices[cryptoSymbol].CryptoPricesArray.Remove(oldestCurrencyPrice);
	}
	CryptoAggregatePrices[cryptoSymbol].CryptoPricesArray.PushBack(float32(currentPrice));

}