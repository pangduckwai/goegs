#!/bin/bash

##############################
# Run flows on remote server #
##############################

. $(dirname "$0")/cfg.sh

if [ $# -eq 1 ]; then
  export RAND_RUN_NUM=$1
fi

echo " Script  : $SCPT"
echo " Project : $PROJ"
echo " Runs.   : $RAND_RUN_NUM"

echo
echo "# start running rand benchmarks..."

cd $PROJ

./cmd/old/old 1
sleep 1
./cmd/old/old 2
sleep 1
./cmd/ver2/ver2 1
sleep 1
./cmd/ver2/ver2 2
sleep 1
./cmd/ver2/ver2 4
sleep 1
./cmd/fast/fast 1
sleep 1
./cmd/fast/fast 2
sleep 1
./cmd/fast/fast 4

cd $SCPT
