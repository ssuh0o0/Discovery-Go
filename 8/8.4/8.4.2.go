// 8.4.2 추상 팩토리 패턴

// 팩토리들을 여럿 묶어놓은 팩토리를 추상화하는 패턴

package main

import "fmt"

type Button interface {
	Paint()
	OnClick()
}

type Label interface {
	Paint()
}

type WinButton struct{}

func (WinButton) Paint()   { fmt.Println("win button paint") }
func (WinButton) OnClick() { fmt.Println("win button click") }

type WinLabel struct{}

func (WinLabel) Paint() { fmt.Println("win label paint") }

type MacButton struct{}

func (MacButton) Paint()   { fmt.Println("mac button paint") }
func (MacButton) OnClick() { fmt.Println("mac button click") }

type MacLabel struct{}

func (MacLabel) Paint() { fmt.Println("mac label paint") }

type UIFactory interface {
	CreateButton() Button
	CreateLabel() Label
}

type WinFactory struct{}

func (WinFactory) CreateButton() Button {
	return WinButton{}
}

func (WinFactory) CreateLabel() Label {
	return WinLabel{}
}

type MacFactory struct{}

func (MacFactory) CreateButton() Button {
	return MacButton{}
}

func (MacFactory) CreateLabel() Label {
	return MacLabel{}
}

func CreateFactory(os string) UIFactory {
	if os == "win" {
		return WinFactory{}
	} else {
		return MacFactory{}
	}
}

func Run(f UIFactory) {
	button := f.CreateButton()
	button.Paint()
	button.OnClick()
	label := f.CreateLabel()
	label.Paint()
}

func main() {
	f1 := CreateFactory("win")
	Run(f1)
	f2 := CreateFactory("mac")
	Run(f2)
}
