//go:generate fyne bundle -o bundled.go assets

package main

import (
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"

	//"fyne.io/fyne/v2/canvas"

	"fyne.io/fyne/v2/widget"

	formularios "hello/src/pkg/formularios"
	"hello/src/pkg/styles"

	"fyne.io/fyne/v2/container"
)

type Node struct {
	name string
	icon fyne.Resource
	canv bool

	run func(*canvas.Image, *fyne.Container, string, fyne.Window, fyne.App)

	Children []*Node // Slice of pointers to Node
}

func main() {

	os.Setenv("FYNE_DRIVER", "software")

	// Initialize GLFW
	// if err := glfw.Init(); err != nil {
	// 	log.Fatalln("Failed to initialize GLFW:", err)
	// }
	// defer glfw.Terminate()

	// godotenv.Load(".env")

	// postgresURI := os.Getenv("DBURL")
	// if postgresURI == "" {
	// 	apiDatas.Log("\n No URL variable is found on the environment \n", "./logs.txt")
	// 	log.Fatal("No URL variable is found on the environment: ", postgresURI)
	// }

	// ---- create GUI app ----

	a := app.NewWithID("com.finanzas.app")
	w := a.NewWindow("Sistema de gestion de aplicaciones - version 1.0")
	icono := resourceCrypto4Png
	// ------- Create root node ---------

	root := &Node{"Consultas", icono, false, nil, nil}

	// ------ Create child nodes ------

	child1 := &Node{"Cryptos", icono, false, nil, nil}
	child2 := &Node{"Causas", icono, false, formularios.ShowCausas, nil}
	// child3 := &Node{"Baja", nil, false, formularios.ShowBaja, nil}
	// child4 := &Node{"Modificacion", nil, true, nil, nil}
	// child5 := &Node{"Consulta", nil, true, formularios.ShowConsulta, nil}
	// child7 := &Node{"Backup", nil, true, formularios.ShowBackup, nil}
	// child6 := &Node{"Salir", theme.HomeIcon(), true, formularios.ShowSalir, nil}
	// child8 := &Node{"Notificar", theme.HomeIcon(), true, formularios.ShowNotificar, nil}

	root.Children = append(root.Children, child1, child2)

	grandchild1 := &Node{"Precios - historico", nil, false, formularios.ShowCryptosInf, nil}
	grandchild2 := &Node{"Precios - comparativo", nil, false, formularios.ShowCryptosCmp, nil}
	// grandchild4 := &Node{"Contraseñas", nil, false, formularios.ShowContrasenias, nil}
	// grandchild14 := &Node{"Emprendedor", nil, false, formularios.ShowClienteUpdt, nil}
	// grandchild16 := &Node{"Feria", nil, false, formularios.ShowFeria, nil}

	child1.Children = append(child1.Children, grandchild1, grandchild2) // info
	// child4.Children = append(child4.Children, grandchild4, grandchild14, grandchild16) // modificaciones

	// ------- ventana principal ---------

	w.SetMaster()
	w.Resize(fyne.NewSize(960, 580))
	w.CenterOnScreen()

	//dd := styles.MyTheme{}
	dd := styles.NewMyTheme()
	a.Settings().SetTheme(dd)

	//i := 0

	//uri := "./narner.png"
	//u, err := storage.ParseURI(uri)
	//if err != nil {
	//	u =
	//}

	logo := canvas.NewImageFromResource(resourceCrypto4Png)
	logo2 := canvas.NewImageFromResource(resourceGNAPng)
	//logo2 := canvas.NewImageFromResource(resourceLOGOJpeg)
	//logo2 := canvas.NewText("     FERIA", color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff})
	//logo2 := widget.NewLabel("             FERIA\nDE EMPRENDEDORES")

	//logo := canvas.NewImageFromFile("./barber.gif")
	//logo := canvas.NewImageFromURI(styles.NewBarberiaIcon())
	//logo := canvas.NewImageFromURI(u)

	//move := canvas.NewPositionAnimation(fyne.NewPos(0, 0), fyne.NewPos(15, 0), time.Second, logo.Move)
	//move.AutoReverse = true
	//move.Start()

	//logo3 := canvas.NewImageFromResource(resourceCryptoPng)

	//content := container.New(styles.NewBarberiaLayout(150))
	content := container.New(styles.NewBarberiaLayout(220))
	//content4 := container.New(layout.NewFormLayout())
	//content2 := container.New(layout.NewStackLayout(), container.NewPadded(logo))
	//var objects []fyne.CanvasObject
	//objects = append(objects, container.NewPadded(logo), layout.NewSpacer(), container.NewPadded(logo))
	content2 := container.New(styles.NewBarberiaLayout(115))

	//content2.Objects = []fyne.CanvasObject{container.NewPadded(logo), layout.NewSpacer(), container.NewPadded(logo)} //canvas.NewText("	NOMBRE DEL NEGOCIO", color.White)
	content2.Objects = []fyne.CanvasObject{layout.NewSpacer(), canvas.NewText("DEPARTAMENTO DE INVESTIGACIONES DE DELITOS TECNOLÓGICOS", color.White), layout.NewSpacer()}

	//flag := &i
	id := ""

	// ----------- Armado del arbol ------------

	tree := widget.NewTree(
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			switch id {
			case "":
				return []widget.TreeNodeID{"Consultas", "Salir"}
			case "Consultas":
				return []widget.TreeNodeID{"Cryptos", "Causas"}
			case "Cryptos":
				return []widget.TreeNodeID{"Precios - historico", "Precios - comparativo"}
			}
			return []string{}
		},
		func(id widget.TreeNodeID) bool {
			return id == "" || id == "Consultas" || id == "Cryptos"
		},
		func(branch bool) fyne.CanvasObject {

			icon := &canvas.Image{}
			label := widget.NewLabel("Text Editor")
			labelHeight := label.MinSize().Height

			icon.SetMinSize((fyne.NewSize(labelHeight, labelHeight)))

			return container.NewBorder(nil, nil, icon, nil, label)

		},
		func(id widget.TreeNodeID, branch bool, obj fyne.CanvasObject) {
			img := obj.(*fyne.Container).Objects[1].(*canvas.Image)

			//if !branch {
			//		text += " (branch)"
			img.Resource = root.icon
			//}

			obj.(*fyne.Container).Objects[0].(*widget.Label).SetText(id)
			//obj.(*fyne.Container).Objects[1].(*canvas.Image). //SetResource(img.Resource)
			//o.(*widget.Label).Refresh()
		})

	// ---- seleccion / busqueda / ejecucion de operaciones -----

	tree.OnSelected = func(id widget.TreeNodeID) {
		// if *flag == 1 {

		if id == "Salir" {

			w.Close()
			//return // salir de la aplicacion

		} else {

			for _, v := range root.Children {

				if id == v.name {

					if !tree.IsBranch(id) {

						// fmt.Printf("\n branch\n id: %s", id)
						v.run(logo, content, id, w, a)

					} else {

						break
					}

				} else if tree.IsBranchOpen("Cryptos") {

					for _, t := range child1.Children {

						if id == t.name {
							t.run(logo, content, id, w, a)
						}

					}
					//tree.CloseBranch("Cryptos")

				}
			}
		}

	}

	/*

		tree.OnSelected = func(id widget.TreeNodeID) {
			// if *flag == 1 {

				if id == "Registro" {
					root.run(flag, content, id, w, a)
				} else {

					for _, v := range root.Children {

						if id == v.name {

							if !tree.IsBranch(id) {

								v.run(flag, content, id, w, a)

							} else {

								break
							}

						} else if tree.IsBranchOpen("Alta") {

							for _, t := range child2.Children {

								if id == t.name {
									t.run(flag, content, id, w, a)
								}

							}
							tree.CloseBranch("Alta")

						} else if tree.IsBranchOpen("Modificacion") {

							for _, t := range child4.Children {

								if id == t.name {
									t.run(flag, content, id, w, a)
								}

							}
							tree.CloseBranch("Modificacion")

						}
					}

				}

			// } else {
			// 	content.Objects = []fyne.CanvasObject{layout.NewSpacer(), canvas.NewText("Aun no se encuentra identificado en el sistema", color.White), layout.NewSpacer()}
			// 	fun.Wait(2)
			// 	formularios.ShowLogin(flag, content, id, w, a)
			// 	tree.UnselectAll()
			// }
		}

	*/

	// ---- armado layout -----
	formularios.ShowStart(logo2, content, id, w, a)
	logo.FillMode = canvas.ImageFillContain

	split1 := container.NewVSplit(content, content2)

	split1.Offset = 0.90 // determina como se reparte la division
	split1.Refresh()

	split2 := container.NewHSplit(tree, split1)
	split2.Offset = 0.35 // determina como se reparte la division
	split2.Refresh()

	w.SetContent(split2)
	//w.Resize(fyne.NewSize(961, 580))
	split2.Refresh()

	w.ShowAndRun()

}
