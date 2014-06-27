#!/bin/bash
prevGoPath=$GOPATH
prevGoRoot=$GOROOT
git submodule update --init
cd heka
cp ../plugin_loader.cmake cmake/plugin_loader.cmake
. build.sh
go get github.com/go-sql-driver/mysql
cd ../../ 
cd heka/build
make
cd ../../
export GOPATH=$prevGoPath
export GOROOT=$prevGoRoot
echo "DONE." 
