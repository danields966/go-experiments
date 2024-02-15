package main

/*
Check whether the password entered by the user meets the specified requirements.
The password must be at least 5 characters and can only contain digits and Latin letters.
Scan the password string from Stdin. If it meets the requirements, print "Ok", otherwise print "Wrong password"
*/

import (
	"fmt"
	"unicode"
)

func main() {
	var password string

	if _, err := fmt.Scan(&password); err != nil {
		panic(err)
	}
	rs := []rune(password)

	isCorrect := true
	if len(rs) < 5 {
		isCorrect = false
	} else {
		for _, ch := range rs {
			if unicode.In(ch, unicode.Latin, unicode.Digit) == false {
				isCorrect = false
				break
			}
		}
	}

	if isCorrect {
		fmt.Print("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
