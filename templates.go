package gforms

import (
	"bytes"
	"text/template"
)

const defaultTemplates = `
{{define "TextTypeField"}}<input type="text" name="{{.Field.GetName | html}}" value="{{.Value | html}}"></input>{{end}}
{{define "BooleanTypeField"}}<input type="checkbox" name="{{.Field.GetName | html}}"{{if .Checked}} checked{{end}}>{{end}}
{{define "SimpleWidget"}}<input type="{{.Type | html}}" name="{{.Field.GetName | html}}" value="{{.Value | html}}"{{range $attr, $val := .Attrs}} {{$attr | html}}="{{$val | html}}"{{end}}></input>{{end}}
{{define "SelectWidget"}}<select {{if .Multiple }}multiple {{end}}name="{{.Field.GetName | html}}"{{range $attr, $val := .Attrs}}{{$attr | html}}="{{$val | html}}"{{end}}>
{{range $idx, $val := .Options}}<option value="{{$val.Value | html}}"{{if $val.Selected }} selected{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label | html}}</option>
{{end}}</select>{{end}}
{{define "RadioWidget"}}{{$name := .Field.GetName}}{{range $idx, $val := .Options}}<input type="radio" name="{{$name | html}}" value="{{$val.Value | html}}"{{if or $val.Checked (eq $.Field.GetV.RawStr $val.Value) }} checked{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label | html}}
{{end}}{{end}}
{{define "CheckboxMultipleWidget"}}{{$name := .Field.GetName}}{{range $idx, $val := .Options}}<input type="checkbox" name="{{$name | html}}" value="{{$val.Value | html}}"{{if $val.Checked}} checked{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label | html}}
{{end}}{{end}}
{{define "MaterialTextTypeField"}}<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label"><input class="mdl-textfield__input" type="text" id="textfield-{{.Field.GetName | html}}" name="{{.Field.GetName | html}}" value="{{.Value | html}}"><label class="mdl-textfield__label" for="textfield-{{.Field.GetName | html}}">{{.Label | html}}</label></div>{{end}}
{{define "MaterialBooleanTypeField"}}<div class="mdl-textfield mdl-js-textfield"><label class="mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect" for="checkbox-{{.Field.GetName | html}}"><input type="checkbox" id="checkbox-{{.Field.GetName | html}}" class="mdl-checkbox__input" name="{{.Field.GetName | html}}"{{if .Checked}} checked{{end}}><span class="mdl-checkbox__label">{{.Label | html}}</span></label></div>{{end}}
{{define "MaterialSimpleWidget"}}<div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label"><input class="mdl-textfield__input" type="text" id="textfield-{{.Field.GetName | html}}" name="{{.Field.GetName | html}}" value="{{.Value | html}}"{{range $attr, $val := .Attrs}} {{$attr | html}}="{{$val | html}}"{{end}}><label class="mdl-textfield__label" for="textfield-{{.Field.GetName | html}}">{{.Label | html}}</label></div>{{end}}
{{define "MaterialSelectWidget"}}<div class="mdl-selectfield mdl-js-selectfield mdl-selectfield--floating-label"><select class="mdl-selectfield__select" {{if .Multiple }}multiple {{end}}name="{{.Field.GetName | html}}"{{range $attr, $val := .Attrs}}{{$attr | html}}="{{$val | html}}"{{end}}>{{range $idx, $val := .Options}}<option value="{{$val.Value | html}}"{{if $val.Selected }} selected{{end}}{{if $val.Disabled}} disabled{{end}}>{{$val.Label | html}}</option>{{end}}</select><div class="mdl-selectfield__icon"><i class="material-icons">arrow_drop_down</i></div><label class="mdl-selectfield__label" for="{{.Field.GetName | html}}">{{.Label | html}}</label></div>{{end}}
{{define "MaterialRadioWidget"}}{{$name := .Field.GetName}}<div class="{{$name}} mdl-textfield mdl-js-textfield mdl-textfield--floating-label">{{range $idx, $val := .Options}}<label class="mdl-radio mdl-js-radio mdl-js-ripple-effect" for="radio-{{$val.Value | html}}"><input type="radio" id="radio-{{$val.Value | html}}" class="mdl-radio__button" name="{{$name | html}}" value="{{$val.Value | html}}"{{if or $val.Checked (eq $.Field.GetV.RawStr $val.Value) }} checked{{end}}{{if $val.Disabled}} disabled{{end}}><span class="mdl-radio__label">{{$val.Label | html}}</span></label>{{end}}<label class="mdl-textfield__label">{{.Label | html}}</label></div>{{end}}
`

// all templates of Field and Widget
var Template *template.Template

func init() {
	var err error
	Template, err = template.New("gforms").Parse(defaultTemplates)
	if err != nil {
		panic(err)
	}
}

type templateContext struct {
	Field FieldInterface
	Value string
}

func newTemplateContext(f FieldInterface) templateContext {
	ctx := templateContext{
		Field: f,
	}
	v := f.GetV()
	if v != nil {
		ctx.Value = v.RawStr
	}
	return ctx
}

func renderTemplate(name string, ctx interface{}) string {
	var buffer bytes.Buffer
	err := Template.ExecuteTemplate(&buffer, name, ctx)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}
