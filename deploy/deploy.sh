#!/bin/sh
CURRENT_DIR=$(cd $(dirname $0); pwd)
cd $CURRENT_DIR

chmod +x ../backstage/build.sh
../backstage/build.sh

chmod +x ../minigame/build.sh
../minigame/build.sh

export COMPOSE_PROJECT_NAME=game
docker-compose -f docker_compose.yml up -d

rm ../backstage/backstage
rm ../minigame/minigame