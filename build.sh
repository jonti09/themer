#!/bin/bash

# go to client directory
cd client
npm run build

# create tar of build files
cd build
tar -cvzf ../public.tar.gz *
rm -rf ../build