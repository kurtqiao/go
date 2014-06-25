package main

import "fmt"
import "os"
import "flag"
import "strconv"

func help() {
	fmt.Fprintf(os.Stderr, "Usage: %s [Bus][Dev][Func][Reg]\n", os.Args[0])
	fmt.Printf("Please input Bus/Dev/Func/Reg as Dec number\n")
	fmt.Printf("default BECREG is 15(0xF), change it by use:\n")
	fmt.Printf("%s [Bus][Dev][Func][Reg][BECREG]", os.Args[0])
}

func main() {
	flag.Usage = func () {
		help()
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() < 1 {
		help()
		return
	}

	BECREG   := 15
	pciaddr  := 0
	Busno,_  := strconv.Atoi(flag.Arg(0))
	Devno,_  := strconv.Atoi(flag.Arg(1))
	Func ,_  := strconv.Atoi(flag.Arg(2))
	Reg  ,_  := strconv.Atoi(flag.Arg(3))
    if flag.NArg() == 5{
		BECREG,_ = strconv.Atoi(flag.Arg(4))
	}

	pciaddr = BECREG << 28 + Busno << 20 + Devno << 15 + Func << 12 + (Reg & 0xfc)
	fmt.Printf("PCI MMIO address: %X\n", pciaddr)
	pciaddr = 1<<31 + Busno<<16 + Devno<<11 +Func<<8 + (Reg&0xfc)
    fmt.Printf("PCI IO address:   %X", pciaddr)
	
}