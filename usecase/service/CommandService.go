package service

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

type CommandService struct {
}

func NewCommandService() *CommandService {
	return &CommandService{}
}

func (s *CommandService) ListProcess() (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("tasklist")
	} else {
		cmd = exec.Command("ps", "aux")
	}
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	fmt.Println(string(output))
	return string(output), nil
}

func (s *CommandService) KillByPid(sPid string) error {
	pid, err := strconv.Atoi(sPid)
	if err != nil {
		return err
	}
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid))
	} else {
		cmd = exec.Command("kill", strconv.Itoa(pid))
	}
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Lỗi khi chạy lệnh: %v", err)
	}
	return nil
}

func (s *CommandService) KillByName(name string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("taskkill", "/F", "/IM", name)
	} else {
		cmd = exec.Command("pkill", name)
	}
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Lỗi khi chạy lệnh: %v", err)
	}
	return nil
}

func (s *CommandService) ExcuteScript(path string) (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("powershell", "-File", path)
	} else {
		cmd = exec.Command("/bin/bash", path)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	return "", nil
}

func (s *CommandService) OpenFilePath(path string) (interface{}, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "start", path)
	} else {
		cmd = exec.Command("xdg-open", path)
	}
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return nil, nil
}
