package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"

	"path/filepath"

	"github.com/stoewer/go-strcase"
)

var tpl *template.Template
var fm template.FuncMap = template.FuncMap{
	"ucc": func(v string) string {
		return strcase.UpperCamelCase(v)
	},
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gotpl"))
}

type Model struct {
	Name   string
	Table  string
	Fields []string
}

type Helper struct {
	Fields []string
}

func (h *Helper) String() string {
	return fmt.Sprint(h.Fields)
}

func (h *Helper) Set(value string) error {
	fields := strings.Split(value, ",")
	for _, v := range fields {
		h.Fields = append(h.Fields, v)
	}
	return nil
}

func main() {
	var (
		n string
		t string
		h Helper
		d string
	)
	flag.StringVar(&n, "model", "", "nombre del modelo (ej: role)")
	flag.StringVar(&t, "table", "", "nombre de la tabla (ej: roles)")
	flag.Var(&h, "fields", "nombre de los campos de la tabla separados por coma sin espacios (ej: name,phone,address,age)")
	flag.StringVar(&d, "dest", "dist", "nombre del directorio destino (ej: roles)")
	flag.Parse()

	if n == "" {
		log.Fatalln("el modelo es obligatorio: -model=nombre_modelo")
		return
	}
	if t == "" {
		log.Fatalln("el nombre de la tabla es obligatorio: -table=nombre_tabla")
		return
	}
	if len(h.Fields) == 0 {
		flag.PrintDefaults()
		log.Fatalln("el listado de los campos de la tabla es obligatorio: -fields=campo1,campo2,campo3")
		return
	}

	m := Model{n, t, h.Fields}

	fmt.Printf("el valor de d es: %s\n", d)

	createDir(d)
	generateModel(m, d)
	generateStorage(m, d)
	generateSqlServer(m, d)
	generateHandler(m, d)
}

func createDir(d string) {
	_, err := os.Stat(d)
	if os.IsNotExist(err) {
		log.Printf("no existe la carpeta %s. Creandola...", d)
		os.Mkdir(d, os.ModePerm)
	}
}

func generateModel(m Model, d string) {
	generateTemplate(filepath.Join(d, "model.go"), "model.gotpl", m)
}

func generateStorage(m Model, d string) {
	generateTemplate(filepath.Join(d, "storage.go"), "storage.gotpl", m)
}

func generateSqlServer(m Model, d string) {
	generateTemplate(filepath.Join(d, "sqlserver.go"), "sqlserver.gotpl", m)
}

func generateHandler(m Model, d string) {
	generateTemplate(filepath.Join(d, "handler.go"), "handler.gotpl", m)
}

func generateTemplate(dest, source string, m Model) {
	f, err := os.OpenFile(dest, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("no se pudo crear el archivo: %v", err)
	}
	defer f.Close()

	err = tpl.ExecuteTemplate(f, source, m)
	if err != nil {
		log.Printf("error creando el archivo: %v", err)
		return
	}
}
