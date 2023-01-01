package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Name    string
	Address string
}

func main() {

	var name string
	var address string

	person := make(map[string]string)

	in := bufio.NewReader(os.Stdin)

	fmt.Println("Enter a Name : ")
	name, err := in.ReadString('\n')

	fmt.Println("Enter a Address : ")
	address, err = in.ReadString('\n')
	person["name"] = strings.TrimSpace(name)
	person["address"] = strings.TrimSpace(address)

	barr, err := json.Marshal(person)

	if err != nil {
		fmt.Printf("Error occured : %v\n", err)
	}

	fmt.Printf("json : %s\n", string(barr))

}
