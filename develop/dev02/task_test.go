package main

import (
	"testing"
)

func TestUnpackString(t *testing.T) {
	val, err := UnpackString("a4bc2d5e")
	if err != nil {
		t.Fatal(err)
	}
	if val != "aaaabccddddde" {
		t.Fatal("Неккоректный результат")
	}

	val, err = UnpackString("abcd")
	if err != nil {
		t.Fatal(err)
	}
	if val != "abcd" {
		t.Fatal("Неккоректный результат")
	}

	val, err = UnpackString("")
	if err != nil {
		t.Fatal(err)
	}
	if val != "" {
		t.Fatal("Неккоректный результат")
	}

	val, err = UnpackString("45")
	if err == nil {
		t.Fatal(err)
	}
}
