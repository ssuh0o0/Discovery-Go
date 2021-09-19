
// 이 함수는 정수형만 비교할 수 있음.
func assertIntEqual(t *testing.T, a, b int) {
	if a != b {
		t.Errorf("assertion failed: %d != %d" , a,b)
	}
}

//구조체와 배열을 이용해 테이블 기반 테스트

func TestFib(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{0,0},
		{5,5},
		{6,8},
	}

	for _, c := range cases {
		got := seq.Fib(c.in)
		if got !- c.want {
			t.Errorf("Fib(%d) == %d, want : %d", c.in, got, c.want)
		}
	}
}