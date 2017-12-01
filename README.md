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

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"motor z\nget 0.18.002\nmotor x\nget 0.18.002"}'
{"error":"","code":200,"script":"\u003e\u003e\u003e motor z\n\u003e\u003e\u003e get 0.18.002\n\u003c\u003c\u003c Pr-00.18.002: hex=[0x00 0x00 0x00 0x00] dec=[  0   0   0   0] (0)\n\u003e\u003e\u003e motor x\n\u003e\u003e\u003e get 0.18.002\n\u003c\u003c\u003c Pr-00.18.002: hex=[0x00 0x00 0x01 0xc2] dec=[  0   0   1 194] (450)\n"}


### high-level API

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-rpm"}'
{"code":200,"script":"\u003e\u003e\u003e z-rpm\nget-z-rpm=1300\n"}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-rpm 2000"}'
{"code":200,"script":"\u003e\u003e\u003e z-rpm 2000\nset-z-rpm=2000\n"}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-rpm"}'
{"code":200,"script":"\u003e\u003e\u003e z-rpm\nget-z-rpm=2000\n"}

$> curl -u faux-fcs:faux-fcs -X POST clrbinetsrv.in2p3.fr:5555/api/cmd/req-upload-cmds --data '{"motor":"z", "cmds":"z-rpm\nx-rpm"}'
{"code":200,"script":"\u003e\u003e\u003e z-rpm\nget-z-rpm=2000\n\u003e\u003e\u003e x-rpm\nget-x-rpm=1300\n"}

```
