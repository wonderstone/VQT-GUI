package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("VQT-GUI")
	// * 0. lab for output
	// make a label for output with white color
	la := widget.NewLabel("res")
	// * 1. load the dir for BT
	// make a form for BTDir
	BTdirEntry := widget.NewEntry()
	BTdirEntry.SetPlaceHolder("./config/manual/")
	BTdirEntry.Text = "./config/manual/"

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "BTDir", Widget: BTdirEntry},
		},
	}
	// make a button with a callback change la with BTdirEntry.Text
	btn := widget.NewButton("ReadConfig", func() {
		la.SetText(BTdirEntry.Text)
	})

	// split the form and button
	hs := container.NewHSplit(form, btn)
	hs.SetOffset(0.75)

	// make a gray separator
	sep := canvas.NewRectangle(color.Gray{Y: 128})
	// add hs and sep in a vbox
	BTloadCont := container.NewVBox(hs, sep)
	// * 2. make a form for Instruments and Periods and Indicators
	// make a form for Instruments and Periods and Indicators
	BTform := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Instruments", Widget: widget.NewEntry()},
			{Text: "BeginDate", Widget: widget.NewEntry()},
			{Text: "EndDate", Widget: widget.NewEntry()},
			{Text: "Indicators", Widget: widget.NewEntry()},
		},
	}
	// add a button for DataDownload
	btndownload := widget.NewButton("DataDownload", func() {
	})

	// should read the yaml file and get entries
	entries := []string{"MA3", "Var5"}
	ii, IIform := SelectFactor(entries)
	// make two buttons for SetIndiInfo and preprocess
	btn1 := widget.NewButton("SetIndiInfo", func() {})
	btn2 := widget.NewButton("Preprocess", func() {})
	// split the two buttons
	hs1 := container.NewHSplit(btn1, btn2)
	hs1.SetOffset(0.5)
	// make a gray separator
	sep1 := canvas.NewRectangle(color.Gray{Y: 128})

	// add hs1 and sep1 in a vbox
	BTEntryCont := container.NewVBox(BTform, btndownload, ii, IIform, hs1, sep1)

	// * 3. make a type or select input for backtest
	dd := widget.NewSelect(
		[]string{"DMT", "SortAndBuy"},
		func(s string) { la.SetText(s) },
	)

	// make a content for backtest tab with separator
	content := container.NewVBox(BTloadCont, BTEntryCont, dd, la)
	// add a tabcontainer

	tabs := container.NewAppTabs(

		container.NewTabItem("Backtest", content),
		container.NewTabItem("RealTime", canvas.NewText("Tab2", color.White)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}

func SelectFactor(entries []string) (*widget.Select, *widget.Form) {
	// make a form for IndiInfo
	IIform := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Name", Widget: widget.NewEntry()},
			{Text: "IndiType", Widget: widget.NewEntry()},
			{Text: "ParSlice", Widget: widget.NewEntry()},
			{Text: "InfoSlice", Widget: widget.NewEntry()},
		},
	}

	// * 3. make a type or select input for backtest
	dd := widget.NewSelect(
		entries,
		// read the yaml file and make a form for IndiInfoForm
		func(s string) {},
	)
	return dd, IIform
}
