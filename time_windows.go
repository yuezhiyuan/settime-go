package main

import (
	"syscall"
	"time"
	"unsafe"
)

type systemtime struct {
	wYear         uint16
	wMonth        uint16
	wDayOfWeek    uint16
	wDay          uint16
	wHour         uint16
	wMinute       uint16
	wSecond       uint16
	wMilliseconds uint16
}

func SetTime(t time.Time) error {
	time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)
	st := systemtime{
		wYear:         uint16(t.Year()),
		wMonth:        uint16(t.Month()),
		wDayOfWeek:    uint16(t.Weekday()),
		wDay:          uint16(t.Day()),
		wHour:         uint16(t.Hour()),
		wMinute:       uint16(t.Minute()),
		wSecond:       uint16(t.Second()),
		wMilliseconds: uint16(t.Nanosecond() / 1000000),
	}
	r, _, err := syscall.NewLazyDLL("kernel32.dll").NewProc("SetSystemTime").Call(uintptr(unsafe.Pointer(&st)))
	if r == 0 {
		return err
	}

	return nil
}
