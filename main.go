package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

//validate if the given dirPath is mounted and is a cvmfs filesystem or not
func validate(dirPath string) (err error) {
	output, err := exec.Command("df", dirPath).Output()
	if err != nil { return err }
	if !strings.Contains(string(output), "cvmfs") {
		err = errors.New("given path is not a cvmfs filesystem")
	}
	return err
}

func main() {

	var (
		LowerDir = "/cvmfs/unpacked.cern.ch"
		UpperDir = "./upperdir"
		MountDir = "/GSoC/unpacked.cern.ch"
	)

	//check if LowerDir is correctly mounted
	err := validate(LowerDir)
	if err != nil {
		fmt.Println(fmt.Errorf("cvmfs filesystem not mounted"))
		os.Exit(1)
	}

	//make sure the directory with path MountDir exists
	if _, err := os.Stat(MountDir); os.IsNotExist(err) {
		if err := os.MkdirAll(MountDir, 0755); err != nil {
			fmt.Println(fmt.Errorf("%v",err))
			os.Exit(1)
		}
	}

	//create an overlay mount of the two dirs

	//omitting "upperdir=" and "workdir=" as the overlay will be read-only
	params := fmt.Sprintf("lowerdir=%s:%s",LowerDir,UpperDir)
	if err := syscall.Mount("",MountDir,"overlay",0,params); err != nil {
		fmt.Println(fmt.Errorf("error: %v",err))
	}
}
