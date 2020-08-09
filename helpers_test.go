package gofake

import (
	"testing"
)

func TestRandIntRange(t *testing.T) {
	if randIntRange(5, 5) != 5 {
		t.Error("You should have gotten 5 back")
	}
}

func TestGetRandValueFail(t *testing.T) {
	for _, test := range [][]string{nil, {}, {"not", "found"}, {"person", "notfound"}} {
		if getRandValue(test) != "" {
			t.Error("You should have gotten no value back")
		}
	}
}

func TestGetRandIntValueFail(t *testing.T) {
	for _, test := range [][]string{nil, {}, {"not", "found"}, {"status_code", "notfound"}} {
		if getRandIntValue(test) != 0 {
			t.Error("You should have gotten no value back")
		}
	}
}

func TestRandFloat32RangeSame(t *testing.T) {
	if randFloat32Range(5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

func TestRandFloat64RangeSame(t *testing.T) {
	if randFloat64Range(5.0, 5.0) != 5.0 {
		t.Error("You should have gotten 5.0 back")
	}
}

func TestReplaceWithNumbers(t *testing.T) {
	if replaceWithNumbers("") != "" {
		t.Error("You should have gotten an empty string")
	}
}

func BenchmarkReplaceWithNumbers(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		Seed(42)

		b.StartTimer()
		replaceWithNumbers("###☺#☻##☹##")
		b.StopTimer()
	}
}

func TestReplaceWithNumbersUnicode(t *testing.T) {
	for _, test := range []struct{ in, should string }{
		{"#界#世#", "5界7世8"},
		{"☺#☻☹#", "☺5☻☹7"},
		{"\x80#¼#語", "\x805¼7語"},
	} {
		Seed(42)
		got := replaceWithNumbers(test.in)
		if got == test.should {
			continue
		}
		t.Errorf("for '%s' got '%s' should '%s'",
			test.in, got, test.should)
	}
}

func TestReplaceWithLetters(t *testing.T) {
	if replaceWithLetters("") != "" {
		t.Error("You should have gotten an empty string")
	}
}

func TestReplaceWithHexLetters(t *testing.T) {
	if "" != replaceWithHexLetters("") {
		t.Error("You should have gotten an empty string")
	}
}

func TestToFixed(t *testing.T) {
	floats := [][]float64{
		{123.1234567489, 123.123456},
		{987.987654321, 987.987654},
	}

	for _, f := range floats {
		if toFixed(f[0], 6) != f[1] {
			t.Fatalf("%g did not equal %g. Got: %g", f[0], f[1], toFixed(f[0], 6))
		}
	}
}

func TestFuncLookupSplit(t *testing.T) {
	tests := map[string][]string{
		"":                  {},
		"a":                 {"a"},
		"a,b,c":             {"a", "b", "c"},
		"a, b, c":           {"a", "b", "c"},
		"[a,b,c]":           {"[a,b,c]"},
		"a,[1,2,3],b":       {"a", "[1,2,3]", "b"},
		"[1,2,3],a,[1,2,3]": {"[1,2,3]", "a", "[1,2,3]"},
	}

	for input, expected := range tests {
		values := funcLookupSplit(input)
		if len(values) != len(expected) {
			t.Fatalf("%s was not %s", values, expected)
		}
		for i := 0; i < len(values); i++ {
			if values[i] != expected[i] {
				t.Fatalf("expected %s got %s", expected[i], values[i])
			}
		}
	}
}
