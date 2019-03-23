#!/usr/bin/env python
# coding=utf-8

# Copyright ©2019 The go-lsst Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

from __future__ import division, print_function, unicode_literals

import json
import base64
try:
    import httplib
except ImportError:
    import http.client as httplib

from time import sleep


class Client(object):
    def __init__(self, user, pwd, addr, timeout=10, verbose=False):
        self.usr = user
        self.pwd = pwd
        self.addr = addr
        self.verbose = verbose

        self.hdlr = httplib.HTTPConnection(addr, timeout=timeout)

        auth_str = "{}:{}".format(user, pwd)
        try:
            # Python 2
            auth = base64.encodestring(auth_str).strip()
        except TypeError:
            # Python 3
            auth = base64.encodebytes(auth_str.encode())
            auth = auth.decode().strip()

        self.hdr = {
            "Content-Type": "application/json",
            "Authorization": "Basic " + auth,
        }

        self.x = Motor(self, "x")
        self.z = Motor(self, "z")

        print("connecting to {}...".format(addr))
        self._connect(timeout)
        print("connecting to {}... [done]".format(addr))
    
    def _request_info(self):
        self.hdlr.request("GET", "/api/mon", None, self.hdr)
        resp = self.hdlr.getresponse()
        response = resp.read()
        if self.verbose:
            print("response: {!r}".format(response))

        resp_dict = json.loads(response)

        if resp.status != httplib.OK:
            raise RuntimeError(
                "invalid status: {} -- error: {}".format(resp.status, resp_dict["error"])
            )
        return resp_dict

    def _request_action(self, method, action, data):
        if isinstance(data, dict):
            data = json.dumps(data)
        if self.verbose:
            print("request-data: {!r}".format(data))
            print("request-cmd: {!r}".format(action))

        self.hdlr.request(method, "/api/cmd/{}".format(action), data, self.hdr)
        resp = self.hdlr.getresponse()

        response = resp.read()
        resp_dict = json.loads(response)

        if self.verbose:
            print("response: {!r}".format(response))

        if resp.status != httplib.OK:
            raise RuntimeError(
                "invalid status: {} -- error: {}".format(resp.status, resp_dict["error"])
            )

        return resp_dict

    def _connect(self, timeout):
        for _ in range(timeout):
            try:
                self._request_info()
                break
            except RuntimeError:
                sleep(1)
                continue
        else:
            raise RuntimeError(
                "could not establish connection to {} after {}s".format(self.addr, timeout)
            )

    def infos(self):
        resp_dict = self._request_info()
        
        return resp_dict["infos"]

    def is_online(self):
        infos = self.infos()
        for info in infos:
            if not info["online"]:
                return False
        return True

    def wait(self):
        self.x.wait()
        self.z.wait()
    
    def reset(self):
        self.x.reset()
        self.z.reset()

    def stop(self):
        self.x.stop()
        self.z.stop()


class Motor(object):
    def __init__(self, cli, name):
        self.cli = cli
        self.name = name

    def _run(self, cmd):
        resp_dict = self.cli._request_action(
            method="POST",
            action="req-upload-cmds",
            data={"motor": self.name, "cmds": cmd},
        )
        result = resp_dict["script"].split("=")
        
        if len(result) != 2:
            return
        
        return result[-1].strip()
    
    def _get(self, cmd):
        return self.cli._request_action(
            method="GET",
            action=cmd,
            data={"motor": self.name},
        )

    def sleep(self, secs):
        """sleep puts the motor in sleep mode for the provided amount
        of seconds.
        """
        self._run(
            "sleep {}s".format(secs)
        )

    def set_angle(self, pos):
        """
        set_angle takes a floating point value indicating the new
        angle position the testbench should go to.
        """
        self._run(
            "{}-angle-pos {:d}".format(self.name, pos)
        )
        self.wait()

    def get_angle(self):
        """
        get_angle returns the floating point value of the testbench
        position in degrees.
        """
        res_dict = self._get("req-get-angle-pos")
        return float(res_dict["value"])

    angle = property(get_angle, set_angle)

    def set_mode(self, mode):
        want = ("find-home", "pos")
        if not mode in want:
            raise RuntimeError("invalid mode {} (want: {})".format(mode, want))

        self._run(
            "{}-{}".format(self.name, mode)
        )

    def get_mode(self):
        return self.infos()["mode"]

    mode = property(get_mode, set_mode)

    def set_rpm(self, rpm):
        self._run(
            "{}-rpm {:d}".format(self.name, rpm)
        )

    def get_rpm(self):
        res_dict = self._get("req-get-rpm")
        return int(res_dict["value"])

    rpm = property(get_rpm, set_rpm)

    def reset(self):
        self.cli._request_action(
            method="POST", 
            action="req-reset", 
            data={"motor": self.name},
        )
        
    def stop(self):
        self.cli._request_action(
            method="POST", 
            action="req-stop",
            data={"motor": self.name},
        )

    def wait(self):
        self.cli._request_action(
            method="GET", 
            action="req-wait-pos",
            data={"motor": self.name},
        )

    def infos(self):
        for info in self.cli.infos():
            if info["motor"] == self.name:
                return info

        raise RuntimeError("No infos for motor {}".format(self.name))

    def is_online(self):
        return self.infos()["online"]

    def is_ready(self):
        ready = self.status() == "ready"
        return ready and self.is_sync()

    def status(self):
        return self.infos()["status"]

    def temps(self):
        return self.infos()["temps"]

    def is_sync(self):
        return self.infos()["sync"]
