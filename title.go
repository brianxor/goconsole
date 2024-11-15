package goconsole

import (
	"fmt"
	"syscall"
	"unsafe"
)

var setConsoleTitleAddress uintptr

func init() {
	var err error

	kernel32Handle, err := syscall.LoadLibrary("Kernel32.dll")

	if err != nil {
		panic(err)
	}

	defer syscall.FreeLibrary(kernel32Handle)

	setConsoleTitleAddress, err = syscall.GetProcAddress(kernel32Handle, "SetConsoleTitleW")

	if err != nil {
		panic(err)
	}
}

func SetTitle(title string) error {
	titlePtr, err := syscall.UTF16PtrFromString(title)

	if err != nil {
		return err
	}

	ret, _, err := syscall.SyscallN(setConsoleTitleAddress, uintptr(unsafe.Pointer(titlePtr)))

	if ret == 0 {
		return fmt.Errorf("failed to set console title: %w", err)
	}

	return nil
}
