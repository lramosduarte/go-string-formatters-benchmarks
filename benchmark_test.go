package main

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

var data = struct {
	Name string
	Age  int
}{
	Name: "João da Silva",
	Age:  17,
}

func BenchmarkFormatSmallStrings(b *testing.B) {
	b.Run("concat strings with plus(+) operator", func(b *testing.B) {
		func() string {
			return data.Name + " tem " + fmt.Sprint(data.Age) + " anos"
		}()
	})
	b.Run("concat strings with fmt() func", func(b *testing.B) {
		func() string {
			return fmt.Sprintf("%s tem %d anos", data.Name, data.Age)
		}()
	})
	b.Run("concat strings with text/template package", func(b *testing.B) {
		func() string {
			rawTemplate, _ := template.New("template").Parse("{{.Name}} tem {{.Age}} anos")
			buff := new(bytes.Buffer)
			rawTemplate.Execute(buff, data)
			return buff.String()
		}()
	})
}
