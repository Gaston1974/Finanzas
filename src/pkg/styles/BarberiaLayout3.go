package styles

import (
	"fyne.io/fyne/v2"
)

const sideWidth3 = 80

type BarberiaLayout3 struct {
	//left, content, right fyne.CanvasObject
}

func NewBarberiaLayout3() fyne.Layout { //(left, content, right fyne.CanvasObject) fyne.Layout {

	return &BarberiaLayout3{} //{left: left, content: content, right: right}
}

// Layout will manipulate the listed CanvasObjects Size and Position
// to fit within the specified size.
func (l *BarberiaLayout3) Layout(objects []fyne.CanvasObject, size fyne.Size) {

	objects[0].Resize(fyne.NewSize(sideWidth3, size.Height))

	objects[1].Move(fyne.NewPos(sideWidth3, 0))
	objects[1].Resize(fyne.NewSize(100, size.Height))

	//l.right.Move(fyne.NewPos(sideWidth+l.content.MinSize().Width, 0))
	objects[2].Move(fyne.NewPos(180, 0))
	objects[2].Resize(fyne.NewSize(size.Width-180, size.Height))

	//fmt.Printf("\nin..")
	//fmt.Printf("\nin.. \nlen: %d \nwidth: %.2f \nheight: %.2f \n vec: %v", len(objects), size.Width, size.Height, objects)
}

// MinSize calculates the smallest size that will fit the listed
// CanvasObjects using this Layout algorithm.
func (l *BarberiaLayout3) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(10, 10)
}
