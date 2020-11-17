#!/usr/bin/env bash
git pull
./devMode.sh prod

cd docker
./pu.sh
cd ..
./devMode.sh dev