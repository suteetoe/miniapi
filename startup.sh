#! /bin/bash

## run filebeat agent
filebeat -c /usr/share/filebeat/filebeat.yml

## run app
/root/go-app