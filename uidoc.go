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
	width, height     float64
}

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

func (r *UIDoc) SetDocument(doc Element) {
	r.doc = doc
	r.Layout()
}

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

type drawHandler struct {
	parent *UIDoc
}

func (r *drawHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams) {
	if r.parent.doc == nil {
		return
	}
	r.parent.doc.Render(dp, 0, 0)
}

func (r *drawHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
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
