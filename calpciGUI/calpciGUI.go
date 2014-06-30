package main

import (
    "github.com/lxn/walk"
    . "github.com/lxn/walk/declarative"
    "strconv"
    "fmt"
)

func main() {
    var Busno, Devno, Funcno, Regno, outTE *walk.TextEdit
    MainWindow{
        Title:   "PCI address Calculator",
        MinSize: Size{500, 150},
        Layout:  VBox{},
        Children: []Widget{
            HSplitter{
                Children: []Widget{
            Label{Text:"Bus:"},
            TextEdit{
                AssignTo: &Busno,
            },
            Label{Text:"Dev:"},
            TextEdit{
                AssignTo: &Devno,
            },
            Label{Text:"Func:"},
            TextEdit{
                AssignTo: &Funcno,
            },
            Label{Text:"Reg:"},
            TextEdit{
                AssignTo: &Regno,
            },
            PushButton{
                Text: "Get",
                OnClicked: func() {
                    BB,_ := strconv.Atoi(Busno.Text())
                    DD,_ := strconv.Atoi(Devno.Text())
                    FF,_ := strconv.Atoi(Funcno.Text())
                    RR,_ := strconv.Atoi(Regno.Text())
                    pcimmaddr := 15 << 28 + BB << 20 + DD << 15 + FF << 12 + (RR & 0xfc)
                    pciaddr := 1<<31 + BB<<16 + DD<<11 +FF<<8 + (RR&0xfc)
                    outTE.SetText(fmt.Sprintf("pci io address: %X\r\npci mmio address: %X",pciaddr, pcimmaddr))
                },
            },
            },
            },
            TextEdit{AssignTo: &outTE, ReadOnly: true},
            
        },
    }.Run()
}