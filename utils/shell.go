package utils

import (
	"os"
	"os/exec"
	"strings"
)

func ExecShellCommand(bin string, commandArgs []string) error{
	cmd := exec.Command(bin, commandArgs...)
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	return err
}

const (
	mp4Tool = "MP4Box"
	ffmpegTool = "ffmpeg"
)

func Mp4ToTs(videoPath string, toVideoPath string) error{
	var commandArgs = strings.Split("-i 1.mp4 -vcodec copy -acodec copy -vbsf h264_mp4toannexb 1.ts", " ")
	commandArgs[1] = videoPath
	commandArgs[len(commandArgs) - 1] = toVideoPath
	return ExecShellCommand(ffmpegTool, commandArgs)
}

func CombineVideosUseMp4(videoPaths []string, toPath string) error{
	var commandArgs []string
	for _, videoPath := range videoPaths{
		commandArgs = append(commandArgs, "-cat")
		commandArgs = append(commandArgs, videoPath)
	}
	commandArgs = append(commandArgs, "-new")
	commandArgs = append(commandArgs, toPath)
	return ExecShellCommand(mp4Tool, commandArgs)
}

func CombineVideosUseFfmpeg(videoPaths []string, toPath string) error{
	var commandArgs = []string{"-i"}
	var contacts = "concat:"
	for index, videoPath := range videoPaths{
		contacts += videoPath
		if index != len(videoPaths) - 1{
			contacts += "|"
		}
	}
	commandArgs = append(commandArgs, contacts)
	commandArgs = append(commandArgs,
		strings.Split("-acodec copy -vcodec copy -absf aac_adtstoasc " + toPath , " ")...)

	return ExecShellCommand(ffmpegTool, commandArgs)
}

func CheckVideo(videoPath string) error{
	checkVideoCommand := "-v error -i " + videoPath + " -f null -"
	return ExecShellCommand(ffmpegTool, strings.Split(checkVideoCommand, " "))
}

func DeleteFile(filePath string) {
	var args = make([]string, 1)
	args[0] = filePath
	_ = ExecShellCommand("rm", args)
}
