package num

import (
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const DEBUG bool = false

//=====================================================================================
//=====================================================================================

func randBinString(length int64) (binString string) {
	var maxLimit int64 = 1<<uint(length) - 1

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	binString = strconv.FormatInt(r.Int63n(maxLimit), 2)

	if int64(len(binString)) < length {

		if DEBUG {
			log.Printf("num.RandBinString(): Binary string shorter than required, required length:%v, found length:%v", length, len(binString))
			log.Printf("num.RandBinString(): String:[%v]", binString)
		}

		lenDiff := length - int64(len(binString))

		for i := int64(0); i < lenDiff; i++ {
			binString = "0" + binString
		}
		// stringLen := int64(len(binString))
		// if DEBUG {
		// 	log.Printf("num.RandBinString(): binString:[%v] stringLen:[%v] (before padding)", binString, len(binString))
		// }
		// for i := int64(0); stringLen == length; i++ {
		// 	binString = "0" + binString
		// 	stringLen = int64(len(binString))
		// 	if DEBUG {
		// 		log.Printf("num.RandBinString(): binString:[%v] stringLen:[%v]", binString, len(binString))
		// 	}
		// }

		if DEBUG {
			log.Printf("num.RandBinString(): Returning String:[%v] Length:[%v]", binString, len(binString))
		}
	}

	return binString
}

//=====================================================================================

func RandBinCoinString(length int64) (binString string) {
	return "1" + randBinString(length-2) + "1"
}

//=====================================================================================
//=====================================================================================

func NotPrimeForAllBases(binaryNumString string) (isPrime bool) {
	val := new(big.Int)

	for i := 2; i <= 10; i++ { // <<<<<<<<<<< i selects the base for string
		val.SetString(binaryNumString, i)

		isPrime := val.ProbablyPrime(1) //<<<<<<<<<<<<<<<<<<<<<<< PERFORMANCE Sensitive !!

		if DEBUG {
			log.Printf("NotPrimeForAllBases(): Value:%v Base:%v  is prime: %v \n", val, i, isPrime)
		}

		if isPrime {
			return false
		} else {
			//.. do nothing as number for current BASE is NOT Prime
		}

	}
	return true
}

//=====================================================================================
//=====================================================================================

func PrintDivisorsPerBase(binaryNumString string) {
	val := new(big.Int)
	for i := 2; i <= 10; i++ { // <<<<<<<<<<< i selects the base for string
		val.SetString(binaryNumString, i)

		if DEBUG {
			log.Printf("PrintDivisorsPerBase(): Value:%v Base:%v \n", val, i)
		}

		divisor, err := getDivisorAsString(val)
		if err != nil {
			log.Panic("PrintDivisorPerBase(): Divisor not found!!")
		}
		fmt.Printf(" %v", divisor)
	}
}

func GetDivisorsPerBase(binaryNumString string) (divisorRecord string) {
	val := new(big.Int)
	for i := 2; i <= 10; i++ { // <<<<<<<<<<< i selects the base for string
		val.SetString(binaryNumString, i)

		if DEBUG {
			log.Printf("PrintDivisorsPerBase(): Value:%v Base:%v \n", val, i)
		}

		divisor, err := getDivisorAsString(val)
		if err != nil {
			log.Panic("PrintDivisorPerBase(): Divisor not found!!")
		}
		divisorRecord += " " + divisor
	}
	return strings.Trim(divisorRecord, " ")
}

type DivisorNotFoundError struct {
	msg string
}

func (err DivisorNotFoundError) Error() string {
	return "Divisor not found!!!"
}

func getDivisorAsString(num *big.Int) (divisor string, err error) {
	i := new(big.Int)
	i.SetInt64(2)

	mod := new(big.Int)

	zero := new(big.Int)
	zero.SetInt64(0)

	one := new(big.Int)
	one.SetInt64(1)

	inc := new(big.Int)
	inc.SetInt64(1)

	max := new(big.Int)
	max.Sub(num, one)

	var divisorFound bool = false
	for {
		mod.Mod(num, i)
		if mod.Cmp(zero) == 0 {
			divisorFound = true
			break
		}

		i.Add(i, inc)
		if i.Cmp(max) == 0 {
			divisorFound = false
			break
		}
	}

	if divisorFound {
		return i.String(), nil
	} else {
		return "", DivisorNotFoundError{"Divisor not found!!"}
	}
}
