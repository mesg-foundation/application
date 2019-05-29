#!/bin/bash -e

# make sure script is running inside mesg/tools container.
source $(dirname $0)/require-mesg-tools.sh

# generate mocks for container package.
mockery -name=Container -dir ./container -output ./container/mocks

# generate mocks for docker.CommonAPIClient that used by container package.
mockery -name CommonAPIClient -dir ./internal/mocks -output ./utils/docker/mocks

# generate mocks for database package.
mockery -name=ServiceDB -dir ./database -output ./database/mocks
