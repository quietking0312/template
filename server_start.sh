#!/bin/bash

service_name="./server"
service_chmod="chmod 777 $service_name"
log_name=server.log
pid_file=server.pid

SCRIPT=$0
OPERATOR=$1
is_pid=0

usage() {
  echo "Usage: sh $SCRIPT [app_name] [start|stop|restart]"
  exit 1
}

if [ $# != 1 ]; then
  usage
fi

function isExist() {
    if [ ! -f $pid_file ]; then
      is_pid=1
    fi
}

function start() {
  if [ $is_pid -ne 1 ]; then
    echo "$pid_file is exist"
    exit 1
  fi
  ${service_chmod} &
  nohup ${service_name} admin -c datacore.toml >${log_name} 2>&1 &
  # shellcheck disable=SC2181
  if [[ $? -eq 0 ]]; then
    echo $! > ${pid_file}
  else exit 1
  fi
}

function stop() {
    if [ $is_pid -eq 1 ]; then
      echo "$pid_file not is exist"
      exit 1
    fi
    # shellcheck disable=SC2046
    kill -9 $(cat ${pid_file})
    # shellcheck disable=SC2181
    if [[ $? -eq 0 ]]; then
      rm -f ${pid_file}
    else exit 1
    fi
}

function run() {
    isExist
    case "$OPERATOR" in
    "start")
      start
      ;;
    "stop")
      stop
      ;;
    "restart")
      stop
      isExist
      start
      ;;
    *)
      usage
      ;;
    esac
}

run

