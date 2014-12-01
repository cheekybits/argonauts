#!/bin/bash
go build -o csvtojson
cat test/test.csv | ./csvtojson
rm ./csvtojson