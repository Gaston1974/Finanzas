package styles

import (
	"fyne.io/fyne/v2"
)

const sideWidth2 = 100

type BarberiaLayout2 struct {
	//left, content, right fyne.CanvasObject
}

func NewBarberiaLayout2() fyne.Layout { //(left, content, right fyne.CanvasObject) fyne.Layout {

	return &BarberiaLayout2{} //{left: left, content: content, right: right}
}

// Layout will manipulate the listed CanvasObjects Size and Position
// to fit within the specified size.
func (l *BarberiaLayout2) Layout(objects []fyne.CanvasObject, size fyne.Size) {

	objects[0].Resize(fyne.NewSize(sideWidth2, size.Height))

	objects[1].Move(fyne.NewPos(sideWidth2, 0))
	objects[1].Resize(fyne.NewSize(size.Width-sideWidth2*2, size.Height))

	//l.right.Move(fyne.NewPos(sideWidth+l.content.MinSize().Width, 0))
	objects[2].Move(fyne.NewPos(size.Width-sideWidth2, 0))
	objects[2].Resize(fyne.NewSize(sideWidth2, size.Height))

	//fmt.Printf("\nin..")
	//fmt.Printf("\nin.. \nlen: %d \nwidth: %.2f \nheight: %.2f \n vec: %v", len(objects), size.Width, size.Height, objects)
}

// MinSize calculates the smallest size that will fit the listed
// CanvasObjects using this Layout algorithm.
func (l *BarberiaLayout2) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(10, 10)
}
