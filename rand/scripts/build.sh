#!/bin/bash

#########################
# Build all executables #
#########################

. $(dirname "$0")/cfg.sh

echo " Script  : $SCPT"
echo " Project : $PROJ"

echo
echo "# building $PROJ/cmd/..."

cd $PROJ/cmd/fast
if [ -f pgo/cpu.pprof ]; then
  echo "Building with './pgo/cpu.pprof'..."
  go build -pgo=pgo/cpu.pprof
else
  go build
fi
ls -l fast

cd $PROJ/cmd/old
if [ -f pgo/cpu.pprof ]; then
  echo "Building with 'pgo/cpu.pprof'..."
  go build -pgo=pgo/cpu.pprof
else
  go build
fi
ls -l old

cd $PROJ/cmd/ver2
if [ -f pgo/cpu.pprof ]; then
  echo "Building with 'pgo/cpu.pprof'..."
  go build -pgo=pgo/cpu.pprof
else
  go build
fi
ls -l ver2

cd $SCPT
