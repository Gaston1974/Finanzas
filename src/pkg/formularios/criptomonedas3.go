package formularios

import (
	"fmt"
	"hello/src/pkg/scripts"
	"hello/src/pkg/styles"

	// "hello/src/pkg/dao"
	// funciones "hello/src/pkg/funciones"
	// handlers "hello/src/pkg/handlers"
	"image/color"

	//"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Show loads a modificacion example window for the specified app context
func ShowCryptosMM(im *canvas.Image, cont *fyne.Container, id string, win fyne.Window, a fyne.App) {

	cont.Layout = styles.NewBarberiaLayout(120)

	titulo := canvas.NewText("GENERACION DE INFORMES", color.White)
	titulo.TextStyle = fyne.TextStyle{Bold: true, Italic: false, Underline: true, Symbol: true}
	titulo.Alignment = fyne.TextAlignCenter

	// lbl := widget.NewLabel("seleccione ruta de descarga..")
	lbl2 := widget.NewLabel("")
	text1 := widget.NewLabel("SE CLASIFICA CRIPTOMONEDAS SEGUN PRECIOS \nENTRE LOS SIGUIENTES TIPOS: \n\n " +
		"🔹BTCUSDT(Bitcoin) \n 🔹ETHUSDT(Ethereum) \n 🔹SOLUSDT \n 🔹ADAUSDT \n\n DE ACUERDO A BINANCE.")

	form := widget.NewForm()

	item1 := widget.NewFormItem("", titulo)
	item2 := widget.NewFormItem("", widget.NewLabel(" "))

	item6 := widget.NewFormItem("", lbl2)
	item7 := widget.NewFormItem("", text1)

	form.Items = []*widget.FormItem{item1, item2, item2, item7, item6}

	form.OnSubmit = func() {
		fmt.Printf("\nconsulta enviada\n")

		// obj, res := scripts.Info3()
		scripts.Info3()
		// if res == 1 {
		// 	showTable(a, obj)

		// } else {
		// 	lbl2.SetText("error..")
		// }

	}

	form.SubmitText = "Generar"

	cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), form, layout.NewSpacer()}
	cont.Refresh()

}
