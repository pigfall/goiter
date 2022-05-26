package iter

import(
		"testing"
)

func TestFlat(t *testing.T){
	values :=[][]string{
		[]string{
			"1","2",
		},
		[]string{
			"3","4",
		},
	}

	elems,err := Collect(Flatten(Slice(ListSliceToIter(values))))
	if err != nil{
		 t.Fatal(err)
	}
	if len(elems) != 4{
		t.Fatal("unexpected")
	}
	expect := []string{"1","2","3","4"}
	for i:=0;i<len(expect);i++{
		if expect[i] != elems[i]{
			t.Fatal("unexpected")
		}
	}
	t.Log(elems)

}
