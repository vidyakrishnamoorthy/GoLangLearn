// int -> integer
// float -> decimal
// complex -> imaginary numbers
// bool -> boolean
// int8
// int16
// int32
// int64
// unit8
// uint16
// uint32
// uint64
// string

package main

import "fmt"

func main() {
	var jellybeanCounter int8 = 10 // default value is 0

	var lengthOfSong uint16
	var isMusicOver bool
	var songRating float32

	lengthOfSong = 15
	isMusicOver = false
	songRating = 5.5
	var myname string = "anju is my name"

	insync := "She likes InSync"

	fmt.Println(jellybeanCounter, lengthOfSong, isMusicOver, songRating)

	fmt.Println("hey there! "+myname+" "+insync, jellybeanCounter, lengthOfSong, isMusicOver, songRating)

}
