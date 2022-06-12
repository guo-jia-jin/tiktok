package middleware

import (
	"os/exec"
)

func SaveFirstFrame(videopath string, coverpath string) (err error) {
	command := "E:\\ffmpeg\\bin\\ffmpeg.exe"
	cmd := exec.Command(command, "-i", videopath, "-ss", "00:00:01", "-vframes", "1", coverpath)
	// fmt.Printf("cmd: %v\n", cmd)
	// var stderr bytes.Buffer
	// cmd.Stderr = &stderr
	if err = cmd.Run(); err != nil {
		return err
	}
	return
}
