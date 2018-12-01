package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// error checking function

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// open file, check for error
	dat, err := os.Open("input")
	check(err)

	// create a new scanner from bufio
	// init the tally
	datScan := bufio.NewScanner(dat)
	var tally int

	// interate through the scanned data
	// convert each line to int and add to tally
	for datScan.Scan() {
		lineStr := datScan.Text()
		num, _ := strconv.Atoi(lineStr)
		tally = tally + num
	}
	// print tally
	fmt.Println(tally)

}
