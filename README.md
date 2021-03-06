[![Build Status](https://api.travis-ci.org/folkengine/goname.svg?branch=master)](https://travis-ci.org/folkengine/goname)

# goname
Port of the [Random Name Generator](https://github.com/folkengine/random_name_generator) in go.

# Install

```
$> go get github.com/folkengine/goname/...
```

# Usage

```bash
$> goname -h
Usage of goname:
  -elven
    	Create Elven names
  -fantasy
    	Create Fantasy names
  -goblin
    	Create Goblin names
  -roman
    	Create Roman names (default)
```

# Testing

* [ginko](https://github.com/onsi/ginkgo)
* [gomega](https://github.com/onsi/gomega)

# Docker 

```bash
$> docker build -t goname .
$> docker run -it goname
```

# Concourse

```bash
$> fly --target psgr login --concourse-url https://concourse.example.com
$> fly set-pipeline --target psgr --config ci/pipeline.yaml --pipeline goname --non-interactive
$> fly unpause-pipeline --target psgr --pipeline goname
```

To destroy:
```bash
$> fly -t psgr destroy-pipeline --pipeline goname
```