#!/bin/bash
docker-machine stop $DROPLET_NAME
docker-machine rm -y $DROPLET_NAME
