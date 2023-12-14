#!/bin/bash

for f in provider_cmd_*.go
do

    array=("provider_cmd_aws.go" "provider_cmd_azure.go" "provider_cmd_google.go")

    if ! [[ ${array[@]} =~ $f ]]
    then
        mv "$f" "$(echo "$f" | sed s/go/go.txt/)"
    fi

done