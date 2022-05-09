package main

import (
	"fyne.io/fyne/v2/layout"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type myTheme struct{}

var _ fyne.Theme = (*myTheme)(nil)

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(name, variant)
}

func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Bold {
		return resourceMsyhbdTtc
	}
	if style.Italic {
		return theme.DefaultTheme().Font(style)
	}
	if style.Monospace {
		return theme.DefaultTheme().Font(style)
	}
	return resourceMsyhTtc
}

func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

var ip = widget.NewLabel("")
var position = widget.NewLabel("")
var isp = widget.NewLabel("")
var ip1 = widget.NewLabel("")
var ip2 = widget.NewLabel("")
var ip3 = widget.NewLabel("")
var myApp = app.New()

func yuming(urlhttp string) string {

	return "err"

}

func ipdizhi(urlhttp string) string {

	return "err"
}
func icp(urlhttp string) string {

	return "err"
}

func icpadd(urlhttp string) string {

	return "err"
}

func gonsi(urlhttp string) string {

	return "err"

}

func zhToUnicode(raw []byte) ([]byte, error) {

	return "err"
}

func yumings(urlhttp string) string {

	return "err"
	

}

func info(res1 string, res2 string, res3 string, res4 string, res5 string, res6 string) fyne.CanvasObject {

	return "err"

}

func query() fyne.CanvasObject {
	return "err"
}

func main() {
	myApp.Settings().SetTheme(&myTheme{})
	myWindow := myApp.NewWindow("IPconfig-Summ1e233")
	myWindow.Resize(fyne.NewSize(800, 600))

	//myWindow.SetContent(container.New(layout.NewHBoxLayout(), b1, form))
	//myWindow.SetContent(container.New(layout.NewHBoxLayout(), query(), info()))
	myWindow.SetContent(fyne.NewContainerWithLayout(layout.NewGridLayoutWithColumns(2), info("", "", "", "", "", ""), query()))
	//myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
