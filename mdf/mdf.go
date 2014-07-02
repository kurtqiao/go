package main

import (
    "github.com/lxn/walk"
    "crypto/md5"
    "hash/crc32"
    "fmt"
    "io"
    "os"
    "flag"
)

func main() {
    flag.Parse()
    fn := flag.Arg(0)

    fl, err := os.Open(fn)
    if err !=nil {
        fmt.Println(fn, err)
        return
    }
    defer fl.Close()
    fmt.Printf("%s:\n", fn)
    md5h := md5.New()
    io.Copy(md5h, fl)
    fmt.Printf("MD5: 0x%X\n", md5h.Sum([]byte("")))
    c := crc32.NewIEEE()
    fl.Seek(0, os.SEEK_SET)  //reset file pointer to start
    io.Copy(c, fl)
    crch:=c.Sum32()
    fmt.Printf("CRC32: 0x%X", crch)
    //set crc32 result to windows clipboard
    walk.Clipboard().SetText(fmt.Sprintf("0x%X", crch))
}