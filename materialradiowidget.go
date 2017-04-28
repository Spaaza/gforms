package gforms

import (
	"bytes"
)

type materialRadioSelectWidget struct {
	Attrs map[string]string
	Maker RadioOptionsMaker
	Widget
}
//
//type radioOptionValue struct {
//	Label    string
//	Value    string
//	Checked  bool
//	Disabled bool
//}
//
//type radioOptionValues []*radioOptionValue
//
type materialRadioContext struct {
	Label   string
	Field   FieldInterface
	Attrs   map[string]string
	Options radioOptionValues
}
//
//type RadioOptionsMaker func() RadioOptions
//
//type RadioOptions interface {
//	Label(int) string
//	Value(int) string
//	Checked(int) bool
//	Disabled(int) bool
//	Len() int
//}
//
//type StringRadioOptions [][]string
//
//func (opt StringRadioOptions) Label(i int) string {
//	return opt[i][0]
//}
//
//func (opt StringRadioOptions) Value(i int) string {
//	return opt[i][1]
//}
//
//func (opt StringRadioOptions) Checked(i int) bool {
//	checked := opt[i][2]
//	if checked == "true" {
//		return true
//	} else {
//		return false
//	}
//}
//
//func (opt StringRadioOptions) Disabled(i int) bool {
//	disabled := opt[i][3]
//	if disabled == "true" {
//		return true
//	} else {
//		return false
//	}
//}
//
//func (opt StringRadioOptions) Len() int {
//	return len(opt)
//}

func (wg *materialRadioSelectWidget) html(f FieldInterface) string {
	var buffer bytes.Buffer
	ctx := new(materialRadioContext)
	ctx.Label = wg.Attrs["label"]
	opts := wg.Maker()
	for i := 0; i < opts.Len(); i++ {
		ctx.Options = append(
			ctx.Options,
			&radioOptionValue{
				Label:    opts.Label(i),
				Value:    opts.Value(i),
				Checked:  opts.Checked(i),
				Disabled: opts.Disabled(i),
			})
	}
	ctx.Field = f
	ctx.Attrs = wg.Attrs
	err := Template.ExecuteTemplate(&buffer, "MaterialRadioWidget", ctx)
	if err != nil {
		panic(err)
	}
	return buffer.String()
}

// Generate radio input field: <input type="radio" ...>
func MaterialRadioSelectWidget(attrs map[string]string, mk RadioOptionsMaker) *materialRadioSelectWidget {
	wg := new(materialRadioSelectWidget)
	wg.Attrs = attrs
	wg.Maker = mk
	return wg
}
