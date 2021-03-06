package main

import (
	"errors"
	"fmt"
	"github.com/bluele/gforms"
	"net/http"
	"path"
	"reflect"
	"runtime"
	"text/template"
)

type Lang struct {
	Name string `gforms:"name"`
}

type CustomValidator struct {
	Langs []string
	gforms.Validator
}

func (vl CustomValidator) Validate(fi *gforms.FieldInstance, fo *gforms.FormInstance) error {
	v := fi.V
	if v.IsNil || v.Kind != reflect.String || v.Value == "" {
		return nil
	}
	for _, t := range vl.Langs {
		if v.Value == t {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unknown lang: %v", v.Value))
}

func main() {
	tpl := template.Must(template.ParseFiles(path.Join(getTemplatePath(), "post_form.html")))
	langForm := gforms.DefineModelForm(Lang{}, gforms.NewFields(
		gforms.NewTextField(
			"name",
			gforms.Validators{
				gforms.Required(),
				CustomValidator{
					Langs: []string{"golang", "python", "c"},
				},
			},
		),
	))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		form := langForm(r)
		if r.Method != "POST" {
			tpl.Execute(w, form)
			return
		}
		if !form.IsValid() {
			tpl.Execute(w, form)
			return
		}
		lang := form.GetModel().(Lang)
		fmt.Fprintf(w, "ok: %v", lang)
	})
	http.ListenAndServe(":9000", nil)
}

func getTemplatePath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), "templates")
}
