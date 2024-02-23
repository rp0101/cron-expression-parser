package main

import (
	"cron-expression-parser/controller"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Please enter the cron string in correct format")
		os.Exit(1)
	}
	cronStr := os.Args[1]
	mController := controller.NewCronController()
	expandedFields, err := mController.ParseCronController(cronStr)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// Format as a table
	fmt.Printf("%-14s%s\n", "Field", "Times")
	for field, values := range expandedFields {
		fmt.Printf("%-14s%s\n", field, values)
	}
}
