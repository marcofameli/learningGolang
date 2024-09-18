package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

// exemplo 1
//
//	func main() {
//		curso := Curso{"Go", 40}
//		tmp := template.New("CursoTemplate")
//		tmp, _ = tmp.Parse("Curso: {{.Nome}} - Carga Horaria: {{.CargaHoraria}}")
//		err := tmp.Execute(os.Stdout, curso)
//		if err != nil {
//			panic(err)
//		}
//	}
//
// exemplo 2

type Cursos []Curso

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	//t := template.Must(template.New("content.html").ParseFiles(templates...))
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": toUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"Java", 30},
		{"Python", 10},
	})
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", nil)
}

//t := template.Must(template.New("content.html").ParseFiles("content.html"))
//err := t.Execute(os.Stdout, Cursos{
//	{"Go", 40},
//	{"Java", 30},
//	{"C#", 20},
//	{"Python", 10},
//})
//if err != nil {
//	fmt.Println(err)
//}
