package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var Magic8Responses []string = []string{
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
	fmt.Println("WELCOME TO YOUR MAGIC 8 BALL")
	fmt.Println("Ask a question and let fate decide for you!")
	fmt.Println("")
	fmt.Print("What knowledge do you seek? [Press ENTER when done]: ")
	var question string
	input := bufio.NewReader(os.Stdin)
	question, err := input.ReadString('\n')
	question = strings.TrimSpace(question)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nYou asked: %v\n", question)
	num := rand.Int() % len(Magic8Responses)
	fmt.Printf("Magic 8 responds: %v\n", Magic8Responses[num])
}
