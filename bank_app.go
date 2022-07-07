package main

// banking application
// features: deposit, withdraw

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type BankAccount struct {
	balance  float32
	userName string
	password string
}

func (b *BankAccount) Deposit(amt float32) int64 {
	b.balance += amt
	fmt.Printf("Your new balance is %.2f\n", b.balance)
	return time.Now().Unix()
}

func (b *BankAccount) Withdraw(amt float32) int64 {

	b.balance -= amt
	fmt.Printf("Your new balance is %.2f\n", b.balance)
	return time.Now().Unix()
}

func (b *BankAccount) ToString() string {
	var str string = b.userName
	var balStr string
	balStr = fmt.Sprintf("%.2f", b.balance)
	return str + ": " + balStr
}

func (b *BankAccount) alter(userName string, password string) {
	b.userName = userName
	b.password = password
}

type BankApplication struct {
	accounts   []*BankAccount //sort based on name or balance
	transDates []int64
}

func NewApplication() *BankApplication {
	return &BankApplication{}
}

func (b *BankApplication) Deposit() {
	var user string
	var password string
LOGIN: // Login block
	fmt.Print("Username: ")
	fmt.Scanf("%s", &user)
	fmt.Scanf(" %c")
	fmt.Printf("Password: ")
	fmt.Scanf("%s", &password)
	fmt.Scanf(" %c")
	for _, val := range b.accounts {
		if val.userName == user && val.password == password {
		ENTERAMOUNT:
			var amt float32
			fmt.Print("Login successful.\nPlease enter an amount to deposit: ")
			fmt.Scanf("%f", &amt)
			str := fmt.Sprintf("%.2f", amt)
			strconv.ParseFloat(str, 32)
			//fmt.Printf("%.2f\n", amt)
			fmt.Scanf(" %c")
			if amt <= 0.0000 {
				fmt.Println("Invalid amount, please enter a valid amount.")
				goto ENTERAMOUNT
			}
			b.transDates = append(b.transDates, val.Deposit(amt))
			return
		}
	}
	fmt.Println("Incorrect username or password.")
	goto LOGIN
}

func (b *BankApplication) Withdraw() {
	var user string
	var password string
LOGIN: // Login block
	fmt.Print("Username: ")
	fmt.Scanf("%s", &user)
	fmt.Scanf(" %c")
	fmt.Printf("Password: ")
	fmt.Scanf("%s", &password)
	fmt.Scanf(" %c")
	for _, val := range b.accounts {
		if val.userName == user && val.password == password {
			var amt float32
		ENTERAMOUNTWITHDRAWAL:
			fmt.Print("Login successful.\nPlease enter an amount to withdraw: ")
			fmt.Scanf("%f", &amt)
			str := fmt.Sprintf("%.2f", amt)
			strconv.ParseFloat(str, 32)
			fmt.Scanf(" %c")

			if amt > val.balance {
				fmt.Println("You have insufficient funds for this withdrawl.")
				goto ENTERAMOUNTWITHDRAWAL
			} else {
				b.transDates = append(b.transDates, val.Withdraw(amt))
				return
			}
		}
	}
	fmt.Println("Incorrect username or password.")
	goto LOGIN
}

func (b *BankApplication) AddAccount() {
	var username string
	var password string
	fmt.Print("Username of new account: ")
	fmt.Scanf("%s", &username)
	fmt.Scanf(" %c")
	fmt.Print("Password of new account: ")
	fmt.Scanf("%s", &password)
	fmt.Scanf(" %c")
	var acc *BankAccount = &BankAccount{balance: 0.00, userName: username, password: password}
	b.accounts = append(b.accounts[:(len(b.accounts))], acc)

}

func (b *BankApplication) AlterAccount() {
	var user string
	var password string
LOGIN: // Login block
	fmt.Print("Username: ")
	fmt.Scanf("%s", &user)
	fmt.Scanf(" %c")
	fmt.Printf("Password: ")
	fmt.Scanf("%s", &password)
	fmt.Scanf(" %c")
	for _, val := range b.accounts {
		if val.userName == user && val.password == password {
		YES_OR_NO:
			fmt.Print("Would you like to change your username? (Y/N): ")
			fmt.Scanf("%s", &user)
			fmt.Scanf(" %c")
			if user == "Y" || user == "y" {
				fmt.Print("What would you like the new username to be? ")
				fmt.Scanf("%s", &user)
				fmt.Scanf(" %c")
				val.alter(user, val.password)
				fmt.Println("Username successfully changed.")
				return
			} else {
				fmt.Print("Would you like to change your password? (Y/N): ")
				fmt.Scanf("%s", &password)
				fmt.Scanf(" %c")
				if password == "Y" || password == "y" {
					fmt.Print("What would you like the new password to be? ")
					fmt.Scanf("%s", &password)
					fmt.Scanf(" %c")
					val.alter(val.userName, password)
					fmt.Println("Password successfully changed")
					return
				} else {
					fmt.Println("Invalid input, please try again.")
					goto YES_OR_NO
				}
			}
		}
	}
	fmt.Println("Incorrect username or password.")
	goto LOGIN
}

func (b *BankApplication) SortAccountsName() {
	for i := 0; i < len(b.accounts)-1; i++ {
		for j := 0; j < len(b.accounts)-1-i; j++ {
			if strings.Compare(b.accounts[j+1].userName, b.accounts[j].userName) == -1 {
				temp := b.accounts[j+1]
				b.accounts[j+1] = b.accounts[j]
				b.accounts[j] = temp
			}
		}
	}
}

func (b *BankApplication) SortBalance() {
	for i := 0; i < len(b.accounts)-1; i++ {
		for j := 0; j < len(b.accounts)-1-i; j++ {
			if b.accounts[j+1].balance < b.accounts[j].balance {
				temp := b.accounts[j+1]
				b.accounts[j+1] = b.accounts[j]
				b.accounts[j] = temp
			}
		}
	}
}

func (b *BankApplication) PrintAll() {
	for _, val := range b.accounts {
		fmt.Printf("Username:	%s	Balance:	%.2f\n", val.userName, val.balance)
	}
}

func (b *BankApplication) RetrieveAccountInfo() {
	var user string
	var password string
LOGIN: // Login block
	fmt.Print("Username: ")
	fmt.Scanf("%s", &user)
	fmt.Scanf(" %c")
	fmt.Printf("Password: ")
	fmt.Scanf("%s", &password)
	fmt.Scanf(" %c")
	for _, val := range b.accounts {
		if val.userName == user && val.password == password {
			fmt.Println(val.ToString())
			return
		}
	}
	fmt.Println("Incorrect username or password.")
	goto LOGIN
}

func main() {
	var quit bool = false

	ba := NewApplication()
	b := BankAccount{balance: 200.00, userName: "johnsmith3", password: "password3"}
	b2 := BankAccount{balance: 32.45, userName: "johnsmith2", password: "password2"}
	b3 := BankAccount{balance: 4500.99, userName: "johnsmith", password: "password"}
	ba.accounts = append(ba.accounts, &b, &b2, &b3)

	for !quit {
		var choice int16
		fmt.Print("Hello, and thank you for choosing Golang Bank! Please choose from the options below.\n1 - Deposit\n2 - Withdraw\n3 - Create an account\n4 - View account information\n5 - Change account information\n6 - Exit\n")
	READRESPONSE:
		fmt.Scanf("%d", &choice)
		fmt.Scanf(" %c")
		if choice == 1 {
			ba.Deposit()
		} else if choice == 2 {
			ba.Withdraw()
		} else if choice == 3 {
			ba.AddAccount()
		} else if choice == 4 {
			ba.RetrieveAccountInfo()
		} else if choice == 5 {
			ba.AlterAccount()
		} else if choice == 6 {
			break
		} else {
			goto READRESPONSE
		}
	}
	ba.PrintAll()
	// ba.AddAccount()
	// ba.AddAccount()
	// ba.AddAccount()
	ba.SortAccountsName()
	//ba.Withdraw()
	// for _, val := range ba.accounts {
	// 	fmt.Println(val.userName)
	// }
	// b.Withdraw(210.00)
	// b.Deposit(40.00)
	// b.Withdraw(210.00)
}
