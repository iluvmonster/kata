#!/bin/bash -xe

docker build -t kata .
docker run -ti kata bash
