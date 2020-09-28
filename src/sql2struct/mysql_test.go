package sql2struct

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	info := &DbInfo{
		UserName: "root",
		Password: "ankang123",
		Host:     "localhost",
		Port:     "3306",
		Charset:  "utf8mb4",
		DbType:   "mysql",
	}

	model := NewDbModel(info)
	err := model.Connect()
	if err != nil {
		fmt.Println("error", err)
	}

	columns, err := model.GetColumns("ankang", "blog_tag")
	if err != nil {
		fmt.Println("error", err)
	}

	template := NewStructTemplate()
	templateColumns := template.AssemblyColumns(columns)
	err = template.Generate("blog_tag", templateColumns)
	if err != nil {
		fmt.Println("error", err)
	}
}
