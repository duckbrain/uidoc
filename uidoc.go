package uidoc

import (
	"fmt"

	"github.com/andlabs/ui"
)

var _ fmt.Formatter

type UIDoc struct {
	area, measureArea *ui.Area
	box               *ui.Box
	doc               Element
	focus             Interacter
	focusLayout       layout
	width, height     float64
}

// Creates a new UIDoc control. This can be added to a ui.Window, ui.Box, etc.
func New() *UIDoc {
	r := &UIDoc{}
	r.area = ui.NewScrollingArea(&drawHandler{r}, 400, 400)
	r.measureArea = ui.NewArea(&measureHandler{r})
	r.box = ui.NewVerticalBox()
	toolbarContainer := ui.NewHorizontalBox()
	toolbarHeightRetainer := ui.NewHorizontalSeparator()
	toolbarContainer.Append(toolbarHeightRetainer, false)
	toolbarContainer.Append(r.measureArea, true)
	r.box.Append(r.area, true)
	r.box.Append(toolbarContainer, false)
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
	r.area.SetSize(int(r.width), int(r.height))
	r.area.QueueRedrawAll()
}

// Wrap the box to make this element behave as a control

func (r *UIDoc) Destroy() {
	r.box.Destroy()
}
func (r *UIDoc) LibuiControl() uintptr {
	return r.box.LibuiControl()
}
func (r *UIDoc) Handle() uintptr {
	return r.box.Handle()
}
func (r *UIDoc) Show() {
	r.box.Show()
}
func (r *UIDoc) Hide() {
	r.box.Hide()
}
func (r *UIDoc) Enable() {
	r.box.Enable()
}
func (r *UIDoc) Disable() {
	r.box.Disable()
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
	if f := r.parent.doc.Fill(); f != nil {
		p := ui.NewPath(ui.Winding)
		p.AddRectangle(dp.ClipX, dp.ClipY, dp.ClipWidth, dp.ClipHeight)
		p.End()
		dp.Context.Fill(p, f)
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

type measureHandler struct {
	parent *UIDoc
}

func (r *measureHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams) {
	if r.parent.width != dp.AreaWidth {
		fmt.Sprintf("Size changed: %v\n", dp.AreaWidth)
		r.parent.layout(dp.AreaWidth)
	}

	//Fill background
	p := ui.NewPath(ui.Winding)
	p.AddRectangle(0, 0, dp.AreaWidth, dp.AreaHeight)
	p.End()
	dp.Context.Fill(p, &ui.Brush{
		A: 1,
	})
}
func (r *measureHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
}
func (r *measureHandler) MouseCrossed(a *ui.Area, left bool) {
}
func (r *measureHandler) DragBroken(a *ui.Area) {
}
func (r *measureHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) bool {
	return false
}
