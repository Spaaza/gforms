package gforms

import (
	"bytes"
)

type materialSelectWidget struct {
	Multiple bool
	Attrs    map[string]string
	Maker    SelectOptionsMaker
	Widget
}

type materialSelectContext struct {
	Multiple bool
	Label    string
	Field    FieldInterface
	Attrs    map[string]string
	Options  selectOptionsValues
}

func (wg *materialSelectWidget) html(f FieldInterface) string {
	var buffer bytes.Buffer
	context := new(materialSelectContext)
	context.Label = wg.Attrs["label"]
	context.Field = f
	context.Multiple = wg.Multiple
	opts := wg.Maker()
	for i := 0; i < opts.Len(); i++ {
		sel := false
		if f.GetV().RawStr == opts.Value(i) {
			sel = true
		} else if len(f.GetV().RawStr) == 0 {
			sel = opts.Selected(i)
		}
		context.Options = append(context.Options, &selectOptionValue{Label: opts.Label(i), Value: opts.Value(i), Selected: sel, Disabled: opts.Disabled(i)})
	}
	context.Attrs = wg.Attrs
	err := Template.ExecuteTemplate(&buffer, "MaterialSelectWidget", context)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Generate select and options field: <select><option></option></select>
func MaterialSelectWidget(attrs map[string]string, mk SelectOptionsMaker) *materialSelectWidget {
	wg := new(materialSelectWidget)
	if attrs == nil {
		attrs = map[string]string{}
	}
	if isNilValue(mk) {
		mk = func() SelectOptions {
			return StringSelectOptions([][]string{})
		}
	}
	wg.Maker = mk
	wg.Attrs = attrs
	return wg
}

// Generate select-multiple and options field: <select multiple><option></option></select>
func MaterialSelectMultipleWidget(attrs map[string]string, mk SelectOptionsMaker) *materialSelectWidget {
	wg := MaterialSelectWidget(attrs, mk)
	wg.Multiple = true
	return wg
}
