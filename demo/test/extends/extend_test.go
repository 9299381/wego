package extends

import (
	"fmt"
	"testing"
)

func TestExtend(t *testing.T) {
	p := &pc{}
	testInterface(p)
}

func testInterface(usb usbInterface) {
	usb.SayHello()
	usb.DoSth()
}

type usbInterface interface {
	SayHello()
	DoSth()
}
type usb struct {
}

func (i *usb) SayHello() {
	fmt.Println("usb....sayhello")
}

func (i *usb) DoSth() {
	fmt.Println("usb....do ")
}

type pc struct {
	*usb //可以直接引用usb的 dosth 无需new
}

func (i *pc) SayHello() {
	fmt.Println("pc....sayhello")
}
