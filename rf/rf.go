/*
simple program to search binary file for specific string,
the magic string should live in define boundary(16byte in this sample).
just test it for BIOS binary search key words.
*/
package main
import (
  "fmt"
  "os"
  "io"
  "strings"
  "flag"
)
//define seek jump boundary, seek 16 bytes every time
const (
  boundary = 16
)

func help() {
	fmt.Fprintf(os.Stderr, "Usage: %s filename magic\n", os.Args[0])
	fmt.Printf("filename: binary file name\n")
	fmt.Printf("magic: the magic key word you want to search in binary file\n")
}
/*
find the position of subs in string s
s: long string that may contain subs string
subs: key word string
*/
func seek_index(s string, subs string) int {

    n := len(subs)
    if n == 0 {
       return 0
    }
    for i := 0; i+n <= len(s); i++ {
       if s[i:i+n] == subs {
               return i
       }
    }
    return -1
}
/*
Usage:
rf.exe filename magic
filename: binary file name
magic: the magic key word you want to search in binary file
*/
func main() {

	flag.Parse()

	if flag.NArg() < 1 {
		help()
		return
	}

	fn := flag.Arg(0)
	magic := flag.Arg(1)
	fl, err := os.Open(fn)
	if err !=nil {
		fmt.Println(fn, err)
		return
	}
	defer fl.Close()
	buf := make([]byte, boundary)
	fptr := 0
	for{
		n,err := fl.Read(buf)
		if (n==0||err == io.EOF){
			fmt.Printf("magic not found!")
			break
		}
		if strings.Contains(string(buf), magic){
			i:=seek_index(string(buf), magic)
			fmt.Printf("find magic at 0x%X\n", fptr+i)
			break
		}
		fptr += boundary;
	}
}