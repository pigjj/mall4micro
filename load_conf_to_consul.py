# !/usr/bin/env python
# -*- coding: utf-8 -*-

# @Time: 2021/12/23 6:27 PM
# @Author: Haibo Jiang
# @Email: haibojiang@bitorobotics.ltd
# @File: load_conf_to_consul.py
# @Version: V0.0.1
# @license: (C) Copyright 2017-2030, Bito Robotics Co.Ltd.
# @desc:
import os

import requests

YAML_CONF_PATH = os.path.join(os.path.dirname(__file__), "yaml_conf")

CONSUL_URL = "http://127.0.0.1:8500/v1/kv"
TOKEN = ""

HEADER = {"Authorization": TOKEN, "Content-Type": "application.json"}


def upload_config_2_consul(file_name):
    with open(os.path.join(YAML_CONF_PATH, file_name)) as f:
        conf_lines = f.readlines()
    response = requests.put(
        "{}/{}".format(CONSUL_URL, file_name), headers=HEADER, data="".join(conf_lines))
    return response.status_code == 200


def run():
    file_list = os.listdir(YAML_CONF_PATH)
    for file_name in file_list:
        success = upload_config_2_consul(file_name)
        if success:
            print("[Load Config] file: {} success!".format(file_name))
        else:
            print("[Load Config] file: {} failed!".format(file_name))


if __name__ == '__main__':
    run()
