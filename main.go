package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"os/exec"
)

func getJokeViaCurl() (joke string, err error) {
	cmd, _ := exec.Command("curl", "-H \"Accept: text/plain\"", "https://icanhazdadjoke.com/").Output()
	return string(cmd), nil
}

func main() {
	a := app.New()
	w := a.NewWindow("Another joke for your kids")
	w.Resize(fyne.NewSize(640, 200))

	jokeText := widget.NewLabel("")
	jokeText.Wrapping = fyne.TextWrapWord

	button := widget.NewButton("Ask dad for a joke!", func() {
		joke, err := getJokeViaCurl()
		if err != nil {
			dialog.ShowError(err, w)
		} else {
			jokeText.SetText(joke)
		}
	})

	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), hBox, widget.NewSeparator(), jokeText)

	a.Settings().SetTheme(theme.DarkTheme())

	w.SetContent(vBox)
	w.ShowAndRun()
}
