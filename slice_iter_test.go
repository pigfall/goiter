package iter

import(
	"fmt"
		"testing"
)

func TestForEach(t *testing.T){
	slice := []string{"12","34"}

	err := ForEach(Map(Slice(slice),func(from string)(int,error){
		return 0,fmt.Errorf("stop")
	}),func(i int)error{
		t.Log(i)
		return nil
	})
	if err != nil{
		t.Log(err)
	}
}
