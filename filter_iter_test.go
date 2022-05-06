package iter

import(
		"testing"
)


func TestFilter(t *testing.T){
	sValue := []int{1,2,-1,}
	s := Slice(sValue)
	
	result,_ := Collect(Filter(s,func(v int)(bool,error){
		return v < 0,nil
	}))

	t.Log(result)
	if len(result) != 1 ||  result[0] != -1{
		t.Fatal("Failed")
	}

}
