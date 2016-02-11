package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/andlabs/ui"
	"github.com/duckbrain/uidoc"
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
	document, err := uidoc.Parse(buffer.Bytes())
	if err != nil {
		panic(err)
	}

	err = ui.Main(func() {
		font := ui.LoadClosestFont(&ui.FontDescriptor{
			Family: "Deja Vu",
			Size:   12,
		})
		name := ui.NewEntry()
		button := ui.NewButton("Greet")
		doc := uidoc.NewDoc()
		doc.SetDocument(document)
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), false)
		box.Append(name, false)
		box.Append(button, false)
		box.Append(doc, true)
		window := ui.NewWindow("Hello", 400, 700, false)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			fmt.Printf("Setting name")
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
