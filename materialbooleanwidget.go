package gforms

import (
	"bytes"
)

type materialBooleanInputWidget struct {
	Type  string
	Attrs map[string]string
	Widget
}

type materialBooleanContext struct {
	Checked bool
	Label string
	Type  string
	Field FieldInterface
	Value string
	Attrs map[string]string
}

func (wg *materialBooleanInputWidget) html(f FieldInterface) string {
	var buffer bytes.Buffer
	checked, _ := f.GetV().Value.(bool)
	err := Template.ExecuteTemplate(&buffer, "MaterialBooleanTypeField", materialBooleanContext{
		Label: wg.Attrs["label"],
		Type:  wg.Type,
		Checked: checked,
		Field: f,
		Attrs: wg.Attrs,
		Value: f.GetV().RawStr,
	})
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

func MaterialBooleanInputWidget(attrs map[string]string) Widget {
	w := new(materialBooleanInputWidget)
	w.Type = "checkbox"
	if attrs == nil {
		attrs = map[string]string{}
	}
	w.Attrs = attrs
	return w
}
