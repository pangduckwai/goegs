package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type readCfg struct {
	Game struct {
		Variant string
		Players int
	}

	Database struct {
		Pool int
		Conn []struct {
			Host string
			Port string
			Name string
			User string
			Pass string
		}
	}

	Simulations struct {
		Concurrent int
		Runs       int
		Sims       int
		Sleep      int
	}

	Levels  int
	Verbose int
	Profile bool
}

func main() {
	file, err := os.ReadFile("test.yaml")
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		} else {
			fmt.Println("file not exists")
			os.Exit(1)
		}
	}

	read := &readCfg{}
	err = yaml.Unmarshal(file, read)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hmmm... %#v\n", read)
}

/*
game:
  variant: BSCTTTCCDMT5XT0MTD
  players: 3

database:
  pool: 16
  conn:
    - host: localhost
      port: 5432
      name: wdom0
      user: paul
    - name: wdom1

simulations:
  concurrent: 4
  runs: 10
  sims: 10000
  sleep: 1

verbose: 2

profile: true
*/
