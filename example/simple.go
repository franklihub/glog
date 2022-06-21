package main

import (
	"github.com/franklihub/glog"
)

func main() {
	msg := "glog msg"

	glog.Debug(msg)
	glog.Info(msg)
	glog.Warn(msg)
	glog.Error(msg)
	// glog.Fatal(msg)
	glog.Panic(msg)

}
