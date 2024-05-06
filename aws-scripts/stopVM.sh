#!/bin/bash

set -e

INSTANCE_ID=$1

aws ec2 stop-instances --instance-ids $INSTANCE_ID