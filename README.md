# Goiter
An iterator library for golang, inspired by rust iter.

# Usage
## `filter`
```
// build iterator from slice
it := iter.Slice([]string{"void,"c","cp"})

// collect iterator into slice
filterd,_ := iter.Collect(
    // filter with specific prefix 
    iter.Filter(
      it,
      func(v string)(include bool,err error){
      if strings.HasPrefix(v,"c"){
      return true,nil
      }
      return false,nil
      },
      ),
    )

fmt.Println(filterd) // "void"
```
