#!/bin/bash -e

# navigate to core repo
cd $GOPATH/src/github.com/mesg-foundation/core

# container package
mockery -name=Container -dir ./container -output ./container/mocks

# database package
mockery -name=ServiceDB -dir ./database -output ./database/mocks