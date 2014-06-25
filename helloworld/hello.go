package main
import "fmt"
//import "C"

func main(){
	defer fmt.Printf("exit...")
	defer fmt.Printf("Hello World!\n")
//	C.inb(0x70)

}