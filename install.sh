#!/bin/bash

function install {
  echo -n "installing $1... "
  cd $1
  go install
  cd ..
  echo "done"
}

install csvtojson
