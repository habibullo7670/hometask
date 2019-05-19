package main

import "fmt"
import "time"

func main() {
	var exitTerminal bool
	var operator int
	var number string
	var sum float32

	for !exitTerminal {
		var operatorCheck bool
		var numberCheck bool
		var sumCheck bool
		var commandCheck bool
		var command string
        var a string
		PrintMainScreen()
        
		for !operatorCheck {
			fmt.Scan(&operator)

			if operator >= 1 && operator <= 4 {
				operatorCheck = true
				if operator==1 {a="Megafon"}
				if operator==2 {a="Babilon"}
				if operator==3 {a="Tcell"}
				if operator==4 {a="Beeline"}
			} else {
				fmt.Println("Please, choose right operator")
			}
		}
        fmt.Println("Please, enter your phone number")
		
		for !numberCheck {
			
			 fmt.Scan(&number)
			 
			 if CheckNumber(operator, number) {
				numberCheck = true
			} else {
				fmt.Println("Please, enter right number")
			}
		}
            fmt.Println("Please, enter your payment")
		for !sumCheck {
			
			fmt.Scan(&sum)

			if sum > 0 {
				sumCheck = true
			} else {
				fmt.Println("Please, enter right sum")
			}
		}   
		    fmt.Println("Operation successful")
	        currentTime:=time.Now()
			fmt.Println("---------------------------------------------------------------")
			fmt.Printf("Your operator: %s \n",a)
			fmt.Printf("Your number: %s \n",number)
			fmt.Printf("Your payment: %f \n",sum)
			fmt.Println(currentTime.Format("02.01.2006 15:04:05"))
			fmt.Println("---------------------------------------------------------------")
	
		    fmt.Println("esc- to go exit, back - to go to terminal")

		for !commandCheck {
			fmt.Scan(&command)

			if command == "esc" {
				commandCheck = true
				exitTerminal = true
			} else if command == "back" {
				commandCheck = true
				continue
			} else {
				fmt.Println("command not recongize")
			}
		}
	}

	fmt.Println("application terminated")
	
	}