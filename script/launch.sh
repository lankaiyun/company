#!/bin/bash

pkill company
rm -rf ./company
go build -o company main.go
nohup ./company &