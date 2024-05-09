#!/bin/sh

wget https://github.com/ppdx999/migr/archive/main.zip
unzip main.zip
cp -r migr-main/migr-* .
rm -rf migr-main main.zip
chmod +x migr-*
