#!/bin/bash
OUTPUT=$(echo -e "protocol=https\nhost=github.com\n" | git credential fill 2> /dev/null)

USERNAME=$(echo $OUTPUT | grep -o -E "username=([^ ]+)" | cut -d '=' -f 2)
PASSWORD=$(echo $OUTPUT | grep -o -E "password=([^ ]+)" | cut -d '=' -f 2)

echo ${USERNAME}:${PASSWORD}