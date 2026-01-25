package styles

import (
	"fyne.io/fyne/v2"
)

//const sideWidth = 150

type BarberiaLayout4 struct {
	leftWidth  float32
	rightWidth float32

	//left, content, right fyne.CanvasObject
}

func NewBarberiaLayout4(leftWidth, rightWidth float32) fyne.Layout { //(left, content, right fyne.CanvasObject) fyne.Layout {

	return &BarberiaLayout4{leftWidth, rightWidth} //{left: left, content: content, right: right}
}

// Layout will manipulate the listed CanvasObjects Size and Position
// to fit within the specified size.
func (l *BarberiaLayout4) Layout(objects []fyne.CanvasObject, size fyne.Size) {

	objects[0].Resize(fyne.NewSize(l.leftWidth, size.Height))

	objects[1].Move(fyne.NewPos(l.leftWidth, 0))
	objects[1].Resize(fyne.NewSize(size.Width-l.leftWidth+l.rightWidth, size.Height))

	//l.right.Move(fyne.NewPos(sideWidth+l.content.MinSize().Width, 0))
	objects[2].Move(fyne.NewPos(size.Width-l.rightWidth, 0))
	objects[2].Resize(fyne.NewSize(l.rightWidth, size.Height))

	//fmt.Printf("\nin..")
	//fmt.Printf("\nin.. \nlen: %d \nwidth: %.2f \nheight: %.2f \n vec: %v", len(objects), size.Width, size.Height, objects)
}

// MinSize calculates the smallest size that will fit the listed
// CanvasObjects using this Layout algorithm.
func (l *BarberiaLayout4) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(10, 10)
}

// --------------------------------------------------
