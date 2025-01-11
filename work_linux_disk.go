package main

import (
	"golang.org/x/sys/unix"
	"os"
)

var stat unix.Statfs_t   

wd, err :=