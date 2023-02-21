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
	"github.com/wonderstone/VQT-GUI/backtest"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("VQT-GUI_V0.1")
	// * Backtest： 0. lab for final result output
	// make a label for output with white color
	laRes := widget.NewLabel("This is the final result:")
	// * 1. load the dir for BT
	// make a form for BTDir
	EntryBTdir := widget.NewEntry()
	EntryBTdir.SetPlaceHolder("./config/manual/")
	EntryBTdir.Text = "./config/manual/"

	formBTdir := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "BTDir", Widget: EntryBTdir},
		},
	}

	// - add the func to read the yaml file and get the info and fill the

	// - form with the info

	// * 2. make a form for Instruments and Periods and Indicators
	// make a form for Instruments and Periods and Indicators
	EntryInstruments := widget.NewEntry()
	EntryBDate := widget.NewEntry()
	EntryEDate := widget.NewEntry()
	EntryIndicators := widget.NewEntry()
	EntrySubIndicators := widget.NewEntry()
	EntryCalIndicators := widget.NewEntry()

	formBTinfo := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Instruments", Widget: EntryInstruments},
			{Text: "BeginDate", Widget: EntryBDate},
			{Text: "EndDate", Widget: EntryEDate},
			{Text: "Indicators", Widget: EntryIndicators},
			{Text: "SubIndicators", Widget: EntrySubIndicators},
			{Text: "CalIndicators", Widget: EntryCalIndicators},
		},
	}
	//
	// make a button with a callback change la with BTdirEntry.Text
	btnReadConf := widget.NewButton("ReadConfig", func() {

		// read the yaml file from dir and output the info
		instr, bd, ed, inds, sub, cal := backtest.BtnReadConf_Clicked(EntryBTdir.Text)
		// change the formBTinfo with the info
		EntryInstruments.SetText(instr)
		EntryBDate.SetText(bd)
		EntryEDate.SetText(ed)
		EntryIndicators.SetText(inds)
		EntrySubIndicators.SetText(sub)
		EntryCalIndicators.SetText(cal)
	})
	// split the form and button
	hsBTdir := container.NewHSplit(formBTdir, btnReadConf)
	hsBTdir.SetOffset(0.75)
	// make a gray separator
	sepBTinfo := canvas.NewRectangle(color.Gray{Y: 128})
	// add hs and sep in a vbox
	contBTdir := container.NewVBox(hsBTdir, sepBTinfo)
	// add a button for DataDownload
	btnDownload := widget.NewButton("DataDownload", func() {
	})
	// - add a func to download data
	// - void for now!!

	// should read the yaml file and get entries
	// / this string slice part should be derived from the EntryCalIndicators.text
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
	BTEntryCont := container.NewVBox(formBTinfo, btnDownload, ii, IIform, hs1, sep1)

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
	content := container.NewVBox(contBTdir, BTEntryCont, STGCont, hs2, grid, laRes)

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
