package iter

import(
	"fmt"
		"testing"
)

func TestForEach(t *testing.T){
	slice := []string{"12","34"}

	err := ForEach(Map(Slice(slice),func(from string)(int,error){
		return 0,fmt.Errorf("stop")
	}),func(i int){
		t.Log(i)
	})
	if err != nil{
		t.Log(err)
	}
}
