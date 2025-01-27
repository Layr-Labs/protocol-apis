#!/usr/bin/env bash

base_project_dir="gen/python/eigenlabs-protocol-apis"
project_src_dir="${base_project_dir}/src"
module_name="eigenlabs-protocol-apis"
version=${VERSION:-"0.1.0"}

mkdir -p $base_project_dir || true

cat > "${base_project_dir}/pyproject.toml" << EOL
[build-system]
requires = ["setuptools>=61.0"]
build-backend = "setuptools.build_meta"

[project]
name = "${module_name}"
version = "${version}"
description = "Eigenlabs Protocol APIs grpc clients"
requires-python = ">=3.7"
dependencies = [
    "grpcio>=1.32.0",
    "protobuf==4.25.5",
]
EOL

cat > "${base_project_dir}/setup.py" << EOL
from setuptools import setup, find_namespace_packages

setup(
    name="${module_name}",
    version="${version}",
    packages=find_namespace_packages(include=['eigenlabs.*']),
    package_dir={'': 'src'},
    install_requires=[
        "grpcio>=1.32.0",
        "protobuf==4.25.5",
    ],
)
EOL

# Move generated files to src directory
mv gen/pre-python "${base_project_dir}/src/"

find $base_project_dir -type d -exec touch {}/__init__.py \;
