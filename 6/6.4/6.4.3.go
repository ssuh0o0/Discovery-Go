// 에러 처리

// **** 에러에 추가 정보 실어서 보내기 ****
type error interface {
	Error() string
}

return id, errors.New("htmlHandler: ID is empty")

//id가 음수 일때 Errorf는 error()을 호출하는데 Sprintf 와 비슷함.
return fmt.Errorf("ID %d is negative",id)


// error가 인터페이스 라는 것을 생각해서 에러 자료형 하나 만든다.
type ErrNegativeID ID

func (e ErrVegativeID) Error() string {
	return rmt.Sprintf("ID %d is negative", e)
}
var err error = ErrNegativeID(-100)
fmt.Println(err)


// **** 반복된 에러 처리 피하기 ****
// 에러 보내는 함수는 에러를 받는 함수가 어떻게 할지 모르기 때문에, 안끝내고 있는다.
// 그러나 에러 받는 함수는 끝내고 싶을 수도 있기 때문에 이러한 함수을 반복한다.
if err := f(); err != nil {
	panic(err)
}

//따라서 이러한 함수를 만들어 err가 nil값이 아닐 때는 패닉을 발생 시키면된다.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// 그럼 에러를 반환하는 f()를 호출할 때 이렇게 하면 된다.
Must(f())

// 두 값을 반환하는 함수는 이렇게 하면된다.
func Must(i int64, err error) int64 {
	if err != nil {
		panic(err)
	}
	return i
}

parsed := Must(str.conv.ParseInt("123",10, 64))

// **** 추가 정보와 함께 반복된 처리 피하기 *****
type ResponseError struct {
	Err error
	Code int // 에러 코드 번호
}

resp := Response {
	ID : id,
	Task : t,
	Error : respErr,
}
json.NewEncoder(w).Encode(resp)
w.WriteHeader(resp.Code)

// **** panic 과 recover ****
// 호출을 타고 올라가는 과정에서 더이상 패닉이 전파되지 않도록 하는 방법 : recover
// recover은 defer안에서만 효력이 발생한다.

func f() int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f",r)
		}()
		g() 
		return 100
	}
}

func g() {
	panic("I panic!")
}

// defer은 패닉 발생의 유무와 관계없이 실행된다. defer내에 recover 호출했을 때 nil이면 패닉이 발생하지 않은 경우이다.
// 위 코드가 100반환하는 부분까지 못갔기 때문에, 기본값인 0이 반환되었다.
func Example_f() {
	fmt.Println("f() = ",f())
	//output:
	//Recovered in f
	//0
}

// 반환 값에 i 라는 이름을 붙여주고, defer내에서 값을 할당할 수 있게한다.
// 그럼 패닉이 일어난 경우에만 반환값을 -1 할수 있다.
defer func() {
	if r := recover(); r != nil {
		fmt.Println("Recovered in f",r)
		i = -1
	}
}