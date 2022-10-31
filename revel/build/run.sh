#!/bin/sh

SCRIPTPATH=$(cd "$(dirname "$0")"; pwd)
"$SCRIPTPATH/revel.exe" -importPath revel -srcPath "$SCRIPTPATH/src" -runMode dev
