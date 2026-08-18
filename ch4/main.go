package main

import (
	"learning/ch4-logger/pocketlog"
	"os"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(os.Stdout))

	lgr.Debugf("Debugging My Way In Life")
	lgr.Infof("Looking For Some Information")
	lgr.Errorf("Found Errors In My Ways!!!")

}
