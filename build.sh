#!/bin/bash
go build -o ../../../../bin/confer

echo "test confer"

# cd the test dir
cd ./test
echo "#####################"
echo "#     Confer TEST   #"
echo "#####################"
echo ""
echo "> confer match peerconfig.json nodes address 10.105.45.253 |  jq ."
confer match config.json nodes address 10.105.45.253 |  jq .

echo "> confer match peerconfig.json nodes id 10.105.45.253"
confer match config.json nodes id 10.105.45.253 

echo "> confer match peerconfig.json nodes id 1 |  jq ."
confer match config.json nodes id 1 | jq .

echo "> confer match peerconfig.json self id 1"
confer match config.json self id 1
