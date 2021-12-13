#!/bin/bash

cat <&0 > ../../scripts/sample/all.yaml

../../scripts/sample/kustomize build ../../scripts/sample
