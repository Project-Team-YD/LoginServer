#!/bin/bash



echo "---------- Start GameServer Build ----------"
GOOS=linux GOARCH=amd64 go build -o login_server
echo "Build Success"
echo "--------------------------------------------"