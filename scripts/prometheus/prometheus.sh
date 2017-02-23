#!/bin/bash

wget -O prometheus.tar.gz https://github.com/prometheus/prometheus/releases/download/v1.5.2/prometheus-1.5.2.linux-amd64.tar.gz
mkdir /opt/prometheus
tar xvfz prometheus.tar.gz -C /opt/prometheus --strip-components=1
cd /opt/prometheus
sudo ./prometheus
