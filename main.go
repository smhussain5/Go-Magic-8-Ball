package main

import (
	"fmt"
	"math/rand"
)

var magic8Responses []string = []string{
	"Without a doubt! [Y]",
	"Very doubtful! [N]",
	"Concentrate and ask again... [M]",
	"You may rely on it! [Y]",
	"Outlook not so good! [N]",
	"Ask again later... [M]",
	"As I see it, yes! [Y]",
	"Absolutely not! [N]",
	"Cannot predict now... [M]",
	"It is certain! [Y]",
	"Chances are slim! [N]",
	"Ask again when the stars align... [M]",
}

func main() {
	num := rand.Int() % len(magic8Responses)
	fmt.Println("Hello World!")
	fmt.Println(num)
	fmt.Println(magic8Responses[num])
}
