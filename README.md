fcs-lpc-motor-ctl
=================

A `html5` application using the `github.com/go-lsst/ncs` control system to control the `m702` motors.

## Installation

```sh
sh> go get github.com/go-lsst/fcs-lpc-motor-ctl
sh> fcs-lpc-motor-ctl -addr=:7777 &
sh> open 0.0.0.0:7777
```

## API

```sh
$> curl -u faux-fcs:faux-fcs clrbinetsrv.in2p3.fr:5555/api/mon
{"error":"","code":200,"infos":[{"motor":"x","online":true,"status":"ready","mode":"home","rpms":1300,"angle":0,"temps":[36,41,45,51],"histos":null,"webcam":""},{"motor":"z","online":true,"status":"ready","mode":"home","rpms":1300,"angle":0,"temps":[36,40,46,51],"histos":null,"webcam":""}]}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-pos --data '{"motor":"x", "value":1}'
{error:"", code:200}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-angle-pos --data '{"motor":"x", "value":45}'
{error:"", code:200}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-pos --data '{"motor":"z", "value":1}'
{error:"", code:200}

$> curl -u faux-fcs:faux-fcs clrbinetsrv.in2p3.fr:5555/api/mon
{"error":"","code":200,"infos":[{"motor":"x","online":true,"status":"ready","mode":"pos","rpms":1300,"angle":45,"temps":[37,41,46,50],"histos":null,"webcam":""},{"motor":"z","online":true,"status":"ready","mode":"pos","rpms":1300,"angle":0,"temps":[36,40,47,51],"histos":null,"webcam":""}]}

```
