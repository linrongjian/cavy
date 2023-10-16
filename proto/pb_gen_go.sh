pb_files=""
for file in `ls`; do
  if [[ $file =~ "proto" ]];then
    pb_files="${pb_files} ${file}"
  fi
done
protoc --go_out=./../go ${pb_files}


# protoc --go_out=../go protocol.proto