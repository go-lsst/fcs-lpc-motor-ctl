#!/usr/bin/env python

# Copyright ©2019 The go-lsst Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

import base64
import json
import httplib

class Client(object):
    def __init__(self, user="", pwd="", addr="", timeout=10, verbose=False):
        self.usr = user
        self.pwd = pwd
        self.addr = addr
        self.verbose = verbose

        self.hdlr = httplib.HTTPConnection(addr, timeout=timeout)
        auth = base64.encodestring(user+':'+pwd).replace("\n","")
        self.hdr = {
            "Content-Type":  "application/json",
            "Authorization": "Basic "+auth,
        }

        self.x = Motor(self, "x")
        self.z = Motor(self, "z")

        print("connecting to {}...".format(addr))
        self._connect(timeout)
        print("connecting to {}... [done]".format(addr))

    def _connect(self, timeout):
        import time
        for i in range(timeout):
            self.hdlr.request("GET", "/api/mon", None, self.hdr)
            resp = self.hdlr.getresponse()
            v = resp.read()
            if self.verbose:
                print("response: %r" % (v,))
                print("response: %s" % (json.loads(v),))
            v = json.loads(v)
            if resp.status == httplib.OK:
                return
            time.sleep(1)
            pass
        raise RuntimeError("could not establish connection to {} after {}s".format(self.addr, timeout))


    def infos(self):
        self.hdlr.request("GET", "/api/mon", None, self.hdr)
        resp = self.hdlr.getresponse()
        v = resp.read()
        if self.verbose:
            print("response: %r" % (v,))
            print("response: %s" % (json.loads(v),))
        v = json.loads(v)
        if resp.status != 200:
            raise RuntimeError("invalid status: %s -- error: %s" % (resp.status,v["error"]))
        return v["infos"]

    def is_online(self):
        infos = self.infos()
        for info in infos:
            if not info["online"]:
                return False
        return True

    def _run(self, data):
        data =json.dumps(data)
        if self.verbose:
            print("request-data: '%s'" % (data,))
        self.hdlr.request("POST", "/api/cmd/req-upload-cmds", data, self.hdr)
        resp = self.hdlr.getresponse()
        v = resp.read()
        if self.verbose:
            print("response: %r" % (v,))
            print("response: %s" % (json.loads(v),))
        v = json.loads(v)
        if resp.status != 200:
            raise RuntimeError("invalid status: %s -- error: %s" % (resp.status,v["error"]))
        v = v["script"].replace("\n","")
        if not "=" in v:
            return
        i = v.find("=")
        return v[i+1:]

    pass # class Client

class Motor(object):
    def __init__(self, cli, name):
        self.cli = cli
        self.name = name

    def _run(self, cmd):
        return self.cli._run(cmd)

    def sleep(self, secs):
        """sleep puts the motor in sleep mode for the provided amount
        of seconds.
        """
        self._run({"motor":self.name, "cmds":"sleep %ss" % (secs,)})

    def set_angle(self, pos):
        """
        set_angle takes a floating point value indicating the new
        angle position the testbench should go to.
        """
        self._run({"motor":self.name, "cmds":self.name+"-angle-pos %s" % (int(pos),)})

    def get_angle(self):
        """
        get_angle returns the floating point value of the testbench
        position in degrees.
        """
        return int(self._run({"motor":self.name, "cmds":self.name+"-angle-pos"}))

    angle = property(get_angle, set_angle)

    def set_mode(self, mode):
        want = ("find-home", "pos")
        if not mode in want:
            raise RuntimeError("invalid mode %s (want: %s)" % (mode, want))
        self._run({"motor":self.name, "cmds":self.name+"-"+mode})

    def get_mode(self):
        return self.infos()["mode"]

    mode = property(get_mode, set_mode)

    def set_rpm(self, rpm):
        self._run({"motor":self.name, "cmds":self.name+"-rpm %s" % (int(rpm),)})

    def get_rpm(self):
        return int(self._run({"motor":self.name, "cmds":self.name+"-rpm"}))

    rpm = property(get_rpm, set_rpm)

    def infos(self):
        infos = self.cli.infos()
        for info in infos:
            if info["motor"] == self.name:
                return info
        raise RuntimeError("no infos for motor %s" % (self.name,))

    def is_online(self):
        return self.infos()["online"]

    def is_ready(self):
        return self.status() == "ready"

    def status(self):
        return self.infos()["status"]

    def temps(self):
        return self.infos()["temps"]

    pass # class Motor
