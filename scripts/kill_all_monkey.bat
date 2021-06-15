adb shell "pids=`top -n 1 -d 1 | grep monkey | busybox awk '{print $1}'`;array=(${pids});for var in ${array[@]};do kill $var;done"
