package main

import "fmt"
import "os"
import "flag"
import "strconv"
//please check intel host bridge for PCIEBAR (0:0:0:60)
//usualy it's 0xE0000000
func help() {
	fmt.Fprintf(os.Stderr, "Usage: %s [Bus][Dev][Func][Reg]\n", os.Args[0])
	fmt.Printf("Please input Bus/Dev/Func/Reg as Dec number\n")
	fmt.Printf("default PCIEBAR is 14(0xE), change it by use:\n")
	fmt.Printf("%s [Bus][Dev][Func][Reg][PCIEBAR]", os.Args[0])
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

	PCIEBAR   := 14
	pciaddr  := 0
	Busno,_  := strconv.Atoi(flag.Arg(0))
	Devno,_  := strconv.Atoi(flag.Arg(1))
	Func ,_  := strconv.Atoi(flag.Arg(2))
	Reg  ,_  := strconv.Atoi(flag.Arg(3))
    if flag.NArg() == 5{
		PCIEBAR,_ = strconv.Atoi(flag.Arg(4))
	}

	pciaddr = PCIEBAR << 28 + Busno << 20 + Devno << 15 + Func << 12 + (Reg & 0xfc)
	fmt.Printf("PCI MMIO address: %X\n", pciaddr)
	pciaddr = 1<<31 + Busno<<16 + Devno<<11 +Func<<8 + (Reg&0xfc)
    fmt.Printf("PCI IO address:   %X", pciaddr)
	
}