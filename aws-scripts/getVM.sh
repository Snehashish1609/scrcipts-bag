#!/bin/bash

set -e

INSTANCE_ID=$1

# Get the command output
output=$(aws ec2 describe-instances --instance-ids $INSTANCE_ID)

# echo $output
# Convert the command output to JSON
json=$(jq -c "$output")

# Retrieve the field value
name=$(jq -c '.name' <<< "$json")

# Print the field value
echo "$name"