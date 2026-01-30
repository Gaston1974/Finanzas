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
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Show loads a modificacion example window for the specified app context
func ShowCryptosCmp(im *canvas.Image, cont *fyne.Container, id string, win fyne.Window, a fyne.App) {

	cont.Layout = styles.NewBarberiaLayout(120)

	titulo := canvas.NewText("GENERACION DE INFORMES", color.White)
	titulo.TextStyle = fyne.TextStyle{Bold: true, Italic: false, Underline: true, Symbol: true}
	titulo.Alignment = fyne.TextAlignCenter

	// lbl := widget.NewLabel("seleccione ruta de descarga..")
	lbl2 := widget.NewLabel("")
	text1 := widget.NewLabel("SE CLASIFICA CRIPTOMONEDAS SEGUN PRECIOS \nENTRE LAS SIGUIENTES: \n\n " +
		"BTCUSDT \n ETHUSDT \n SOLUSDT \n ADAUSDT \n\n DE ACUERDO A BINANCE.")

	form := widget.NewForm()

	item1 := widget.NewFormItem("", titulo)
	item2 := widget.NewFormItem("", widget.NewLabel(" "))

	item6 := widget.NewFormItem("", lbl2)
	item7 := widget.NewFormItem("", text1)

	form.Items = []*widget.FormItem{item1, item2, item2, item7, item6}

	form.OnSubmit = func() {
		fmt.Printf("\nconsulta enviada\n")

		lbl2.SetText("GENERANDO ARCHIVO .....")
		form.Refresh()

		obj := scripts.Info2()

		showTable

		// if res == 0 {

		// 	lbl2.SetText("ARCHIVO GENERADO")

		// } else {

		// 	lbl2.SetText("ERROR: ARCHIVO NO GENERADO")

		// }

	}

	form.SubmitText = "Generar"

	cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), form, layout.NewSpacer()}
	cont.Refresh()

}

func showTable(w fyne.Window, ob []scripts.Classify) {

	fil := len(ob)
	atributes := ob[0].Atr()
	col := len(atributes)

	data := make([][]string, fil)

	for i := 0; i < fil; i++ {
		data[i] = make([]string, col)
	}

	for i, v := range ob {

		for f, at := range atributes {

			data[i][f] = v.Get(at)
			// fmt.Printf("\n objeto: %s", v.Caratula)
			// fmt.Printf("\n atributo: %s", at)
			// fmt.Printf("\n i: %d f: %d", i, f)
			// fmt.Printf("\n valor: %s", v.Get(at))
			// apiDatas.Wait(3)

			//fmt.Printf("/n valor: %s", v.Get(at))

		}

	}

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	w.SetContent(list)
	w.Resize(fyne.NewSize(1300, 600))
	w.SetContent(container.NewVScroll(list))
	w.ShowAndRun()

}
