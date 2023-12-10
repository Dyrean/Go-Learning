package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

const FILE_NAME = "balance.txt"

func writeBalanceToFile(balace float64) error {
	err := os.WriteFile(FILE_NAME, []byte(fmt.Sprint(balace)), 0644)

	if err != nil {
		return err
	}

	return nil
}

func createBalanceFile(balance float64) error {
	return os.WriteFile(FILE_NAME, []byte(fmt.Sprint(balance)), 0644)
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(FILE_NAME)

	if err != nil {
		if err.Error() == fmt.Sprintf(`open %v: The system cannot find the file specified.`, FILE_NAME) {
			for {
				fmt.Printf("Currently there is no balance for you would you like to create? (y/n): ")
				var choice string
				_, err := fmt.Scan(&choice)

				if err != nil {
					fmt.Println("Open Balance Choice Input Error: ", err)
				}

				switch choice {
				case "y", "Y":
					err = createBalanceFile(0.0)

					if err != nil {
						return 0, err
					}

					return 0.0, nil

				case "n", "N":
					return 0.0, errors.New("balance did not created")

				default:
					fmt.Println("Select right input!")
				}
			}
		}
		return 0.0, err
	}

	balance, err := strconv.ParseFloat(string(data), 64)

	if err != nil {
		return 0, err
	}

	return balance, nil
}

func main() {
	now := time.Now()

	defer func() {
		fmt.Println("Time Elapse:", time.Since(now))
	}()

	balance, err := getBalanceFromFile()
	if err != nil {
		fmt.Println("Get Balance Error: ", err)
		return
	}

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice uint
		fmt.Print("Your choice: ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Choice Input Error: ", err)
			return
		}

		switch choice {
		case 1:
			fmt.Println("Your balance is", balance)

		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			_, err := fmt.Scan(&depositAmount)

			if err != nil {
				fmt.Println("Deposit Input Error:", err)
				return
			}

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			balance += depositAmount
			fmt.Printf("Balance updated! %v amount deposited to balance. New balance: %v.\n", depositAmount, balance)
			writeBalanceToFile(balance)

		case 3:
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			_, err := fmt.Scan(&withdrawAmount)

			if err != nil {
				fmt.Println("Withdraw Input Error:", err)
				return
			}

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if balance < withdrawAmount {
				fmt.Println("Invalid amount. You can't withdraw more than your balance.")
				continue
			}

			balance -= withdrawAmount
			fmt.Printf("Balance updated! %v amount withdrawed from balance. New balance: %v.\n", withdrawAmount, balance)
			writeBalanceToFile(balance)

		case 4:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}
