package uidoc

import (
	"math"

	"github.com/andlabs/ui"
)

// Convenience struct to implement some of the Element Interface with fields
type ElementBase struct {
	MarginLeft, MarginRight, MarginTop, MarginBottom     float64
	PaddingLeft, PaddingRight, PaddingTop, PaddingBottom float64
	BorderWidth                                          float64
	Background, Border                                   *ui.Brush
	LayoutMode                                           LayoutMode
}

func (e *ElementBase) Fill() *ui.Brush {
	return e.Background
}
func (e *ElementBase) Stroke() (*ui.Brush, float64) {
	return e.Border, e.BorderWidth
}
func (e *ElementBase) Mode() LayoutMode {
	return e.LayoutMode
}
func (e *ElementBase) Margins() (top, right, bottom, left float64) {
	return e.MarginTop, e.MarginRight, e.MarginBottom, e.MarginLeft
}
func (e *ElementBase) SetMargins(n float64) {
	e.MarginTop = n
	e.MarginRight = n
	e.MarginBottom = n
	e.MarginLeft = n
}
func (e *ElementBase) Padding() (top, right, bottom, left float64) {
	return e.PaddingTop, e.PaddingRight, e.PaddingLeft, e.PaddingBottom
}
func (e *ElementBase) SetPadding(n float64) {
	e.PaddingTop = n
	e.PaddingRight = n
	e.PaddingBottom = n
	e.PaddingLeft = n
}

type Text struct {
	ElementBase
	Font, lFont *ui.Font
	Text, lText string
	Wrap        bool
	layout      *ui.TextLayout
}

func NewText(text string, font *ui.Font) *Text {
	t := new(Text)
	t.Wrap = true
	t.Text = text
	t.Font = font
	t.layout = ui.NewTextLayout(text, font, 0)
	t.lText = text
	t.lFont = font
	return t
}

func (e *Text) Layout(width float64) (w float64, h float64) {
	if !e.Wrap {
		width = math.MaxFloat64
	}
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
