#!/bin/bash
cd ~/Movies/douyin/download/
filenames=$(ls)
for file in ${filenames};do
    to_file=${file/".mp4"/".ts"}
    echo $to_file_name
    `ffmpeg -i $file -vcodec copy -acodec copy -vbsf h264_mp4toannexb $to_file`
done