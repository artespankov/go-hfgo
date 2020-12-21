package prose

import (
	"fmt"
	"testing"
)

type testData struct {
	elements []string
	want string
}

func TestJoinWithCommas(t *testing.T){
	tests := []testData{
		{elements: []string{}, want: ""},
		{elements: []string{"apple"}, want: "apple"},
		{elements: []string{"apple", "orange"}, want: "apple and orange"},
		{elements: []string{"apple", "orange", "peach"}, want: "apple, orange, and peach"},
	}

	for _, test := range tests {
		got := JoinWithCommas(test.elements)
		if got != test.want{
			t.Error(errorString(test.elements, got, test.want))
		}
	}
}

func TestOneElement(t *testing.T){
	elements := []string{"apple"}
	want := "apple"
	got := JoinWithCommas(elements)
	if got != want{
		t.Error(errorString(elements, got, want))
	}
}

func TestTwoElements(t *testing.T){
	elements := []string{"apple", "orange"}
	want := "apple and orange"
	got := JoinWithCommas(elements)
	if got != want {
		t.Error(errorString(elements, got, want))
	}
}

func TestThreeElements(t *testing.T){
	elements := []string{"apple", "orange", "peach"}
	want := "apple, orange, and peach"
	got := JoinWithCommas(elements)
	if got != want {
		t.Error(errorString(elements, got, want))
	}
}


func errorString(elements []string, got string, want string) string{
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", elements, got, want)
}