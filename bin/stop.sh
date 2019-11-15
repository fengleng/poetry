#!/bin/sh
dir="/work/poetry/bin/server.pid"
pid=`cat $dir`

kill -QUIT $pid


#kill -3 $pid
#`kill -QUIT $pid`
echo "ok"
