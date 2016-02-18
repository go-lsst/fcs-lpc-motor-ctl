fcs-lpc-motor-ctl
=================

A `html5` application using the `github.com/go-lsst/ncs` control system to control the `m702` motors.

## Installation

```sh
sh> go get github.com/go-lsst/fcs-lpc-motor-ctl
sh> fcs-lpc-motor-ctl -addr=:7777 &
sh> open 0.0.0.0:7777
```
