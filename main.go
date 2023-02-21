package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("VQT-GUI_Simple_Illustration")
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
	// add a blank form for Strategy info
	StgForm := &widget.Form{
		Items: []*widget.FormItem{},
	}

	dd := widget.NewSelect(
		[]string{"DailyMarketTiming(DMT)", "SortAndBuy(SAB)"},
		func(s string) {
			if s == "DailyMarketTiming(DMT)" {
				StgForm.Items = []*widget.FormItem{
					{Text: "TimeCritic", Widget: widget.NewEntry()},
				}
			}
			if s == "SortAndBuy(SAB)" {
				StgForm.Items = []*widget.FormItem{
					{Text: "TimeCritic", Widget: widget.NewEntry()},
					{Text: "SortType", Widget: widget.NewEntry()},
				}
			}
			// refresh
			StgForm.Refresh()
		},
	)
	// make a button for backtest
	btn3 := widget.NewButton("Run", func() {})
	// make a gray separator
	sep2 := canvas.NewRectangle(color.Gray{Y: 128})
	STGCont := container.NewVBox(dd, StgForm, btn3, sep2)

	// * 4. make a progress bar for backtest

	progress := widget.NewProgressBar()
	var i float64
	go func() {
		for i = 0; i <= 1.0; i += 0.01 {
			time.Sleep(time.Millisecond * 250)
			progress.SetValue(i)
		}
	}()
	// make a lable show the progress
	la2 := widget.NewLabel("BT progress: ")
	// form the progress bar and label
	hs2 := container.NewHSplit(la2, progress)
	hs2.SetOffset(0.1)

	// * 5. draw the result
	// position1 is a point

	line := canvas.NewLine(color.Gray{Y: 128})
	line.Position1 = fyne.NewPos(0, 0)
	line.Position2 = fyne.NewPos(-10, -10)
	line.StrokeWidth = 1

	// put line in a grid
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)),
		line)

	// make a content for backtest tab with separator
	content := container.NewVBox(BTloadCont, BTEntryCont, STGCont, hs2, grid, la)

	// + Realtime 1. load the dir for RT
	// make a form for RTDir
	RTdirEntry := widget.NewEntry()
	RTdirEntry.SetPlaceHolder("./config/manual/")
	RTdirEntry.Text = "./config/manual/"

	RTform := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "RTDir", Widget: RTdirEntry},
		},
	}
	// make a button with a callback change la with BTdirEntry.Text
	RTbtn := widget.NewButton("ReadConfig", func() {

	})

	// add a separator
	sep3 := canvas.NewRectangle(color.Gray{Y: 128})

	// split the form and button
	RThs := container.NewHSplit(RTform, RTbtn)
	RThs.SetOffset(0.75)

	// make a content for backtest tab with separator
	RTDircontent := container.NewVBox(RThs, sep3)
	// add a tabcontainer

	// + Realtime 2. sub instr & indi and cal indi form
	RTsubform := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "RTSubInstr", Widget: widget.NewEntry()},
			{Text: "RTSubInfo", Widget: widget.NewEntry()},
			{Text: "RTCalInfo", Widget: widget.NewEntry()},
		},
	}
	// make two buttons for preload and connect
	RTbtn1 := widget.NewButton("Preload", func() {})
	RTbtn2 := widget.NewButton("Connect", func() {})
	// split the two buttons
	RThs1 := container.NewHSplit(RTbtn1, RTbtn2)
	RThs1.SetOffset(0.5)
	// make a gray separator
	sep4 := canvas.NewRectangle(color.Gray{Y: 128})
	// add hs1 and sep1 in a vbox
	RTCContent := container.NewVBox(RTsubform, RThs1, sep4)

	// + Realtime 3. make two lable to show the button results
	RTla1 := widget.NewLabel("Preload not Done!!!")
	RTla2 := widget.NewLabel("Check the VDS connection!!!")
	// split the two lables
	RThs2 := container.NewHSplit(RTla1, RTla2)
	RThs2.SetOffset(0.5)

	// add a button to preload and conect the VDS
	RTcontent := container.NewVBox(RTDircontent, RTCContent, RThs2)
	tabs := container.NewAppTabs(

		container.NewTabItem("Backtest", content),
		container.NewTabItem("RealTime", RTcontent),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	myWindow.SetContent(tabs)
	myWindow.Resize(fyne.NewSize(600, 800))
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
