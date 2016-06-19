package uidoc

import (
	"fmt"

	"github.com/andlabs/ui"
)

var _ fmt.Formatter

type UIDoc struct {
	area          *ui.Area
	doc           Element
	focus         Interacter
	focusLayout   layout
	width, height float64
}

// Creates a new UIDoc control. This can be added to a ui.Window, ui.Box, etc.
func New() *UIDoc {
	r := &UIDoc{}
	r.area = ui.NewScrollingArea(&drawHandler{r}, -1, 0)
	toolbarContainer := ui.NewHorizontalBox()
	toolbarHeightRetainer := ui.NewHorizontalSeparator()
	toolbarContainer.Append(toolbarHeightRetainer, false)
	return r
}

// Frees the document if it is not nil. This is not done in SetDocument in case you want to reuse a document.
func (r *UIDoc) Free() {
	if r.doc != nil {
		r.doc.Free()
	}
}

// Gets the document displayed
func (r *UIDoc) Document() Element {
	return r.doc
}

// Sets the document displayed, lays it out, and renders it
func (r *UIDoc) SetDocument(doc Element) {
	r.doc = doc
	r.focus = nil
	r.Layout()
}

// Triggers document Layout. This should be called after modifying elements in
// a way that would affect layout. It is also called when a document area
// resizes.
func (r *UIDoc) Layout() {
	r.layout(r.width)
}

func (r *UIDoc) layout(width float64) {
	if r.doc == nil {
		return
	}
	_, h := r.doc.Layout(width)
	r.width = width
	r.height = h
	r.area.QueueRedrawAll()
	r.area.SetSize(-1, int(r.height))
	r.area.QueueRedrawAll()
}

// Wrap the area to make this element behave as a control

func (r *UIDoc) Destroy() {
	r.area.Destroy()
}
func (r *UIDoc) LibuiControl() uintptr {
	return r.area.LibuiControl()
}
func (r *UIDoc) Handle() uintptr {
	return r.area.Handle()
}
func (r *UIDoc) Show() {
	r.area.Show()
}
func (r *UIDoc) Hide() {
	r.area.Hide()
}
func (r *UIDoc) Enable() {
	r.area.Enable()
}
func (r *UIDoc) Disable() {
	r.area.Disable()
}

func (r *UIDoc) handleGroupMouse(g *Group, me *ui.AreaMouseEvent) {
	hit := func(l layout, me *ui.AreaMouseEvent) bool {
		top, right, bottom, left := l.e.Padding()
		return me.X >= l.x-left && me.X <= l.x+l.w+left+right && me.Y >= l.y-top && me.Y <= l.y+l.h+top+bottom
	}

	if r.focus != nil {
		if me.Down == 1 && r.focus != nil {
			r.focus.SetPressed(true)
			r.area.QueueRedrawAll()
			return
		}

		if me.Up == 1 && r.focus != nil {
			if hit(r.focusLayout, me) {
				r.focus.SetPressed(false)
				r.focus.OnActivate()
				r.area.QueueRedrawAll()
			}
			return
		}

		for _, button := range me.Held {
			if button == 1 {
				r.focus.SetPressed(hit(r.focusLayout, me))
				r.area.QueueRedrawAll()
				return
			}
		}
	}

	for _, l := range g.layouts {
		i, ok := l.e.(Interacter)
		if !ok {
			continue // Skip non-Interactors
		}
		if !ok || l.y > me.Y { // if it is too far down, stop
			return
		}
		if hit(l, me) {
			if r.focus == i {
				return
			}
			defer r.area.QueueRedrawAll()
			if r.focus != nil {
				r.focus.SetFocused(false)
			}
			r.focusLayout = l
			r.focus = i
			i.SetFocused(true)
			return
		}
	}
}

type drawHandler struct {
	parent *UIDoc
}

func (r *drawHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams) {
	if r.parent.doc == nil {
		return
	}
	if r.parent.width != dp.AreaWidth {
		r.parent.layout(dp.AreaWidth)
	}
	r.parent.doc.Render(dp, 0, 0)
}

func (r *drawHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	if r.parent.doc == nil {
		return
	}
	if g, ok := r.parent.doc.(*Group); ok {
		r.parent.handleGroupMouse(g, me)
	}
}

func (r *drawHandler) MouseCrossed(a *ui.Area, left bool) {
}
func (r *drawHandler) DragBroken(a *ui.Area) {
}
func (r *drawHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	return false
}
