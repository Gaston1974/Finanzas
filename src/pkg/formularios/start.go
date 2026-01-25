package formularios

import (

	// "hello/src/pkg/dao"
	// funciones "hello/src/pkg/funciones"
	// handlers "hello/src/pkg/handlers"

	//"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
)

// Show loads a modificacion example window for the specified app context
func ShowStart(im *canvas.Image, cont *fyne.Container, id string, win fyne.Window, a fyne.App) {

	// titulo := canvas.NewText(" ", color.White)
	// titulo.TextStyle = fyne.TextStyle{Bold: true, Italic: true, Underline: true, Symbol: true}
	// titulo.Alignment = fyne.TextAlignCenter

	//logo := canvas.NewImageFromResource(resourceCrypto4Png)

	// cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), im, layout.NewSpacer()}

	// top := canvas.NewLine(color.White)
	// bottom := canvas.NewLine(color.White)
	// left := canvas.NewLine(color.White)
	// right := canvas.NewLine(color.White)
	//middle := canvas.NewLine(color.White)

	//c := container.NewBorder(top, bottom, left, right, im)

	cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), im, layout.NewSpacer()}
	im.FillMode = canvas.ImageFillContain
	cont.Refresh()

}
