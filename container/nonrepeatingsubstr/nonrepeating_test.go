package substr

import "testing"

/**
测试：go test .
结果：
=== RUN   TestSubstr
--- PASS: TestSubstr (0.00s)
PASS

代码覆盖率：go test -coverprofile=c.out
结果：
PASS
coverage: 92.9% of statements
ok      gostudy/container/nonrepeatingsubstr    0.256s
*/
func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		//Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},

		//Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbbb", 1},
		{"abcabcabcd", 4},

		//Chinese cases
		{"这里是慕课网", 6},
		{"一二三二一", 3},
	}

	for _, tt := range tests {
		if actual := nonRepeatingSubstr(tt.s); actual != tt.ans {
			t.Errorf("got %d input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

/**
性能测试：go test -bench .
结果：
goos: windows
goarch: amd64
pkg: gostudy/container/nonrepeatingsubstr
BenchmarkSubstr
BenchmarkSubstr-8   	  802208	      1321 ns/op
PASS

性能调优：go test -bench . -cpuprofile cpu.out
结果：
goos: windows
goarch: amd64
pkg: gostudy/container/nonrepeatingsubstr
BenchmarkSubstr-8        1048669              1149 ns/op
PASS
ok      gostudy/container/nonrepeatingsubstr    2.529s

性能调优（工具）：go tool pprof cpu.out   ---> web

*/
func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"
	ans := 8

	for i := 0; i < b.N; i++ {
		actual := nonRepeatingSubstr(s)
		if actual != ans {
			b.Errorf("got %d input %s; expected %d", actual, s, ans)
		}
	}
}
