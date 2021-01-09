#!/bin/bash

if [ $1 == 'dev' ]
then
    go run main.go
else
    cd client
    npm run build

    cp -r build/* ../public/
    rm -rf build

    cd ..
fi
