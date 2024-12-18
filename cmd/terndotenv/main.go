package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error running tern migrate: ", err)
		fmt.Println("Output: ", string(output))
		panic(err)
	}

	fmt.Println("Migration ran successfully: ", string(output))
}