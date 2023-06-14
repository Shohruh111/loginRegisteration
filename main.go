package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *User) Login(username, password string) bool {
	if l.Username == username && l.Password == password {
		return true
	}
	return false
}

func (r *User) Registartion(user *User) *User {
	r.Password = user.Password
	r.Username = user.Username
	return r
}
func main() {
	var (
		choice   int
		username string
		password string
	)
	var reg []User

	file, err := os.Open("data.json")
	if err != nil {
		log.Println(err)
		return
	}
	err = json.NewDecoder(file).Decode(&reg)
	if err != nil {
		log.Print("Error!!!")
		return
	}

STEP1:
	fmt.Println("\t\t\t\t1. Login")
	fmt.Println("\t\t\t\t2.Registartion")
	fmt.Scan(&choice)

	if choice == 1 {
	loginStep:
		fmt.Print("Enter username: ")
		fmt.Scan(&username)
		fmt.Print("Enter password: ")
		fmt.Scan(&password)
		for _, user := range reg {
			if user.Login(username, password) {
				fmt.Println("Successfully Login!")
				return
			}
		}
		fmt.Println("Please, enter right password and username!!!")
		goto loginStep
	} else if choice == 2 {
		fmt.Print("Enter a username for registration:  ")
		fmt.Scan(&username)

		fmt.Print("Enter a password for registration:  ")
		fmt.Scan(&password)
		user := User{username, password}
		reg = append(reg, user)

		body, err := json.MarshalIndent(reg, "", "	")
		if err != nil {
			log.Print("error while writing file!!!", err)
			return
		}
		err = ioutil.WriteFile("data.json", body, os.ModePerm)
		fmt.Println("Succesfully registred!")
		goto STEP1
	} else {
		goto STEP1
	}
}
