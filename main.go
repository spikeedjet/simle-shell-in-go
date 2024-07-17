package main

import(
	"bufio"
	"fmt"
	// "errors"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error{
	// //remove the newline character
	// input = strings.TrimSuffix(input,"\n")

	// //prepare the command to execute
	// cmd := exec.Command(input)

	// Remove the newline character and split the input into command and arguments
	input = strings.TrimSuffix(input,"\n")
	args := strings.Split(input, " ")
	cmd := exec.Command(args[0], args[1:]...)


	//set the corrent output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	//execute the command and return the error
	return cmd.Run()
}


func main(){
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		//read the keyboad input
		input, err := reader.ReadString('\n')
		if err!=nil{
			fmt.Fprintln(os.Stderr,err)
		}

		//handle the execution of the input
		if err = execInput(input); err!=nil{
			fmt.Fprintln(os.Stderr,err)
		}
	}
}




