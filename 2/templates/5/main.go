package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"go", 40},
		{"java", 20},
		{"python", 20},
	})
	if err != nil {
		panic(err)
	}
}
