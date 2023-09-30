// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/AlecAivazis/survey/v2"
// )

// func main() {
// 	// defer keyboard.Close()
// 	// keyboard.Open()

// 	// options := []string{"Option 1", "Option 2", "Option 3"}
// 	// selectedIndex := 0
// 	// colors := []func(format string, a ...interface{}) string{color.RedString, color.GreenString, color.BlueString}

// 	// fmt.Println("Use arrow keys to select an option and press Enter to confirm:")
// 	// renderOptions(options, selectedIndex, colors)

// 	// for {
// 	// 	_, key, err := keyboard.GetKey()
// 	// 	if err != nil {
// 	// 		fmt.Println("Error reading key:", err)
// 	// 		os.Exit(1)
// 	// 	}

// 	// 	switch key {
// 	// 	case keyboard.KeyArrowLeft:
// 	// 		if selectedIndex > 0 {
// 	// 			selectedIndex--
// 	// 		}
// 	// 	case keyboard.KeyArrowRight:
// 	// 		if selectedIndex < len(options)-1 {
// 	// 			selectedIndex++
// 	// 		}
// 	// 	case keyboard.KeyEnter:
// 	// 		color.Cyan("You selected: %s\n", options[selectedIndex])
// 	// 		return
// 	// 	case keyboard.KeyEsc:
// 	// 		fmt.Println("Exiting without selection.")
// 	// 		return
// 	// 	}
// 	// 	fmt.Print("\033[H\033[2J")
// 	// 	renderOptions(options, selectedIndex, colors)

// 	// }

// 	example()

// }

// func renderOptions(options []string, selectedIndex int, colors []func(format string, a ...interface{}) string) {
// 	for i, option := range options {

// 		if i == selectedIndex {
// 			fmt.Printf("[%s] ", colors[i]("--> "+option))
// 		} else {
// 			fmt.Printf("[ ] %s ", option)
// 		}
// 	}
// 	fmt.Println()

// }

// func example() {
// 	var selectedOption string

// 	// Define a slice of options for the select menu
// 	options := []string{"Option 1", "Option 2", "Option 3"}

// 	// Create a survey select menu
// 	prompt := &survey.Select{
// 		Message: "Select an option:",
// 		Options: options,
// 	}

// 	// Ask the user to select an option
// 	if err := survey.AskOne(prompt, &selectedOption); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Print the selected option
// 	fmt.Printf("You selected: %s\n", selectedOption)
// }

package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// the questions to ask
var qs = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "What is your name?"},
		Validate:  survey.Required,
		Transform: survey.Title,
	},
	{
		Name: "color",
		Prompt: &survey.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue", "green"},
			Default: "red",
		},
	},
	{
		Name:   "age",
		Prompt: &survey.Input{Message: "How old are you?"},
	},
}

func main() {
	// the answers will be written to this struct
	answers := struct {
		Name          string // survey will match the question and field names
		FavoriteColor string `survey:"color"` // or you can tag fields to match a specific name
		Age           int    // if the types don't match, survey will convert it
	}{}

	// perform the questions
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s chose %s.", answers.Name, answers.FavoriteColor)
}
