package gforms

import (
	"bytes"
)

type materialTextInputWidget struct {
	Type  string
	Attrs map[string]string
	Widget
}

func (wg *materialTextInputWidget) html(f FieldInterface) string {
	var buffer bytes.Buffer
	err := Template.ExecuteTemplate(&buffer, "MaterialTextTypeField", widgetContext{
		Label: wg.Attrs["label"],
		Type:  wg.Type,
		Field: f,
		Attrs: wg.Attrs,
		Value: f.GetV().RawStr,
	})
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Generate text input fiele: <input type="text" ...>
func MaterialTextInputWidget(attrs map[string]string) Widget {
	w := new(materialTextInputWidget)
	w.Type = "text"
	if attrs == nil {
		attrs = map[string]string{}
	}
	w.Attrs = attrs
	return w
}
