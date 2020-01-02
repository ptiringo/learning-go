package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var correctAnswer = 10

	for {
		var answer = rand.Intn(100) + 1

		fmt.Printf("選んだ数は %3v です：", answer)

		if answer == correctAnswer {
			fmt.Println("正解です！")
			break
		} else if answer < correctAnswer {
			fmt.Println("小さすぎます。")
		} else if answer > correctAnswer {
			fmt.Println("大きすぎます。")
		}
	}
}
