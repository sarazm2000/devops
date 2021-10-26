#!/bin/bash

docker-compose up -d

bash setip.sh

ping -c 4 web.hw1
