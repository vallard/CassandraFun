package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {

	fmt.Printf("Test\n")
	cluster := gocql.NewCluster("192.168.99.100")
	cluster.Port = 32869
	cluster.ProtoVersion = 4
	cluster.Keyspace = "ucstechspecs"
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	m := make(map[string]string)
	m["color"] = "red"
	m["car"] = "ferrari"

	pictures := []string{"http://fi.1.org", "http://lmgtfy"}

	err = session.Query(`INSERT INTO fabricInterconnects (partNumber, name, attributes, pictures) VALUES (?, ?, ?, ?)`,
		"UCSFI-61234", "UCS FI 6120", m, pictures).Exec()

	log.Println("Inserted data")
	if err != nil {
		log.Fatal(err)
	}

	var model string
	iter := session.Query(`SELECT name FROM fabricInterconnects WHERE partNumber = ? `,
		"UCSFI-61234").Iter()
	for iter.Scan(&model) {
		fmt.Println("Model: ", model)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}
