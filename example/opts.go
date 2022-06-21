package main

import (
	"time"

	"github.com/franklihub/glog"
)

func main() {
	opts := []glog.Option{
		glog.WithCutOption(glog.CutPer60M),
		glog.WithAppName("appName"),
		glog.WithBufferSize(1024),
		glog.WithFlushDuration(500 * time.Millisecond),
		glog.WithWriteOption(glog.WriteByMerged),
		glog.WithLogPath("logs"),
		glog.WithWriteLevel(glog.LevelPanic | glog.LevelFatal | glog.LevelError | glog.LevelWarning | glog.LevelInfo | glog.LevelDebug),
		glog.WithStdout(true),
		glog.WithMaxTime(24 * 60 * 60),
	}
	glog.Log(opts...)
	//

	msg := "glog msg"
	glog.Debug(msg)
	glog.Info(msg)
	glog.Warn(msg)
	glog.Error(msg)
	// glog.Fatal(msg)
	glog.Panic(msg)

}
