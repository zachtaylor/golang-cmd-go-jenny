package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/template"

	"taylz.io/env"
)

// Version is the module version number
const Version = "v0.0.0"

// newenv returns env.Values that are required
func newenv() env.Values {
	return env.Values{
		"p": "",
		"k": "",
		"v": "",
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print(Usage)
		return
	} else if os.Args[1] == "version" {
		fmt.Println("go-jenny version", Version)
		return
	}

	def := newenv()
	env := newenv().ParseArgs(os.Args[1:])

	if len(env["h"]) > 0 || len(env["help"]) > 0 {
		fmt.Print(Usage)
		return
	}

	for k := range env {
		if env[k] == def[k] {
			fmt.Println("go-jenny: requires:", k)
			fmt.Print(Usage)
			return
		}
	}

	fileName := "jenny.go"

	if f := env["f"]; f != "" {
		fileName = f
	}

	if _, err := os.Stat(fileName); err == nil {
		os.Remove(fileName)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("go-jenny:", err)
		return
	}
	defer file.Close()

	stdlib := []string{"\"sync\""}
	remote := []string{}

	for _, i := range strings.Split(env["i"], ",") {
		if len(i) < 1 {
			continue
		}
		if isep := strings.Index(i, " "); isep < 0 {
			i = escape(i)
		} else {
			i = i[0:isep] + " " + escape(i[isep+1:])
		}
		if !strings.Contains(i, ".") {
			stdlib = append(stdlib, i)
		} else {
			remote = append(remote, i)
		}
	}

	sort.Strings(stdlib)
	sort.Strings(remote)

	typeName := "Map"

	if t := env["t"]; t != "" {
		typeName = t
	}

	val := env["v"]

	tpl := template.Must(template.New("").Parse(Template))
	data := Options{
		Package: env["p"],
		Type:    typeName,
		Key:     env["k"],
		Val:     val,
		Off:     defaultValue(val),
		Stdlib:  stdlib,
		Remote:  remote,
	}
	if err := tpl.Execute(file, data); err != nil {
		fmt.Println("go-jenny:", err)
	}
}

func escape(str string) string {
	if len(str) < 2 {
		return ""
	} else if str[0] != '"' {
		str = "\"" + str
	}
	if str[len(str)-1] != '"' {
		str = str + "\""
	}
	return str
}

func defaultValue(t string) string {
	switch t {
	case "bool":
		return "false"
	case "int":
		fallthrough
	case "uint":
		fallthrough
	case "float32":
		fallthrough
	case "float64":
		return "0"
	case "string":
		return `""`
	default:
		return "nil"
	}
}
