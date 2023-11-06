package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	project_name := os.Args[1]

	err := os.Mkdir(project_name, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	cmd := exec.Command("go", "mod", "init", project_name)
	cmd.Dir = project_name + "/"
	stdout, err := cmd.Output()

	file, err := os.Create(project_name + "/" + "main.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.WriteString(`package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	http.ListenAndServe(":3000", r)
}
`)

	fmt.Sprintln(string(stdout))
}
