package formularios

import (
	"fmt"
	"hello/src/pkg/styles"
	"os"
	"os/exec"
	"path/filepath"

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
func ShowCryptosWW(im *canvas.Image, cont *fyne.Container, id string, win fyne.Window, a fyne.App) {

	cont.Layout = styles.NewBarberiaLayout(120)

	titulo := canvas.NewText("GENERACION DE INFORMES", color.White)
	titulo.TextStyle = fyne.TextStyle{Bold: true, Italic: false, Underline: true, Symbol: true}
	titulo.Alignment = fyne.TextAlignCenter

	// lbl := widget.NewLabel("seleccione ruta de descarga..")
	lbl2 := widget.NewLabel("")
	text1 := widget.NewLabel("SE CLASIFICA ACTIVOS FINANCIEROS \nENTRE LOS SIGUIENTES TIPOS: \n\n " +
		"🔹ACCIONES \n 🔹CRYPTOS \n 🔹BONOS \n 🔹DOLARES \n\n .")

	form := widget.NewForm()

	item1 := widget.NewFormItem("", titulo)
	item2 := widget.NewFormItem("", widget.NewLabel(" "))

	item6 := widget.NewFormItem("", lbl2)
	item7 := widget.NewFormItem("", text1)

	form.Items = []*widget.FormItem{item1, item2, item2, item7, item6}

	form.OnSubmit = func() {
		fmt.Printf("\nconsulta enviada\n")

		// Path to the shell script
		scriptPath := "./src/pkg/apiDatas/scriptDatos.py"

		absPath, err := filepath.Abs(scriptPath)
		if err != nil {
			fmt.Println("Error getting absolute path:", err)
			return
		}

		// Create a new command to run the shell script
		cmd := exec.Command("python3", absPath)

		// Set the command's output to be the standard output (terminal)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println("Error executing script:", err.Error())
			return
		} else {

			lbl2.SetText("GRAFICO GENERADO")
		}

	}

	form.SubmitText = "Generar"

	cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), form, layout.NewSpacer()}
	cont.Refresh()

}
