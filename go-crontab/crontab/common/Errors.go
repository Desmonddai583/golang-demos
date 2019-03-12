package common

import "errors"

var (
	// ErrLockAlreadyRequired 锁已被占用
	ErrLockAlreadyRequired = errors.New("锁已被占用")

	// ErrNoLocalIPFound 没有找到网卡IP
	ErrNoLocalIPFound = errors.New("没有找到网卡IP")
)
