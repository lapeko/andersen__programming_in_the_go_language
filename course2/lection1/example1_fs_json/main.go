package main

import (
	"encoding/json"
	"log"
	"os"
)

type Users struct {
	Users []User `json:"users"`
}
type User struct {
	Name   string
	Type   string
	Age    int8
	Social Social
}
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	var users Users

	reader, err := os.ReadFile("users_input.json")

	if err != nil {
		log.Fatalln(err)
	}

	if err = json.Unmarshal(reader, &users); err != nil {
		log.Fatalln(err)
	}

	log.Println(users)

	jsonBytes, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		log.Fatalln(err)
	}

	err = os.WriteFile("users_output.json", jsonBytes, 0664)

	if err != nil {
		log.Fatalln(err)
	}

	var untypedUsers map[string]interface{}

	if err = json.Unmarshal(reader, &untypedUsers); err != nil {
		log.Fatalln(err)
	}

	log.Printf("%+v\n", untypedUsers)
}
