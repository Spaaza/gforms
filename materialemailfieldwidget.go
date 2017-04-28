package gforms

import (
	"bytes"
)

type materialEmailInputWidget struct {
	Type  string
	Attrs map[string]string
	Widget
}

func (wg *materialEmailInputWidget) html(f FieldInterface) string {
	var buffer bytes.Buffer
	err := Template.ExecuteTemplate(&buffer, "MaterialSimpleWidget", widgetContext{
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

// Generate email input field: <input type="email" ...>
func MaterialEmailInputWidget(attrs map[string]string) Widget {
	w := new(materialEmailInputWidget)
	w.Type = "email"
	if attrs == nil {
		attrs = map[string]string{}
	}
	w.Attrs = attrs
	return w
}
