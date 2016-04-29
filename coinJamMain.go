package main

import (
	"bufio"
	"codeJam/CoinJam/mint"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const DEBUG bool = false

func main() {

	// Read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Get number of test cases
	scanner.Scan()
	testCase, err := strconv.ParseInt(scanner.Text(), 10, 64) //base:10, bit-size:0

	if err != nil {
		log.Panic("Unable to read number of test cases:", err.Error())
	}

	//=========== Method 1 =========================
	for i := int64(1); i <= testCase; i++ {
		scanner.Scan()
		// startNum, err := strconv.ParseInt(scanner.Text(), 10, 64) // <<<<<<<<<<<<<<, change this, not relevent to coin problem
		if err != nil {
			log.Panic("Unable to read number from input")
		}

		if DEBUG {
			log.Printf("Parameter String : %v", scanner.Text())
		}

		// read coin generation parameters
		inputSlice := strings.Split(scanner.Text(), " ") // space separated strings

		// first parameter is coin length (N)
		coinLen, err := strconv.ParseInt(inputSlice[0], 10, 64)
		if err != nil {
			log.Panic("Couldn't get coin length from input")
		}

		// second parameter is coin count (J)
		coinCount, err := strconv.ParseInt(inputSlice[1], 10, 64)
		if err != nil {
			log.Panic("Couldn't get coin count from input")
		}

		fmt.Printf("Case #%v:\n", i)
		// mint.MakeCoins(coinLen, coinCount)

		mint.GenerateCoins(coinLen, coinCount)

	}
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	//=========== Method 2 ==========================
	//... alternate to read full file
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// 	if scanner.Text() == "" {
	// 		break
	// 	}
	// }
	//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
