package oldies

/*
import (
	"fmt"
	"hello/src/pkg/dao"
	funciones "hello/src/pkg/funciones"
	"hello/src/pkg/handlers"
	"strconv"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ShowConsulta(flag *int, cont *fyne.Container, id string, win fyne.Window, a fyne.App) {

	titulo := canvas.NewText("Consultas", color.White)
	titulo.TextStyle = fyne.TextStyle{Bold: true, Italic: true, Underline: true, Symbol: true}
	titulo.Alignment = fyne.TextAlignCenter

	var table string

	combo := widget.NewSelect([]string{"usuarios", "asignaciones", "emprendedores", "totales", "notificaciones"}, func(value string) {
		table = value
	})
	combo.PlaceHolder = " "

	res, msg := funciones.ActualizaVencimientos()
	if res == 0 {
		cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), canvas.NewText(msg, color.White), layout.NewSpacer()}
		funciones.Wait(2)
		return
	}

	var user funciones.Identificable = dao.Usuario{}
	var cli funciones.Identificable = dao.Cliente{}
	var sto funciones.Identificable = dao.Asignacion{}
	var vent funciones.Identificable = dao.Venta{}
	var tot funciones.Identificable = dao.Total{}
	var not funciones.Identificable = dao.Notificacion{}

	tablas := map[string]funciones.Identificable{
		"usuarios":       user,
		"emprendedores":  cli,
		"asignaciones":   sto,
		"ventas":         vent,
		"totales":        tot,
		"notificaciones": not,
	}

	checkPDF := widget.NewCheck("(solo para asignaciones)", func(value bool) {
		//	fmt.Println("Check set to: ", value)
	})

	subTitulo2 := canvas.NewText("", color.White)

	form := widget.NewForm()

	item1 := widget.NewFormItem("", titulo)
	item2 := widget.NewFormItem("", widget.NewLabel(" "))
	item3 := widget.NewFormItem("Elija una opción", combo)
	item4 := widget.NewFormItem("", subTitulo2)
	item7 := widget.NewFormItem("", subTitulo2)

	item10 := widget.NewFormItem("Generar PDF", checkPDF)

	form.Items = []*widget.FormItem{item1, item2, item3, item10, item4, item7}

	form.OnSubmit = func() {
		fmt.Printf("\n datos de consulta enviados \n")
		funciones.Log("\n datos de consulta enviados \n", nil, "./barberia_logs.txt")

		f := tablas[table]

		//params := []string{in1.Text, in2.Text}

		res, msg, vec := handlers.HandlerCombos(f, table, "", nil, checkPDF.Checked)

		if res == 0 || len(vec) == 0 { // error o no hay datos en la tabla
			cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), canvas.NewText(msg, color.White), layout.NewSpacer()}
			funciones.Log(msg, nil, "./barberia_logs.txt")
			funciones.Wait(2)

			cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), form, layout.NewSpacer()}

		} else {

			var fieldWithStyle *canvas.Text
			atributes := vec[0].Atr()
			fieldsNum := len(atributes)
			grid := container.New(layout.NewGridLayoutWithColumns(fieldsNum))

			for i := 0; i < len(vec); i++ { // nuevo registro en la grilla. Recorro el vector de registros consultados en la base.

				for j := 0; j < fieldsNum; j++ { // completo el registro en la grilla

					if i == 0 && j == 0 { // titulo de la grilla

						for _, v := range atributes {
							titulo := canvas.NewText(v, color.NRGBA{R: 0, G: 180, B: 0, A: 255})
							titulo.TextStyle = fyne.TextStyle{Bold: true, Italic: true, Underline: true, Symbol: true}
							titulo.Alignment = fyne.TextAlignCenter
							grid.Add(titulo)

						}

					}

					// armado de grilla

					field := vec[i].Get(atributes[j])

					switch field {

					case "vacio":
						fieldWithStyle = canvas.NewText(field, color.NRGBA{R: 180, G: 0, B: 0, A: 255})

					default:
						fieldWithStyle = canvas.NewText(field, color.White)
					}

					if atributes[j] == "Precio" {

						fieldWithStyle = canvas.NewText(field, color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff})

					}

					if table == "stocks" {
						if atributes[j] == "Cantidad" {

							cant, _ := strconv.Atoi(vec[i].Get("Cantidad"))
							cantMin, _ := strconv.Atoi(vec[i].Get("Cantidad Minima"))

							if cant <= cantMin {

								fieldWithStyle = canvas.NewText(field, color.NRGBA{R: 180, G: 0, B: 0, A: 255})
							}

						} else if atributes[j] == "Descripcion" {
							fieldWithStyle.TextSize = 12
						}

					} /*else if table == "clientes" {
						if atributes[j] == "Puntos" {

							p, _ := strconv.Atoi(vec[i].Get("Puntos"))

							if p >= min {

								fieldWithStyle = canvas.NewText(field, color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff})
								fieldWithStyle.TextSize = 11
							}

						}

					} */

/*
					fieldWithStyle.Alignment = fyne.TextAlignCenter
					fieldWithStyle.TextStyle.Bold = false
					grid.Add(fieldWithStyle)

				}

			}

			if checkPDF.Checked && table == "asignaciones" {
				i, msg := funciones.GenerarPDF(vec)
				if i != 1 {
					funciones.Log(msg, nil, "./barberia_logs.txt")
					fmt.Printf("\n %s \n", msg)
				}
			}

			r := a.NewWindow(id)
			r.Resize(fyne.NewSize(830, 410))
			r.CenterOnScreen()
			r.SetContent(container.NewScroll(grid))
			r.Show()
		}

	}

	form.SubmitText = "Consultar"

	cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), form, layout.NewSpacer()}

}



*/
