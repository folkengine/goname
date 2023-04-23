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
