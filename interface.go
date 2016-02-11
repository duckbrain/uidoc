package uidoc

import "github.com/andlabs/ui"

type Element interface {
	Fill() *ui.Brush
	Margins() (top, right, bottom, left float64)
	Padding() (top, right, bottom, left float64)
	Mode() LayoutMode
	Layout(width float64) (float64, float64)
	Render(dp *ui.AreaDrawParams, x, y float64)
	Free()
}

type Interacter interface {
	Element
	MouseEvent(me *ui.AreaMouseEvent)
	MouseCrossed(left bool)
	SetFocused(bool)
	OnActivate()
}

type LayoutMode int

const (
	LayoutBlock LayoutMode = iota
	LayoutInline
)

type Alignment int

const (
	AlignLeft Alignment = iota
	AlignCenter
	AlignRight
)
