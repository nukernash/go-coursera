package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	fmt.Println("Please enter name")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println("Please enter address")
	_, err = fmt.Scan(&address)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	mapNameAddress := map[string]string{"name": name, "address": address}
	jsonMap, err := json.Marshal(mapNameAddress)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}
	fmt.Println("JSON ouput:", string(jsonMap))
	return
}
