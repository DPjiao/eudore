package test

import (
	"fmt"
	"github.com/eudore/eudore"
	"syscall"
	"testing"
	"time"
)

func TestSignal(t *testing.T) {
	e := eudore.NewEudore()
	e.RegisterSignal(syscall.Signal(0x01), true, func() error {
		t.Log("0x01")
		return nil
	})
	e.RegisterSignal(syscall.Signal(0x00), true, func() error {
		t.Log("222")
		return nil
	})
	e.RegisterSignal(syscall.Signal(0x00), true, func() error {
		t.Log("1111")
		return nil
	})
	e.RegisterSignal(syscall.Signal(0x00), false, func() error {
		t.Log("3333")
		return fmt.Errorf("error test 333")
	})
	t.Log(e.HandleSignal(syscall.Signal(0x00)))
	t.Log(e.HandleSignal(syscall.Signal(0x01)))
	time.Sleep(1 * time.Second)
}
