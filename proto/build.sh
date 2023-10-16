#!/bin/bash
runPath=$(cd `dirname $0`;pwd)
echo $runPath
for file in $runPath/*
do 
if test -f $file && [[ $file =~ \.proto$ ]]
then
    echo $file 
    arr=(${arr[*]} $file)
    protoc -I $runPath --go_out=../Pickon_Server/src/wjrgit.qianz.com/protobuf/go $file
    protoc -I $runPath --go_out=../battlePlatform/wjrgit.qianz.com/protobuf/go $file
fi  
done