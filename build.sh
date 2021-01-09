#!/bin/bash

cd client
npm run build

mv build ../public

go run main.go