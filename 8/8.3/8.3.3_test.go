// 상속
package main

import "fmt"

// IsA 관계와 HasA 관계가 있다.
//
// HasA )
// 상속보다는 객체 구성이 더 많다.
// 재 사용하고자 하는 구현의 자료형의 변수를 struct에 내장하면 된다.
// 즉, HasA 관계는 그 구현을 필드로 가짖고 있으면 되는 것이다.
//
// IsA )
// 대부분의 경우에 추상클래스를 상속한다.
// 추상 클래스 중 구현이 없는 경우는 go 언어의 인터페이스를 이용하면 된다.
// 추상클래스가 아닌 클래스를 상속받는 경우는
// 인터페이스와 구현을 함께 상속한다.
// 안토페이스와 구조체 내장을 동시에 하면 가능하다.

// **** 메서드 추가 ****
// 기존 코드를 재사용하면서 기능 추가를 하고 싶은 경우에 상속

// 다형성 예제에서 둘레를 구하는 기능을 추가하고자 한다.
type Rectangle struct {
	Width, Height float32
}

func (r Rectangle) Area() float32 {
	return r.Width * r.Height
}

// 5장에서 배웠던 구조체 내장을 사용하자.
type RectangleCircum struct{ Rectangle }

func (r RectangleCircum) Circum() float32 {
	return 2 * (r.Width + r.Height)
}

// 이제 모든 기능을 사용할 수 있다.

func ExampleRectangleCircum() {
	r := RectangleCircum{Rectangle{3, 4}}
	fmt.Println(r.Area())
	fmt.Println(r.Circum())
	//output:
	//12
	//14
}

// 상속과 함께 생성자도 만들어줌
// func NewRectangleCircum(width, height float32) *RectangleCircum {
// 	return &RectangleCircum(Rectangle{width, height})
// }

// **** 오버라이딩 ****
// 기존의 구현을 다른 구현으로 대체한다

// 기존 넓이의 2배를 출력해보자
type WrongRectangle struct{ Rectangle }

func (r WrongRectangle) Area() float32 {
	return r.Rectangle.Area() * 2
}

func ExampleWrongRectangle() {
	r := WrongRectangle{Rectangle{3, 4}}
	fmt.Println(r.Area())
	//output:24
}

// **** 서브타입 ****
// 기존 객체가 쓰이던 곳에 상속받은 객체를 쓰고자 상속하기도 한다.

type Shape interface {
	Area() float32
}

type Square struct {
	Size float32
}

func (s Square) Area() float32 {
	return s.Size * s.Size
}

type Triangle struct {
	Width, Height float32
}

func (t Triangle) Area() float32 {
	return 0.5 * t.Width * t.Height
}

func TotalArea(shapes []Shape) float32 {
	var total float32
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

func ExampleTotakArea() {
	fmt.Println(TotalArea([]Shape{
		Square{3},
		Rectangle{4, 5},
		Triangle{6, 7},
		RectangleCircum{Rectangle{8, 9}},
		WrongRectangle{Rectangle{1, 2}},
	}))
	//output: 126
}

func check() {
	//impl 에는 shape 인터페이스를 구현하는지 여부가 기록됨
	impl := reflect.TypeOf(RectangleCircum{}).Implements{
		reflect.TypeOf((*Shape)(nil)).Elem()
	}
	field, ok := reflect.TypeOf(RectangleCircum{}).FieldByName("Rectangle")
	//rectangle이 rectanglecircum에 내장되어 있는지 여부가 기록됨
	emb := ok && field.Anonymous && field.Type == reflect.TypeOf(Rectangle{})
}
