// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 1.

// Helloworld is our first Go program.
//!+
package main

import (
	"fmt"
	"strings"
)

func add1(r rune) rune  {
	return r+1
	
}
func main() {
	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
}

//!-
