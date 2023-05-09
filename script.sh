#!/bin/sh

echo $1

 if [ "$#" -eq  "0" ]
   then
     echo "No arguments supplied"
     a="default_name"
 else
     echo "Hello world"
     a=$1
 fi

echo "$a"
mkdir "$a"
cd "$a"
mkdir "greetings"
cd greetings
touch greetings.go
go mod init example.com/greetings
git add go.mod
git add greetings.go

cd ..
mkdir hello
cd hello
touch hello.go
go mod init example.com/hello

echo "replace example.com/greetings => ../greetings" >> go.mod
go mod tidy
git add -A ./