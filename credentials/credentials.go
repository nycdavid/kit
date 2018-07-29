package credentials

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	User     string `json:"user"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

func ReadFile() Config {
	var cnf Config
	f, err := os.Open(".kit.json")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(f)
	err = dec.Decode(&cnf)
	if err != nil {
		log.Fatal(err)
	}
	return cnf
}
