#!/bin/bash

cd initialize/
gmct tpl --clean
gmct tpl --dir ../views
gmct static --clean
gmct static --dir ../static
cd ..

go build -ldflags "-s -w" -o gmcadmin

rm -rf admin-release
mkdir admin-release

mv gmcadmin admin-release/
cp -R conf admin-release/

rm -rf admin-release.tar.gz

tar zcfv admin-release.tar.gz admin-release

cd initialize/
gmct tpl --clean
gmct static --clean
cd ..

rm -rf admin-release