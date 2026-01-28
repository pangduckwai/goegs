#!/bin/bash

#############################################
# Prepare the environment for other scripts #
#############################################

#############
# Detect OS
Detect() {
	uout=$(uname -s)
	case $uout in
		Linux*)
			rslt=linux
			;;
		Darwin*)
			rslt=mac
			;;
		CYGWIN*)
			rslt=cygwin
			;;
		*)
			rslt=$out
			;;
	esac
	echo $rslt
}

########################################
# Expand a given path to absolute path
Dir() {
	pthc=$(dirname "$1")
	if [[ "$pthc" == "." ]]; then
		rslt=$(pwd)
	else
		if [[ $0 == /* ]]; then
			rslt=$pthc
		else
			rslt="$(pwd)/$pthc"
		fi
	fi
	echo $rslt
}

# Environment variables
OSYS=$(Detect)
SCPT=$(Dir "$0")
SELF=$(basename "$0")
PROJ=$(dirname "$SCPT")
NOW=$(date '+%Y%m%d%H%M')

case $OSYS in
	linux)
		USR=$(stat -c "%U" $PROJ)
		GRP=$(stat -c "%G" $PROJ)
		;;
	mac)
		USR=$(stat -f "%u" $PROJ)
		GRP=$(stat -f "%g" $PROJ)
		;;
esac
