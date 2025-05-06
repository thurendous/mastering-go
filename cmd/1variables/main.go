// 1. Create variable name (myFavNum)
// 2. Store your favorite number inside it
// 3. Print that variable to the console
// 4. Create new variable name (uniqueNumber)
// 5. Store the value of (10) inside it.
// 6. Print that to the console

package main

import (
	"fmt"
	"mastering-go/pkg/common"
)

var myFavNum = 9

func main() {
	fmt.Println(myFavNum)

	uniqueNumber := 10
	fmt.Println(uniqueNumber)

	common.SharedFunction()
}
