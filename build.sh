#!/bin/bash

# go to client directory
cd client
npm run build

# remove old files
rm -rf ../public/static/

# move the new build file to public directory
mv build/* ../public/