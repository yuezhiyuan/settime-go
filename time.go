//go:build !windows

package main

import (
	"syscall"
	"time"
)

func SetTime(t time.Time) error {
	// 生成timeval结构体
	tv := syscall.NsecToTimeval(t.UnixNano())

	// 调用Settimeofday函数
	if err := syscall.Settimeofday(&tv); err != nil {
		return err
	}

	return nil
}
