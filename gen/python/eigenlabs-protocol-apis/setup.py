from setuptools import setup, find_namespace_packages

setup(
    name="eigenlabs-protocol-apis",
    version="0.1.0",
    packages=find_namespace_packages(include=['eigenlabs.*']),
    package_dir={'': 'src'},
    install_requires=[
        "grpcio>=1.32.0",
        "protobuf>=3.15.0",
    ],
)
