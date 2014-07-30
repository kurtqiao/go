package main
import (
  "fmt"
  //"strings"
)
func main(){
	ostr := "test string!"
	fmt.Printf("string length: %d\n", len(ostr))

    fmt.Printf("%s",invert_str(ostr))
}

func invert_str(ostr string) (rstr string){
	var tstr string
	for i:= len(ostr)-1; i>=0; i--{
		tstr += string(ostr[i])
	}
	return tstr
}