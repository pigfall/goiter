package iter

import(
		"testing"
)


func TestGroupAsMap(t *testing.T){
	type s struct{
		id string
		group string
		name string
	}

	values := []s{
		{
			id: "idA",
			group: "groupA",
			name: "tzzA",
		},
		{
			id: "idB",
			group: "groupA",
			name: "tzzB",
		},
		{
			id: "idC",
			group: "groupB",
			name: "tzzC",
		},
	}

	groups,err :=GroupAsMap( Slice(values), func(v s)(string){
		return v.group
	})
	if err != nil{
		t.Fatal(err)
	}

	if len(groups) != 2 {
		t.Fatal("failed")
	}
	for _,g := range groups{
		if g[0].group == "groupB"{
			if len(g) != 1{
				t.Fatal("failed")
			}
		}
		if g[0].group == "groupA"{
			if len(g) != 2{
				t.Fatal("faild")
			}
		}
	}
	t.Logf("%+v",groups)
}
