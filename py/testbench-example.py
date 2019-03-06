#!/usr/bin/env python

# Copyright ©2019 The go-lsst Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

import testbench
cli = testbench.Client(addr="0.0.0.0:5555", user="faux-fcs", pwd="faux-fcs")
print("infos: {}".format(cli.infos()))

while not (cli.x.is_ready() and cli.z.is_ready()):
    print("not ready...")
    import time
    time.sleep(1)

cli.x.set_angle(32)
cli.x.get_angle()

print("infos: {}".format(cli.infos()))

print("sleeping...")
cli.x.sleep(5)

print("x-rpm= {}".format(cli.x.rpm))
cli.x.rpm -= 20
print("x-rpm= {}".format(cli.x.rpm))

print("x-online: {}\nz-online: {}".format(cli.x.is_online(),cli.z.is_online()))
print("x-mode: {}\nz-mode: {}".format(cli.x.mode,cli.z.mode))

for motor in [cli.x, cli.z]:
    print("motor {} -- angle: {}".format(motor.name, motor.angle))
    motor.angle += 20
    motor.sleep(2)
    print("motor {} -- angle: {}".format(motor.name, motor.angle))
    print("motor {} -- RPMs:  {}".format(motor.name, motor.rpm))
    motor.rpm -= 100
    print("motor {} -- RPMs:  {}".format(motor.name, motor.rpm))
    print("motor {} -- temperatures: {}".format(motor.name, motor.temps()))
    pass
