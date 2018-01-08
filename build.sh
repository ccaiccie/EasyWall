#!/bin/bash

SCRIPTPATH=$(pwd)
SRCPATH=$SCRIPTPATH/src
GOPATH=/usr/local/go/bin/go
GRUNTPATH=$SRCPATH/node_modules/grunt/bin/grunt
CONFIGPATH=$SCRIPTPATH/config

rm -rf $CONFIGPATH
cd $SRCPATH
$GOPATH build
$GRUNTPATH
mv src easywall
chmod +x easywall
mv easywall $SCRIPTPATH
cd $SCRIPTPATH
./easywall

