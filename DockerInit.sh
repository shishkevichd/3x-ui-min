#!/bin/sh
mkdir -p build/bin
cd build/bin
wget -q "https://github.com/XTLS/Xray-core/releases/download/v25.6.8/Xray-linux-64.zip"
unzip "Xray-linux-64.zip"
rm -f "Xray-linux-64.zip" geoip.dat geosite.dat
mv xray "xray-linux-amd64"
wget -q https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geoip.dat
wget -q https://github.com/Loyalsoldier/v2ray-rules-dat/releases/latest/download/geosite.dat
cd ../../