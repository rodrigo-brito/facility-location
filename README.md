# Facility Location Problem

[![Build Status](https://travis-ci.org/rodrigo-brito/facility-location.svg?branch=master)](https://travis-ci.org/rodrigo-brito/facility-location) [![Go Report Card](https://goreportcard.com/badge/github.com/rodrigo-brito/facility-location)](https://goreportcard.com/report/github.com/rodrigo-brito/facility-location)

Facility location problem with single allocation solved in Go (Golang).

## Requiriments

Install Go - https://golang.org/doc/install

## Usage

```
go build
```

```
./facility-location [input-file] [flags]

Flags:
  -a, --async int    number of async tasks (default 4)
  -b, --best float   value of the best solution
  -h, --help         help for hub-spoke-go
  -v, --verbose      active verbose mode
```

Example:

```
./facility-location ./facility-location data/ap10_2.txt -b 90963539.4763 -v -a 4
```

## Heuristics applied

- GRASP (Greedy Randomized Adaptive Search Procedures)
- VND (Variable Neighborhood Descent Search)
- VNS (Variable Neighborhood Search)

## Dataset

- Australian Post (10 to 100 nodes) - [Andreas T.ErnstMohanKrishnamoorthy](https://www.sciencedirect.com/science/article/pii/S0966834996000113)

## Model

<img src="model.png" />
