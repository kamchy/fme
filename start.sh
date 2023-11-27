#!/bin/bash
docker run -e FMEDIR=/root/data -it --rm -v "$PWD":/root/data kamchyd/fme:v1.0.1
