package main

import (
    "fmt"
    "os"
    "strings"
    "os/exec"
    "bufio"
    "io/ioutil"
    "log"
)

// main function
func main() {
	content, err := ioutil.ReadFile("Your_first_name.txt")

     if err != nil {
          log.Fatal(err)
     }
    name := string(content)
    fmt.Println("Hello "+name+ " Boss. I am JARVIS. An AI Assistant made from Golang. Please enter your commands below :\n")

	for true {
		fmt.Print("\nYour command : ")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			choice := scanner.Text()
			input := strings.ToLower(choice)

			if input == "hi" || input == "hello" || input == "hola"{
			fmt.Println("Hi Boss. Happy to see you !")
			
			} else if input == "exit" || input == "quit" || input == "close"{
				fmt.Println("Thank You Boss for your time. Have a great day !")
				os.Exit(0)

			} else if input == "open notepad" || input == "notepad" || input == "start notepad"{
				cmd := exec.Command("notepad")
				if err := cmd.Run(); err != nil { 
					fmt.Println("Error: ", err)
				}
			} else if input == "open calculator" || input == "calculator" || input == "calc" || input == "start calculator"{
				cmd := exec.Command("calc")
				if err := cmd.Run(); err != nil { 
					fmt.Println("Error: ", err)
				}
			} else if input == "open code" || input == "code" || input == "open visual studio code" || input == "visual studio code"  || input == "start code"{
				cmd := exec.Command("code")
				if err := cmd.Run(); err != nil{
					fmt.Println("Error: ", err)
				}
			} else if input == "how are you" || input == "how do you do"{
				fmt.Println("I am fine Boss. Thanks for asking !")

			} else if input == "who created you" || input == "who made you"{
				fmt.Println("I was created by Joel Shine on May 21")

			} else if input == "who are you" || input == "what do you do" || input == "why are you made"{
				fmt.Println("My name is Jarvis. I am an Assistant who works for "+name+". I am written in Golang.")

			} else {
				fmt.Println("Sorry. Couldn't recognize your command.")
			}
		}

	}

}