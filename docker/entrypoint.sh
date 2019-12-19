#!/bin/sh
echo "取代參數"
bash ./setParameter.sh ./config/monitor.yaml
echo "啟動監控程式"
./redis-monitor