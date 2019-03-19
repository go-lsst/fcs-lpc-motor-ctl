#!/usr/bin/env python
# coding=utf-8

# Copyright ©2019 The go-lsst Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

import time
import testbench

cli = testbench.Client(
        addr="134.158.155.17:5454", 
        user="faux-fcs", pwd="faux-fcs",
        timeout=30, 
        verbose=False)
print("infos: {}".format(cli.infos()))

while not (cli.x.is_ready()):
    print("not ready...")
    time.sleep(1)

print("ready.")
print("sending x-angle-pos -15...")
cli.x.set_angle(-15)
cli.x.rpm = 1300
cli.x.wait()
print("x-angle-pos = {}".format(cli.x.get_angle()))

print("sending z-angle-pos -15...")
cli.z.set_angle(-15)
cli.z.rpm = 1300
cli.z.wait()
print("z-angle-pos = {}".format(cli.z.get_angle()))

print("infos: {}".format(cli.infos()))

print("sleeping...")
cli.x.sleep(5)
print("sleeping... [done]")

print("x-rpm= {}".format(cli.x.rpm))
cli.x.rpm -= 20
print("x-rpm= {}".format(cli.x.rpm))

#print("x-online: {}\nz-online: {}".format(cli.x.is_online(),cli.z.is_online()))
#print("x-mode: {}\nz-mode: {}".format(cli.x.mode,cli.z.mode))
print("x-online: {}".format(cli.x.is_online()))
print("x-mode: {}".format(cli.x.mode))

for motor in [cli.x, cli.z]:
    print("motor {} -- angle: {}".format(motor.name, motor.angle))
    motor.angle += 10
    motor.wait()
    print("motor {} -- angle: {}".format(motor.name, motor.angle))
    print("motor {} -- RPMs:  {}".format(motor.name, motor.rpm))
    motor.rpm -= 100
    print("motor {} -- RPMs:  {}".format(motor.name, motor.rpm))
    print("motor {} -- temperatures: {}".format(motor.name, motor.temps()))
    pass

print("*"*80)
while 1:
    print("-"*80)
    print("motor {} -- angle: {}".format(cli.x.name, cli.x.angle))
    print("motor {} -- angle: {}".format(cli.z.name, cli.z.angle))
    cli.x.angle = 22
    cli.z.angle = 22
    cli.wait()
    print("motor {} -- angle: {}".format(cli.x.name, cli.x.angle))
    print("motor {} -- angle: {}".format(cli.z.name, cli.z.angle))
    cli.x.angle -= 20
    cli.z.angle -= 20
    cli.wait()
    pass
