package uidoc

import "github.com/andlabs/ui"

type Button struct {
	BackgroundNormal, BackgroundFocus, BackgroundPress *ui.Brush
	BorderWidth                                        float64
	Border                                             *ui.Brush
	InnerElement                                       Element
	Callback                                           func()
	focused, pressed                                   bool
}

func NewButton(e Element, callback func()) *Button {
	b := new(Button)
	b.InnerElement = e
	b.Callback = callback

	b.BackgroundNormal = &ui.Brush{
		R: 0.8,
		G: 0.8,
		B: 0.8,
		A: 1,
	}
	b.BackgroundPress = &ui.Brush{
		R: 0.6,
		G: 0.6,
		B: 0.6,
		A: 1,
	}
	b.BackgroundFocus = &ui.Brush{
		R: 0.9,
		G: 0.9,
		B: 0.9,
		A: 1,
	}
	b.Border = &ui.Brush{A: 1}
	b.BorderWidth = 1.0

	return b
}

func (e *Button) Fill() *ui.Brush {
	if e.focused {
		if e.pressed {
			return e.BackgroundPress
		} else {
			return e.BackgroundFocus
		}
	} else {
		return e.BackgroundNormal
	}
}
func (e *Button) Stroke() (*ui.Brush, float64) {
	return e.Border, e.BorderWidth
}
func (e *Button) Margins() (top, right, bottom, left float64) {
	return e.InnerElement.Margins()
}
func (e *Button) Padding() (top, right, bottom, left float64) {
	return e.InnerElement.Padding()
}
func (e *Button) Mode() LayoutMode {
	return e.InnerElement.Mode()
}

func (e *Button) Layout(width float64) (w float64, h float64) {
	// TODO Pass layout through
	return e.InnerElement.Layout(width)
}

func (e *Button) Render(dp *ui.AreaDrawParams, x, y float64) {
	// TODO Pass layout through
	e.InnerElement.Render(dp, x, y)
}

func (e *Button) Free() {
	e.InnerElement.Free()
}

func (e *Button) SetFocused(b bool) {
	e.focused = b
}
func (e *Button) SetPressed(b bool) {
	e.pressed = b
}
func (e *Button) OnActivate() {
	if e.Callback != nil {
		e.Callback()
	}
}
