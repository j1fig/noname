#!/bin/bash
docker-machine create --driver digitalocean --digitalocean-access-token $DROPLET_ACCESS_TOKEN --digitalocean-size $DROPLET_SIZE --digitalocean-region $DROPLET_REGION $DROPLET_NAME
eval $(docker-machine env $DROPLET_NAME)
