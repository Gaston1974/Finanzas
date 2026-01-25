package styles

import (
	"fmt"

	"fyne.io/fyne/v2"
)

//const sideWidth = 150

type BarberiaLayout struct {
	sideWidth float32
	//left, content, right fyne.CanvasObject
}

func NewBarberiaLayout(sideWidth float32) fyne.Layout { //(left, content, right fyne.CanvasObject) fyne.Layout {

	return &BarberiaLayout{sideWidth} //{left: left, content: content, right: right}
}

// Layout will manipulate the listed CanvasObjects Size and Position
// to fit within the specified size.
func (l *BarberiaLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {

	objects[0].Resize(fyne.NewSize(l.sideWidth, size.Height))

	objects[1].Move(fyne.NewPos(l.sideWidth, 0))
	objects[1].Resize(fyne.NewSize(size.Width-l.sideWidth*2, size.Height))

	//l.right.Move(fyne.NewPos(sideWidth+l.content.MinSize().Width, 0))
	objects[2].Move(fyne.NewPos(size.Width-l.sideWidth, 0))
	objects[2].Resize(fyne.NewSize(l.sideWidth, size.Height))

	//fmt.Printf("\nin..")
	//fmt.Printf("\nin.. \nlen: %d \nwidth: %.2f \nheight: %.2f \n vec: %v", len(objects), size.Width, size.Height, objects)
}

// MinSize calculates the smallest size that will fit the listed
// CanvasObjects using this Layout algorithm.
func (l *BarberiaLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(10, 10)
}

// --------------------------------------------------

type BarberiaIcon struct {
	fmt.Stringer
}

func NewBarberiaIcon() fyne.URI {

	return &BarberiaIcon{}
}

func (b *BarberiaIcon) Extension() string {
	return "gif"
}

func (b *BarberiaIcon) Name() string {
	return "Qs1N.gif"
}

func (b *BarberiaIcon) MimeType() string {
	return ""
}

func (b *BarberiaIcon) Scheme() string {
	return "https"
}

func (b *BarberiaIcon) Authority() string {
	return "i.gifer.com"
}

func (b *BarberiaIcon) Path() string {
	return "/Qs1N.gif"
}

func (b *BarberiaIcon) Query() string {
	return ""
}

func (b *BarberiaIcon) Fragment() string {
	return ""
}
