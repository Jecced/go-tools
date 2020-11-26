package test

import (
	"fmt"
	"github.com/Jecced/go-tools/src/strutil"
	"os"
	"strings"
	"testing"
	"text/template"
)

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	o := strutil.UnderscoreToUpperCamelCase("user_name")
	fmt.Println(o)
	o = strutil.CamelCaseToUnderscore(o)
	fmt.Println(o)
}

const templateText = `
Output 0: {{title .Name1}}
Output 1: {{title .Name2}}
Output 2: {{.Name3 | title}}
`

func TestStr(t *testing.T) {
	funcMap := template.FuncMap{"title": strings.Title}
	tpl := template.New("mysql_test")
	tpl, _ = tpl.Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go1",
		"Name2": "go2",
		"Name3": "go3",
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
