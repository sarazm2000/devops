#!/bin/bash
set -e

ip=$(docker inspect -f "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}" nginx)

cat /etc/hosts | grep -v "web.hw1" > /tmp/hosts

echo "$ip	web.hw1" >> /tmp/hosts
sudo mv /tmp/hosts /etc/hosts && sudo chmod 644 /etc/hosts
