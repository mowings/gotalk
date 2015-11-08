package stuff

import (
	"fmt"
	"testing"
)

func TestEnglish(t *testing.T) {
	s := "hello there"
	if out := Reverse(s); out != "ereht olleh" {
		t.Logf("got: %s", out)
		t.Fail()
	}
}

func TestChinese(t *testing.T) {
	s := "你好"
	if out := Reverse(s); out != "好你" {
		t.Logf("got: %s", out)
		t.Fail()
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("hello there")
	}
}

func ExampleReverse() {
	fmt.Println(Reverse("hello"))
	// Output: olleh
}
