package main

import (
	"bytes"
	"os/exec"
	"strings"
	"github.com/AidHamza/optimizers-api/log"
)

func printCommand(cmd *exec.Cmd) {
	log.Logger.Info("Executing Command", "ARGS", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		log.Logger.Error("Executing Command", "ARGS", err.Error())
	}
}

func printOutput(outs []byte) string {
	if len(outs) > 0 {
		return string(outs)
	}
	return ""
}

func compressImage(filepath, filetype string) string {
	var cmd string
	var args []string
	if filetype == "image/jpeg" {
		cmd = "jpegoptim"
		args = []string{"-s", "--max=80", "--dest=" + downloadPath, uploadPath + filepath}
	} else if filetype == "image/png" {
		cmd = "optipng"
		args = []string{"-o2", uploadPath + filepath, "-out", downloadPath + filepath}
	}

	cmdExec := exec.Command(cmd, args...)
	cmdOutput := &bytes.Buffer{}
	cmdExec.Stdout = cmdOutput
	printCommand(cmdExec)
	err := cmdExec.Run() // will wait for command to return
	printError(err)
	return printOutput(cmdOutput.Bytes())
}
