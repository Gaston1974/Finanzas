package formularios

import (
	"fmt"
	"hello/src/pkg/apiDatas"
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
func ShowCryptosInf(im *canvas.Image, cont *fyne.Container, id string, win fyne.Window, a fyne.App) {

	cont.Layout = styles.NewBarberiaLayout(120)

	titulo := canvas.NewText("GENERACION DE INFORMES", color.White)
	titulo.TextStyle = fyne.TextStyle{Bold: true, Italic: false, Underline: true, Symbol: true}
	titulo.Alignment = fyne.TextAlignCenter

	lbl := widget.NewLabel("seleccione ruta de descarga..")
	lbl2 := widget.NewLabel("")
	in1 := widget.NewEntry()
	in1.PlaceHolder = "0"
	in2 := widget.NewEntry()
	in2.PlaceHolder = "0"
	in3 := widget.NewEntry()
	in3.PlaceHolder = "0"
	button := widget.NewButton("DESCARGAS", func() { apiDatas.SetPath(a, lbl) })
	button.Importance = widget.MediumImportance

	form := widget.NewForm()

	item1 := widget.NewFormItem("", titulo)
	item2 := widget.NewFormItem("", widget.NewLabel(" "))
	//item3 := widget.NewFormItem("", widget.NewLabel("Exportar archivo de datos. \n"))

	item4 := widget.NewFormItem("", button)
	item5 := widget.NewFormItem("", lbl)
	item6 := widget.NewFormItem("", lbl2)

	item7 := widget.NewFormItem("ingrese cantidad de dias: ", in1)
	item8 := widget.NewFormItem("ingrese cantidad de meses: ", in2)
	item9 := widget.NewFormItem("ingrese cantidad de anios: ", in3)

	form.Items = []*widget.FormItem{item1, item2, item2, item4, item5, item2, item7, item8, item9, item2, item6, item2}

	form.OnSubmit = func() {
		fmt.Printf("\nconsulta enviada\n")

		lbl2.SetText("GENERANDO ARCHIVO .....")
		form.Refresh()

		res, _ := scripts.Info1(lbl.Text, in1.Text, in2.Text, in3.Text)

		if res == 0 {
			lbl2.SetText("ARCHIVO GENERADO")
			// w := a.NewWindow("informes")
			// w.Resize(fyne.NewSize(350, 310))
			// w.CenterOnScreen()
			// dialog.NewCustomConfirm("INFORMES", "ARCHIVO GENERADO", "", nil, nil, w)
			// w.Show()

		} else {

			lbl2.SetText("ERROR: ARCHIVO NO GENERADO")
			//apiDatas.Log(msg, "C:\\log.txt")

		}

	}

	form.SubmitText = "Exportar"

	cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), form, layout.NewSpacer()}
	cont.Refresh()

}
