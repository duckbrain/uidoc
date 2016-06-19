package main

import (
	"bytes"
	"github.com/andlabs/ui"
	"github.com/duckbrain/uidoc"
	"os"
)

func main() {

	file, err := os.Open("document.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	buffer := new(bytes.Buffer)
	_, err = buffer.ReadFrom(file)
	if err != nil {
		panic(err)
	}
	var document uidoc.Element
	//document = uidoc.NewGroup([]uidoc.Element{})

	// TODO: This is commented out because some bug is causing it to crash
	err = ui.Main(func() {
		document, err = uidoc.Parse(buffer.Bytes())
		if err != nil {
			panic(err)
		}

		font := ui.LoadClosestFont(&ui.FontDescriptor{
			Family: "Deja Vu",
			Size:   12,
		})
		name := ui.NewEntry()
		button := ui.NewButton("Greet")
		doc := uidoc.New()
		doc.SetDocument(document)
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), false)
		box.Append(name, false)
		box.Append(button, false)
		box.Append(doc, true)
		window := ui.NewWindow("Hello", 400, 700, false)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			element := uidoc.NewText("Hello, "+name.Text()+"!", font)
			document.(*uidoc.Group).Append(element)
			doc.Layout()
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
