#!/bin/bash

SCRIPTPATH=$(pwd)
SRCPATH=$SCRIPTPATH/src
GOPATH=/usr/local/go/bin/go

cd $SRCPATH
$GOPATH build
mv src easywall
chmod +x easywall
mv easywall $SCRIPTPATH
cd $SCRIPTPATH
./easywall

