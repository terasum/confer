#!/bin/bash
set -e
go build -o ../../../../bin/confer

echo "test confer"

# cd the test dir
cd ./test
echo "#####################"
echo "#     Confer TEST   #"
echo "#####################"
echo "--------------------"
echo " confer read"
echo ">>>"
echo "> confer read config.json nodes | jq ."
confer read config.json nodes | jq .

echo "> confer read config.json self.node_id"
confer read config.json self.node_id

echo "> confer read config.json self.is_origin"
confer read config.json self.is_origin

echo "> confer read config.json self | jq ."
confer read config.json self | jq .

echo "--------------------"
echo " confer write"
echo ">>>"

echo "> confer write config.json config_test.json self.is_origin false -t bool -y"
confer write config.json config_test.json self.is_origin false -t bool -y
echo "> cat config_test.json | jq ."
cat config_test.json | jq .
echo "> rm -f config_test.json"
rm -f config_test.json

echo "> confer write config.json config_test.json self.node_id 2 -t int -y"
confer write config.json config_test.json self.node_id 2 -t int -y
echo "> cat config_test.json | jq ."
cat config_test.json | jq .
echo "> rm -f config_test.json"
rm -f config_test.json

echo "--------------------"
echo " confer writeb"
echo ">>>"

echo "> confer writeb config.json self.is_origin false -t bool | jq ."
confer writeb config.json self.is_origin false -t bool | jq .


echo "> confer writeb config.json self.node_id 2 -t int | jq ."
confer writeb config.json self.node_id 2 -t int | jq .


echo "--------------------"
echo " confer match"
echo ">>>"
echo "> confer match config.json nodes address 10.105.45.253 |  jq ."
confer match config.json nodes address 10.105.45.253 |  jq .

echo "> confer match config.json nodes id 10.105.45.253"
confer match config.json nodes id 10.105.45.253 

echo "> confer match config.json nodes id 1 |  jq ."
confer match config.json nodes id 1 | jq .

echo "> confer match config.json self id 1"
confer match config.json self id 1
