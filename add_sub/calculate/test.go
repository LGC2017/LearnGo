package calculate

import "testing"

func testadd_sub(t *testing.T){
	r := add(2,8)
	if r != 10{
		t.Fatalf("add(2,8)	error,	expect:%d,	actual:%d", 10,r)
	}
	s := sub(8,2)
	if s != 2{
		t.Fatalf("sub(8,2) error, expect:%d,	actual:%d", 6,s)
	}
	t.Logf("test add and sub sucess")
}
