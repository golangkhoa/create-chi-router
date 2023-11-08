package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a chi project",
	Long: `Generates a chi project effortlessly.
For example:

create-chi-router generate -n PROJECT_NAME`,
	Run: generateProject,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringP("name", "n", "chi-router", "Name of the project")
}

func generateProject(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	command := exec.Command("go", "mod", "init", name)
	command.Dir = name + "/"
	stdout, err := command.Output()

	file, err := os.Create(name + "/" + "main.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.WriteString(`package main

import "fmt"
	
func main() {
	fmt.Println("Hello world!")
}
`)

	fmt.Sprintln(string(stdout))
}
