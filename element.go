package uidoc

import "github.com/andlabs/ui"

// Convenience struct to implement some of the Element Interface with fields
type ElementBase struct {
	MarginLeft, MarginRight, MarginTop, MarginBottom     float64
	PaddingLeft, PaddingRight, PaddingTop, PaddingBottom float64
	Background                                           *ui.Brush
	LayoutMode                                           LayoutMode
}

func (e *ElementBase) Fill() *ui.Brush {
	return e.Background
}
func (e *ElementBase) Mode() LayoutMode {
	return e.LayoutMode
}
func (e *ElementBase) Margins() (top, right, bottom, left float64) {
	return e.MarginTop, e.MarginRight, e.MarginBottom, e.MarginLeft
}
func (e *ElementBase) Padding() (top, right, bottom, left float64) {
	return e.PaddingTop, e.PaddingRight, e.PaddingLeft, e.PaddingBottom
}

type Text struct {
	ElementBase
	Font, lFont *ui.Font
	Text, lText string
	layout      *ui.TextLayout
}

func NewText(text string, font *ui.Font) *Text {
	t := new(Text)
	t.Text = text
	t.Font = font
	t.layout = ui.NewTextLayout(text, font, 0)
	t.lText = text
	t.lFont = font
	return t
}

func (e *Text) Layout(width float64) (w float64, h float64) {
	if e.Text != e.lText || e.Font != e.lFont {
		e.layout.Free()
		e.layout = ui.NewTextLayout(e.Text, e.Font, width)
		e.lText = e.Text
		e.lFont = e.Font
	} else {
		e.layout.SetWidth(width)
	}
	return e.layout.Extents()
}

func (e *Text) Render(dp *ui.AreaDrawParams, x, y float64) {
	dp.Context.Text(x, y, e.layout)
}

func (e *Text) Free() {
	e.layout.Free()
}
