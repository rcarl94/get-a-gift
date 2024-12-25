package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	asciiArt "github.com/common-nighthawk/go-figure"

	"github.com/rcarl94/gift-grab/colors"
	"github.com/rcarl94/gift-grab/config"
)

const spacer = "\n🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁 🎁\n\n"

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//fmt.Printf("Loaded config with %d people\n", len(cfg.People))
	asciiArt.NewColorFigure("Merry", "basic", "green", true).Print()
	time.Sleep(2 * time.Second)
	asciiArt.NewColorFigure("Christmas", "basic", "red", true).Print()
	time.Sleep(2 * time.Second)
	fmt.Println(`

             /\
            <  >
             \/
             /\
            /  \
           /++++\
          /  ()  \
          /      \
         /~-~-~-~-\
        /  ()  ()  \
        /          \
       /*&*&*&*&*&*&\
      /  ()  ()  ()  \
      /              \
     /++++++++++++++++\
    /  ()  ()  ()  ()  \
    /                  \
   /~~~~~~~~~~~~~~~~~~~~\
  /  ()  ()  ()  ()  ()  \
  /*&*&*&*&*&*&*&*&*&*&*&\
           |   |
          |-----|
          \_____/
	`)

	time.Sleep(3 * time.Second)
	fmt.Print("\n\n\n")
	fmt.Print(spacer)
	time.Sleep(2 * time.Second)

	for {
		peopleRemaining := make([]config.Person, len(cfg.People))
		copy(peopleRemaining, cfg.People)
		loopCount := len(cfg.People)
		for range loopCount {
			randomInt := rand.Intn(len(peopleRemaining))
			person := peopleRemaining[randomInt]
			anotherRandomInt := rand.Intn(len(person.Descriptors))
			fmt.Printf("%s, open a gift that's %s\n", colors.Red(person.Name), colors.Green(person.Descriptors[anotherRandomInt]))
			reader := bufio.NewReader(os.Stdin)
			for {
				line, err := reader.ReadString('\n')
				if err != nil {
					fmt.Printf("Error while reading input: %s\n", err)
					os.Exit(2)
				}
				if line == "\n" {
					break
				} else {
					anotherRandomInt = rand.Intn(len(person.Descriptors))
					fmt.Printf("How about one that's %s?\n", colors.Green(person.Descriptors[anotherRandomInt]))
				}
			}
			fmt.Print(spacer)
			time.Sleep(2 * time.Second)
			peopleRemaining = removePerson(peopleRemaining, person)
		}
	}
}

func removePerson(people []config.Person, person config.Person) []config.Person {
	for i, p := range people {
		if p.Name == person.Name {
			return append(people[:i], people[i+1:]...)
		}
	}
	return people
}
