package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}
	int1 := &Integer{Value: 1}
	int2 := &Integer{Value: 1}
	btrue1 := &Boolean{Value: true}
	btrue2 := &Boolean{Value: true}
	bfalse1 := &Boolean{Value: false}
	bfalse2 := &Boolean{Value: false}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
	if int1.HashKey() != int2.HashKey() {
		t.Errorf("integers with same content have different hash keys")
	}
	if btrue1.HashKey() != btrue2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}
	if bfalse1.HashKey() != bfalse2.HashKey() {
		t.Errorf("booleans with same content have different hash keys")
	}
	if btrue1.HashKey() == bfalse1.HashKey() {
		t.Errorf("booleans with different content have same hash keys")
	}
}
