#!/bin/sh

# set current working directory to swapchain/contrib
cd `dirname $0`

# cleanup
rm -rf osmosis

# clone icq enabled osmosis repo
git clone --branch=v15.0.0 https://github.com/osmosis-labs/osmosis.git
cd osmosis

# copy ignite config for starting node
cp ../osmosis-config.yml ./config.yml

# build and start node
ignite chain serve --reset-once --skip-proto
