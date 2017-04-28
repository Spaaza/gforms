package gforms

import (
	"bytes"
)

type materialDatetimeInputWidget struct {
	Type  string
	Attrs map[string]string
	Widget
}

func (wg *materialDatetimeInputWidget) html(f FieldInterface) string {
	var buffer bytes.Buffer
	err := Template.ExecuteTemplate(&buffer, "MaterialSimpleWidget", widgetContext{
		Type:  wg.Type,
		Label: wg.Attrs["label"],
		Field: f,
		Attrs: wg.Attrs,
		Value: f.GetV().RawStr,
	})
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Generate text input field: <input type="text" ...>
func MaterialDatetimeInputWidget(attrs map[string]string) Widget {
	w := new(materialDatetimeInputWidget)
	w.Type = "text"
	if attrs == nil {
		attrs = map[string]string{}
	}
	w.Attrs = attrs
	return w
}
