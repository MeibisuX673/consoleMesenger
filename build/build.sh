#!/bin/bash

os=$2
architecture=$3

echo $os

echo -e "\n Введите версию:"
read version

echo -e "\n Введите os(optional):"
read os

echo -e "\n Введите архитектуру(optional):"
read architecture

if [ "$version" == "" ]
then
  echo "Нужно указать версию"
  exit 1
fi

outputClient="client-$os-$architecture"
outputServer="server-$os-$architecture"

user=$(whoami)
date=`date`

if [[ "$architecture" != "" && "$os" != "" ]]
then
  GOOS=$s GOARCH=$architecture go build -o bin/$outputClient -v -ldflags="-X 'main.Version=$version' -X 'main.User=$user' -X 'main.Date=$date'" ./tcp-client/tcp-client.go
  GOOS=$os GOARCH=$architecture go build -o bin/$outputServer -v -ldflags="-X 'main.Version=$version' -X 'main.User=$user' -X 'main.Date=$date'" ./tcp-server/tcp-server.go
else
  go build -o bin/tcp-client-bin -v -ldflags="-X 'main.Version=$version' -X 'main.User=$user' -X 'main.Date=$date'" ./tcp-client/tcp-client.go
  go build -o bin/tcp-server-bin -v -ldflags="-X 'main.Version=$version' -X 'main.User=$user' -X 'main.Date=$date'" ./tcp-server/tcp-server.go
fi



