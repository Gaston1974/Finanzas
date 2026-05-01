package formularios

import (
	"fmt"
	"hello/src/pkg/scripts"
	"hello/src/pkg/styles"
	"strconv"

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
	text1 := widget.NewLabel("SE CLASIFICA CRIPTOMONEDAS SEGUN PRECIOS \nENTRE LOS SIGUIENTES TIPOS: \n\n " +
		"🔹BTCUSDT (Bitcoin) \n 🔹ETHUSDT (Ethereum) \n 🔹SOLUSDT \n 🔹ADAUSDT \n\n DE ACUERDO A BINANCE.")

	form := widget.NewForm()

	item1 := widget.NewFormItem("", titulo)
	item2 := widget.NewFormItem("", widget.NewLabel(" "))

	item6 := widget.NewFormItem("", lbl2)
	item7 := widget.NewFormItem("", text1)

	form.Items = []*widget.FormItem{item1, item2, item2, item7, item6}

	form.OnSubmit = func() {
		fmt.Printf("\nconsulta enviada\n")

		obj, res := scripts.Info2(form, lbl2)

		if res == 1 {
			showTable(a, obj)

		} else {
			lbl2.SetText("error..")
		}

	}

	form.SubmitText = "Generar"

	cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), form, layout.NewSpacer()}
	cont.Refresh()

}

func showTable(a fyne.App, ob []scripts.Classify) {

	w := a.NewWindow("RESULTADOS")

	fil := len(ob)
	atributes := []string{"nombre", "precio", "categoria"}
	col := len(atributes)

	data := make([][]string, fil+1)

	for j := 0; j < fil+1; j++ {
		data[j] = make([]string, col)
	}

	for i := range 3 {

		data[0][i] = atributes[i]
	}

	for k, v := range ob {
		k++
		data[k][0] = v.Name.Symbol
		data[k][1] = strconv.FormatFloat(v.Name.Price, 'f', 2, 64)
		data[k][2] = v.Category
	}

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			// return widget.NewLabel("wide content")
			return canvas.NewText("wide content", color.White)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*canvas.Text).Text = data[i.Row][i.Col]
			if i.Row == 0 {

				o.(*canvas.Text).Color = color.NRGBA{R: 0, G: 180, B: 0, A: 255}
			}
			// o.(*widget.Label).
			// o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	w.SetContent(list)
	w.Resize(fyne.NewSize(250, 150))
	w.SetContent(container.NewVScroll(list))
	w.Show()

}
