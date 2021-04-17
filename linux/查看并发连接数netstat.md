netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a,S[a]}'

while true;do  netstat -n | awk '/^tcp/ {++S[$NF]} END {for(a in S) print a,S[a]}' ;sleep 2; echo =========================== ;done

https://www.cnblogs.com/ggjucheng/archive/2012/01/08/2316661.html