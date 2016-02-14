package uidoc

import "github.com/andlabs/ui"

type Element interface {
	Fill() *ui.Brush
	Stroke() (*ui.Brush, float64)
	Margins() (top, right, bottom, left float64)
	Padding() (top, right, bottom, left float64)
	Mode() LayoutMode
	Layout(width float64) (float64, float64)
	Render(dp *ui.AreaDrawParams, x, y float64)
	Free()
}

type Interacter interface {
	Element
	SetPressed(b bool)
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
