package styles

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MyTheme2 struct {
	fyne.Theme
	//Ins *widget.Entry
}

//var _ fyne.Theme = (*MyTheme)(nil)

func NewMyTheme2() fyne.Theme { //(left, content, right fyne.CanvasObject) fyne.Layout {

	return &MyTheme{Theme: theme.DefaultTheme()} //{left: left, content: content, right: right}
}

func (m *MyTheme2) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {

	//green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	//yellow := color.NRGBA{R: 0xff, G: 0x98, B: 0x00, A: 0xff}
	yellow := color.NRGBA{R: 0xf3, G: 0xf3, B: 0xf3, A: 0xff}

	if name == theme.ColorNameBackground {
		return m.Theme.Color(name, theme.VariantDark)

	}

	if name == theme.ColorNameForeground {
		return yellow

	}

	if name == theme.ColorNameInputBackground {
		return color.Black

	}

	if name == theme.ColorNameButton {
		//return color.NRGBA{R: 0x28, G: 0x29, B: 0x2e, A: 0xff}
		return color.NRGBA{R: 0, G: 0, B: 0, A: 255}

	}

	if name == theme.ColorNameMenuBackground {
		return color.NRGBA{R: 0, G: 0, B: 180, A: 255}

	}

	return theme.DefaultTheme().Color(name, variant)

}

func (m MyTheme2) Icon(name fyne.ThemeIconName) fyne.Resource {
	if name == theme.IconNameHome {
		return nil
		//fyne.NewStaticResource("myHome", homeBytes)    // resources creados por mi
	}

	return theme.DefaultTheme().Icon(name)
}

func (m MyTheme2) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m MyTheme2) Size(name fyne.ThemeSizeName) float32 {

	switch name {
	case theme.SizeNameText:
		return 10
	}

	return theme.DefaultTheme().Size(name)
}
