package mint

import (
	"fmt"
	"github.com/GwiBo/go/CoinJam/num"
	"log"
)

const DEBUG bool = false

//=========================================================
//=========================================================

func MakeCoins(length int64, count int64) {
	if DEBUG {
		log.Printf("MakeCoins() : length:%v count:%v", length, count)
	}

	// find random binary string of required length

	var coinString string

	coinMap := make(map[string]int)

	for i := int64(0); i < count; i++ {
		for {
			for {
				//<<<<<<< De-Dup !!! >>>>>>>>>>
				coinString = num.RandBinCoinString(length)
				_, ok := coinMap[coinString]
				if !ok {
					// key/coin-string doesn't exist, OK to proceed
					break
				} else {
					if DEBUG {
						log.Printf("MakeCoins(): Found Duplicate(%v), searching again", coinString)
					}
				}
			}

			if DEBUG {
				log.Printf("MakeCoins():Coin String Received:%v", coinString)
			}

			if num.NotPrimeForAllBases(coinString) {
				coinMap[coinString] = 1
				break
			}
		}

		// number is NOT primary, find divisors
		fmt.Printf(coinString)
		num.PrintDivisorsPerBase(coinString)
		fmt.Printf("\n")

	}

} // end func MakeCoins()

func GetCoins(length int64, count int64) string {

	if DEBUG {
		log.Printf("MakeCoins() : length:%v count:%v", length, count)
	}

	// find random binary string of required length

	var coinString string
	var coinRecord string
	coinMap := make(map[string]int)

	for i := int64(0); i < count; i++ {
		for {
			for {
				//<<<<<<< De-Dup !!! >>>>>>>>>>
				coinString = num.RandBinCoinString(length)
				_, ok := coinMap[coinString]
				if !ok {
					// key/coin-string doesn't exist, OK to proceed
					break
				} else {
					if DEBUG {
						log.Printf("GetCoins(): Found Duplicate(%v), searching again", coinString)
					}
				}
			}

			if DEBUG {
				log.Printf("GetCoins():Coin String Received:%v", coinString)
			}

			if num.NotPrimeForAllBases(coinString) {
				coinMap[coinString] = 1
				break
			}
		}

		if DEBUG {
			log.Printf("GetCoins():Coin String Unique Candidate Found:%v", coinString)
		}

		coinRecord = coinString + " " + num.GetDivisorsPerBase(coinString) + "\n"

	}
	return coinRecord

} // end func GetCoin()

func GenerateCoins(length int64, count int64) {
	coinRequestChan := make(chan int64, 4)
	coinChan := make(chan string, 4)
	quitChan := make(chan int, 4)

	var coinString string

	var routineCount int = 0
	var maxRoutineCount int = 200 //<<<<<<<<<<<<<<<<<<<< Performance Control

	go UniqueCoinProvider(coinRequestChan, coinChan, quitChan)

	if DEBUG {
		log.Printf("GenerateCoins(): Started coin provider")
	}

	resultChan := make(chan string, 4)
	var resultString string
	var resultCount int64 = 0

	coinRequestChan <- length
	if DEBUG {
		log.Printf("GenerateCoins(): First ping sent")
	}

	for {
		select {

		case resultString = <-resultChan:
			routineCount--
			resultCount++

			fmt.Printf("%v\n", resultString) //<<<<<<<<<<<<<<<<<<<<<<<< output
			if resultCount >= count {
				quitChan <- 1
				return
			}

			for {
				coinRequestChan <- length
				routineCount++
				go getResultAsync(<-coinChan, resultChan)
				if routineCount >= maxRoutineCount {
					break
				}
			}

		case coinString = <-coinChan:
			go getResultAsync(coinString, resultChan)
			routineCount++
		}
	}
}

func getResultAsync(coinString string, resultChan chan string) {
	resultChan <- getResult(coinString)
}

func getResult(coinString string) (result string) {
	return coinString + " " + num.GetDivisorsPerBase(coinString)
}

func UniqueCoinProvider(pingCoinLength chan int64, coin chan string, quit chan int) {
	coinMap := make(map[string]int)
	var coinString string

	for {
		select {

		case length := <-pingCoinLength:
			for {
				for {
					//<<<<<<< De-Dup !!! >>>>>>>>>>
					coinString = num.RandBinCoinString(length)
					_, ok := coinMap[coinString]
					if !ok {
						// key/coin-string doesn't exist, OK to proceed
						break
					} else {
						if DEBUG {
							log.Printf("GetCoins(): Found Duplicate(%v), searching again", coinString)
						}
					}
				} // end for De-dup

				if DEBUG {
					log.Printf("GetCoins():Coin String Received:%v", coinString)
				}

				if num.NotPrimeForAllBases(coinString) {
					coinMap[coinString] = 1
					break
				}
			} // end for NotPrimeCondition
			coin <- coinString

		case <-quit:
			return
		}
	}
}
