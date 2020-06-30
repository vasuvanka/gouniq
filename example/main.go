package main

import (
	"fmt"
	"log"
	"github.com/vasuvanka/gouniq"
)

func main()  {
	ID, err := gouniq.NewID(10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ID)
}
