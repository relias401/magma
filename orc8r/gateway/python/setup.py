"""
Copyright (c) 2016-present, Facebook, Inc.
All rights reserved.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree. An additional grant
of patent rights can be found in the PATENTS file in the same directory.
"""

import os

from setuptools import setup

# We can use an environment variable to pass in the package version during
# build. Since we don't distribute this on its own, we don't really care what
# version this represents. 'None' defaults to 0.0.0.
VERSION = os.environ.get('PKG_VERSION', None)

setup(
    name='orc8r',
    version=VERSION,
    packages=[
        'magma.common',
        'magma.common.health',
        'magma.common.redis',
        'magma.configuration',
        'magma.magmad',
        'magma.magmad.generic_command',
        'magma.magmad.check',
        'magma.magmad.check.kernel_check',
        'magma.magmad.check.machine_check',
        'magma.magmad.check.network_check',
        'magma.magmad.logging',
        'magma.magmad.upgrade',
    ],
    scripts=[
        'scripts/checkin_cli.py',
        'scripts/directoryd_cli.py',
        'scripts/generate_dnsd_config.py',
        'scripts/generate_lighttpd_config.py',
        'scripts/generate_nghttpx_config.py',
        'scripts/generate_service_config.py',
        'scripts/health_cli.py',
        'scripts/magma_conditional_service.py',
        'scripts/magma_get_config.py',
        'scripts/magmad_cli.py',
        'scripts/service303_cli.py',
        'scripts/show_gateway_info.py',
    ],
    install_requires=[
        'Cython>=0.29.1',
        'pystemd==0.5.0',
        'docker>=4.0.2',
        # Waiting for a new fire release (v0.1.3 is too old)
        'fire @ git+https://www.github.com/google/python-fire@'
        '024fbad9424cfdb0d3c93c86c856aedbac0f9d48',
        'aioh2==0.2.2',
        'redis>=2.10.5',  # redis-py (Python bindings to redis)
        'redis-collections>=0.4.2',
        'aiohttp>=0.17.2',
        'grpcio==1.16.1',
        'protobuf==3.6.1',
        'Jinja2>=2.8',
        'netifaces>=0.10.4',
        'pylint>=1.7.1',
        'PyYAML>=3.12',
        'pytz>=2014.4',
        'prometheus_client==0.3.1',
        'snowflake>=0.0.3',
        'psutil>=5.2.2',
        'cryptography>=1.9',
        'systemd-python>=234',
        'itsdangerous>=0.24',
        'click>=5.1',
        'pycares>=2.3.0',
        'python-dateutil>=1.4',
    ]
)
