package formularios

import (
	"database/sql"
	"fmt"
	"image/color"
	"log"
	"strings"

	"hello/src/pkg/apiDatas"
	"hello/src/pkg/dao"
	"hello/src/pkg/styles"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/tealeg/xlsx"
)

func ShowCausas(im *canvas.Image, cont *fyne.Container, id string, win fyne.Window, a fyne.App) {

	cont.Layout = styles.NewBarberiaLayout(120)

	titulo := canvas.NewText("GENERACION DE INFORMES", color.White)
	titulo.TextStyle = fyne.TextStyle{Bold: true, Italic: false, Underline: true, Symbol: true}
	titulo.Alignment = fyne.TextAlignCenter

	lbl := widget.NewLabel("seleccione la carpeta de descarga ..")
	lbl2 := widget.NewLabel("")
	lbl3 := widget.NewLabel("")

	in1 := widget.NewEntry()
	in1.PlaceHolder = "0"

	button := widget.NewButton("DESCARGAS", func() { apiDatas.SetPath(a, lbl) })
	button.Importance = widget.MediumImportance
	form := widget.NewForm()

	item1 := widget.NewFormItem("", titulo)
	item2 := widget.NewFormItem("", widget.NewLabel(" "))
	item3 := widget.NewFormItem("", lbl)
	item4 := widget.NewFormItem("", lbl2)
	item5 := widget.NewFormItem("INGRESE N DE CAUSA: ", in1)
	item7 := widget.NewFormItem("", lbl3)

	item6 := widget.NewFormItem("", button)

	form.Items = []*widget.FormItem{item1, item2, item3, item6, item2, item5, item7, item4, item2, item2, item2}

	form.OnSubmit = func() {
		fmt.Printf("\nconsulta enviada\n")

		var sqlStatement string
		var err error
		var rows *sql.Rows
		var object []dao.Causa
		var cau dao.Causa

		lbl2.SetText("GENERANDO ARCHIVO .... ")
		form.Refresh()

		db := apiDatas.Acceso{}

		_, val, msg := db.SetCliente()

		if val != 1 {
			fmt.Println("\nctx: " + msg)
			lbl2.Text = msg
			form.Refresh()
		} else {
			fmt.Println("\nctx: OK")
		}

		// Test the connection to the database
		if err := db.Cliente.Ping(); err != nil {
			fmt.Printf("\ntest connection: %s", err.Error())
			lbl2.Text = msg
			form.Refresh()
		} else {
			log.Println("\nSuccessfully Connected")
		}

		defer db.Cliente.Close()

		if in1.Text != "" {

			sqlStatement = " SELECT numero_causa , COALESCE(caratula, 'VACIO') , j.nombre , f.nombre , COALESCE(a_cargo_del_magistrado, 'VACIO') , " +
				" preventor , preventor_auxiliar,  COALESCE(p.nombre, 'vacio') , COALESCE(l.nombre, 'vacio') , COALESCE(domicilio, '')  , " +
				" COALESCE(nro_mto, 'VACIO') , COALESCE(nro_sgo, 'VACIO')  , " +
				" COALESCE(tipo_delito, 'VACIO'), COALESCE(nombre_fantasia, 'VACIO'), COALESCE(fecha_llegada, 'VACIO') , COALESCE(providencia, 'VACIO'), " +
				" COALESCE(estado, 'VACIO'), '' as ipaddress , " +
				" COALESCE(d.nombre_archivo, 'VACIO'),  COALESCE(d.ruta_archivo, 'VACIO'), COALESCE(d.tipo_documento, 'VACIO'), " +
				" COALESCE(d.tamano, 'VACIO'), " +
				" COALESCE(d.subido_por, 'VACIO'), COALESCE(n.contenido, 'VACIO') " +
				" FROM causas c " +
				" LEFT JOIN documentos_causa d ON c.id = d.causa_id " +
				" LEFT JOIN notas_causa n ON c.id = n.causa_id " +
				" INNER JOIN fiscalias f ON fiscalia_id = f.id " +
				" INNER JOIN juzgados j ON juzgado_id = j.id  " +
				" INNER JOIN provincias p ON c.provincia_id = p.id " +
				" INNER JOIN localidades l ON c.localidad_id = l.id " +
				" WHERE numero_causa = ? " +
				" ORDER BY numero_causa;"

			rows, err = db.Cliente.Query(sqlStatement, in1.Text)

		} else {

			sqlStatement = " SELECT numero_causa , COALESCE(caratula, 'VACIO') , j.nombre , f.nombre , COALESCE(a_cargo_del_magistrado, 'VACIO') , " +
				" preventor , preventor_auxiliar,  provincia_id , localidad_id , COALESCE(domicilio, '')  , " +
				" COALESCE(nro_mto, 'VACIO') , COALESCE(nro_sgo, 'VACIO')  , " +
				" COALESCE(tipo_delito, 'VACIO'), COALESCE(nombre_fantasia, 'VACIO'), COALESCE(fecha_llegada, 'VACIO') , COALESCE(providencia, 'VACIO'), " +
				" COALESCE(estado, 'VACIO'), '' as ipaddress , " +
				" COALESCE(d.nombre_archivo, 'VACIO'),  COALESCE(d.ruta_archivo, 'VACIO'), COALESCE(d.tipo_documento, 'VACIO'), " +
				" COALESCE(d.tamano, 'VACIO'), " +
				" COALESCE(d.subido_por, 'VACIO'), COALESCE(n.contenido, 'VACIO') " +
				" FROM causas c " +
				" LEFT JOIN documentos_causa d ON c.id = d.causa_id " +
				" LEFT JOIN notas_causa n ON c.id = n.causa_id " +
				" INNER JOIN fiscalias f ON fiscalia_id = f.id " +
				" INNER JOIN juzgados j ON juzgado_id = j.id  " +
				" ORDER BY providencia;"

			rows, err = db.Cliente.Query(sqlStatement)

		}

		if err == nil {
			for rows.Next() {

				cau.Load(rows)
				object = append(object, cau)
			}

			fileRes := GenXLSX(object, lbl.Text)
			if fileRes == "1" {
				apiDatas.Wait(2)
				lbl2.Text = ""
				lbl3.Text = "ARCHIVO GENERADO"
				form.Refresh()
			}
			// showTable(win, object)

		} else {

			lbl2.Text = "ERROR EN LA CONSULTA A LA BASE DE DATOS"

		}

	}

	form.SubmitText = "Consultar"

	cont.Objects = []fyne.CanvasObject{layout.NewSpacer(), form, layout.NewSpacer()}
	cont.Refresh()

}

func GenXLSX(ob []dao.Causa, lbl string) string {

	// Create a new Excel file
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf("Failed to add sheet: %s\n", err)
		return "Failed to add sheet"
	}

	style := apiDatas.AddStyle(2)

	// Add row and cells

	roww := sheet.AddRow()

	causa := roww.AddCell()
	causa.SetStyle(style)
	causa.Value = "NRO CAUSA"
	caratula := roww.AddCell()
	caratula.SetStyle(style)
	caratula.Value = "CARATULA"
	juzgado := roww.AddCell()
	juzgado.SetStyle(style)
	juzgado.Value = "JUZGADO"
	fiscalia := roww.AddCell()
	fiscalia.SetStyle(style)
	fiscalia.Value = "FISCALIA"
	magistrado := roww.AddCell()
	magistrado.SetStyle(style)
	magistrado.Value = "MAGISTRADO"
	preventor := roww.AddCell()
	preventor.SetStyle(style)
	preventor.Value = "PREVENTOR"
	prevAux := roww.AddCell()
	prevAux.SetStyle(style)
	prevAux.Value = "PREVENTOR AUXILIAR"
	provincia := roww.AddCell()
	provincia.SetStyle(style)
	provincia.Value = "PROVINCIA ID"
	localidad := roww.AddCell()
	localidad.SetStyle(style)
	localidad.Value = "LOCALIDAD ID"
	domicilio := roww.AddCell()
	domicilio.SetStyle(style)
	domicilio.Value = "DOMICILIO"
	nroSgo := roww.AddCell()
	nroSgo.SetStyle(style)
	nroSgo.Value = "NRO SGO"
	nroMto := roww.AddCell()
	nroMto.SetStyle(style)
	nroMto.Value = "NRO MTO"
	tipoDelito := roww.AddCell()
	tipoDelito.SetStyle(style)
	tipoDelito.Value = "TIPO DELITO"
	nombreFan := roww.AddCell()
	nombreFan.SetStyle(style)
	nombreFan.Value = "NOMBRE FANTASIA"
	fecha := roww.AddCell()
	fecha.SetStyle(style)
	fecha.Value = "FECHA"
	providencia := roww.AddCell()
	providencia.SetStyle(style)
	providencia.Value = "PROVIDENCIA"
	estado := roww.AddCell()
	estado.SetStyle(style)
	estado.Value = "ESTADO"
	ip := roww.AddCell()
	ip.SetStyle(style)
	ip.Value = "IP"
	archivo := roww.AddCell()
	archivo.SetStyle(style)
	archivo.Value = "ARCHIVO"
	ruta := roww.AddCell()
	ruta.SetStyle(style)
	ruta.Value = "RUTA"
	formato := roww.AddCell()
	formato.SetStyle(style)
	formato.Value = "FORMATO"
	tam := roww.AddCell()
	tam.SetStyle(style)
	tam.Value = "TAMANO"
	user := roww.AddCell()
	user.SetStyle(style)
	user.Value = "USUARIO"
	nota := roww.AddCell()
	nota.SetStyle(style)
	nota.Value = "NOTA"

	maxColumn0Lenght, maxColumn1Lenght, maxColumn2Lenght,
		maxColumn3Lenght, maxColumn4Lenght, maxColumn5Lenght,
		maxColumn6Lenght, maxColumn7Lenght, maxColumn8Lenght,
		maxColumn9Lenght, maxColumn10Lenght, maxColumn11Lenght,
		maxColumn12Lenght, maxColumn13Lenght, maxColumn14Lenght,
		maxColumn15Lenght, maxColumn16Lenght, maxColumn17Lenght,
		maxColumn18Lenght, maxColumn19Lenght, maxColumn20Lenght,
		maxColumn21Lenght, maxColumn22Lenght, maxColumn23Lenght := 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0

	for _, v := range ob {

		style = apiDatas.AddStyle(1)
		style3 := apiDatas.AddStyle(3)
		roww = sheet.AddRow()

		causa := roww.AddCell()
		causa.SetStyle(style3)
		causa.Value = v.Nro_causa
		if float64(len(causa.Value))*0.7 > maxColumn0Lenght {
			maxColumn0Lenght = float64(len(causa.Value)) * 0.7
			sheet.SetColWidth(0, 0, maxColumn0Lenght)
		}

		caratula := roww.AddCell()
		caratula.SetStyle(style)
		caratula.Value = v.Caratula
		if float64(len(caratula.Value))*0.6 > maxColumn1Lenght {
			maxColumn1Lenght = float64(len(caratula.Value)) * 0.6
			sheet.SetColWidth(1, 1, maxColumn1Lenght)
		}

		juzgado := roww.AddCell()
		juzgado.SetStyle(style)
		juzgado.Value = v.Juzgado
		if float64(len(juzgado.Value))*0.6 > maxColumn2Lenght {
			maxColumn2Lenght = float64(len(juzgado.Value)) * 0.6
			sheet.SetColWidth(2, 2, maxColumn2Lenght)
		}

		fiscalia := roww.AddCell()
		fiscalia.SetStyle(style)
		fiscalia.Value = v.Fiscalia
		if float64(len(strings.Trim(fiscalia.Value, " ")))*0.6 > maxColumn3Lenght {
			maxColumn3Lenght = float64(len(strings.Trim(fiscalia.Value, " "))) * 0.6
			sheet.SetColWidth(3, 3, maxColumn3Lenght)
		}

		magistrado := roww.AddCell()
		magistrado.SetStyle(style)
		magistrado.Value = v.Magistrado
		if float64(len(magistrado.Value)) > maxColumn4Lenght {
			maxColumn4Lenght = float64(len(magistrado.Value))
			sheet.SetColWidth(4, 4, maxColumn4Lenght)
		}

		preventor := roww.AddCell()
		preventor.SetStyle(style)
		preventor.Value = v.Preventor
		if float64(len(preventor.Value))*0.6 > maxColumn5Lenght {
			maxColumn5Lenght = float64(len(preventor.Value)) * 0.6
			sheet.SetColWidth(5, 5, maxColumn5Lenght)
		}

		prevAux := roww.AddCell()
		prevAux.SetStyle(style)
		prevAux.Value = v.Preventor_auxiliar
		if float64(len(prevAux.Value)) > maxColumn6Lenght {
			maxColumn6Lenght = float64(len(prevAux.Value))
			sheet.SetColWidth(6, 6, maxColumn6Lenght)
		}

		provincia := roww.AddCell()
		provincia.SetStyle(style)
		provincia.Value = v.Provincia_id
		if float64(len(provincia.Value))*9 > maxColumn7Lenght {
			maxColumn7Lenght = float64(len(provincia.Value)) * 9
			sheet.SetColWidth(7, 7, maxColumn7Lenght)
		}

		localidad := roww.AddCell()
		localidad.SetStyle(style)
		localidad.Value = v.Localidad_id
		if float64(len(localidad.Value))*9 > maxColumn8Lenght {
			maxColumn8Lenght = float64(len(localidad.Value)) * 9
			sheet.SetColWidth(8, 8, maxColumn8Lenght)
		}

		domicilio := roww.AddCell()
		domicilio.SetStyle(style)
		domicilio.Value = v.Domicilio
		if float64(len(domicilio.Value))*0.6 > maxColumn23Lenght {
			maxColumn23Lenght = float64(len(domicilio.Value)) * 0.6
			sheet.SetColWidth(9, 9, maxColumn23Lenght)
		}

		nroSgo := roww.AddCell()
		nroSgo.SetStyle(style)
		nroSgo.Value = v.Nro_sgo
		if float64(len(nroSgo.Value))*0.6 > maxColumn9Lenght {
			maxColumn9Lenght = float64(len(nroSgo.Value)) * 0.6
			sheet.SetColWidth(10, 10, maxColumn9Lenght)
		}

		nroMto := roww.AddCell()
		nroMto.SetStyle(style)
		nroMto.Value = v.Nro_mto
		if float64(len(nroMto.Value)) > maxColumn10Lenght {
			maxColumn10Lenght = float64(len(nroMto.Value))
			sheet.SetColWidth(11, 11, maxColumn10Lenght)
		}

		tipoDelito := roww.AddCell()
		tipoDelito.SetStyle(style)
		tipoDelito.Value = v.Tipo_delito
		if float64(len(tipoDelito.Value)) > maxColumn11Lenght {
			maxColumn11Lenght = float64(len(tipoDelito.Value))
			sheet.SetColWidth(12, 12, maxColumn11Lenght)
		}

		nombreFan := roww.AddCell()
		nombreFan.SetStyle(style)
		nombreFan.Value = v.Nombre_fantasia
		if float64(len(nombreFan.Value))*1.1 > maxColumn12Lenght {
			maxColumn12Lenght = float64(len(nombreFan.Value)) * 1.1
			sheet.SetColWidth(13, 13, maxColumn12Lenght)
		}

		fecha := roww.AddCell()
		fecha.SetStyle(style)
		fecha.Value = v.Fecha
		if float64(len(fecha.Value)) > maxColumn13Lenght {
			maxColumn13Lenght = float64(len(fecha.Value))
			sheet.SetColWidth(14, 14, maxColumn13Lenght)
		}

		providencia := roww.AddCell()
		providencia.SetStyle(style)
		providencia.Value = v.Providencia
		if float64(len(providencia.Value))*1.1 > maxColumn14Lenght {
			maxColumn14Lenght = float64(len(providencia.Value)) * 1.1
			sheet.SetColWidth(15, 15, maxColumn14Lenght)
		}

		estado := roww.AddCell()
		estado.SetStyle(style)
		estado.Value = v.Estado
		if float64(len(estado.Value)) > maxColumn15Lenght {
			maxColumn15Lenght = float64(len(estado.Value))
			sheet.SetColWidth(16, 16, maxColumn15Lenght)
		}

		ip := roww.AddCell()
		ip.SetStyle(style)
		ip.Value = v.IpAdress
		if float64(len(ip.Value))*1.1 > maxColumn16Lenght {
			maxColumn16Lenght = float64(len(ip.Value)) * 1.1
			sheet.SetColWidth(17, 17, maxColumn16Lenght)
		}

		archivo := roww.AddCell()
		archivo.SetStyle(style)
		archivo.Value = v.Nombre_archivo
		if float64(len(archivo.Value))*1.1 > maxColumn17Lenght {
			maxColumn17Lenght = float64(len(archivo.Value)) * 1.1
			sheet.SetColWidth(18, 18, maxColumn17Lenght)
		}

		ruta := roww.AddCell()
		ruta.SetStyle(style)
		ruta.Value = v.Ruta_archivo
		if float64(len(ruta.Value))*1.1 > maxColumn18Lenght {
			maxColumn18Lenght = float64(len(ruta.Value)) * 1.1
			sheet.SetColWidth(19, 19, maxColumn18Lenght)
		}

		formato := roww.AddCell()
		formato.SetStyle(style)
		formato.Value = v.Tipo_documento
		if float64(len(formato.Value))*1.1 > maxColumn19Lenght {
			maxColumn19Lenght = float64(len(formato.Value)) * 1.1
			sheet.SetColWidth(20, 20, maxColumn19Lenght)
		}

		tam := roww.AddCell()
		tam.SetStyle(style)
		tam.Value = v.Tamano
		if float64(len(tam.Value))*1.1 > maxColumn20Lenght {
			maxColumn20Lenght = float64(len(tam.Value)) * 1.1
			sheet.SetColWidth(21, 21, maxColumn20Lenght)
		}

		user := roww.AddCell()
		user.SetStyle(style)
		user.Value = v.UsuarioId
		if float64(len(user.Value))*1.1 > maxColumn21Lenght {
			maxColumn21Lenght = float64(len(user.Value)) * 1.1
			sheet.SetColWidth(22, 22, maxColumn21Lenght)
		}

		nota := roww.AddCell()
		nota.SetStyle(style)
		nota.Value = v.Nota_causas
		if float64(len(nota.Value))*1.1 > maxColumn22Lenght {
			maxColumn22Lenght = float64(len(nota.Value)) * 1.1
			sheet.SetColWidth(23, 23, maxColumn22Lenght)
		}

	}

	// ******************** ESCRITURA EXCEL *******************

	// Save the file
	err = file.Save(lbl + "/CAUSAS.xlsx")
	if err != nil {
		fmt.Printf("Failed to save file: %s\n", err)
		return "0"
	} else {
		fmt.Println("Excel file 'CAUSAS.xlsx' created successfully.")
	}

	return "1"

	// *********************************************************

}
