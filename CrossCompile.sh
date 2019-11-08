#!/bin/bash
OUTDIR=builds

PREFIX32=GOARCH=386
PREFIX64=GOARCH=amd64
INDEXTRACKER_OUTNAME=indextracker
TRACKERMANAGER_OUTNAME=./cmd/trackermanager


GOOS="linux darwin windows freebsd netbsd openbsd dragonfly plan9 nacl android"

PREFIX=GOPATH=$(pwd)
COMPILER=go

function err(){
	echo $1
	exit 1;
}
function compile(){
#	if ! test -d $OUTDIR; then
#		mkdir $OUTDIR || err "Cant create $OUTDIR"
#	fi;
	export $PREFIX
	for g in $GOOS; do
		echo "comple for $g"

		export GOOS=$g 
		export $PREFIX32 
		$COMPILER build -o ${INDEXTRACKER_OUTNAME}${GOOS}32

		export GOOS=$g 
		export $PREFIX64
	       	$COMPILER build -o ${INDEXTRACKER_OUTNAME}${GOOS}64




	done;
}
compile


