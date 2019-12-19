# Redis Monitor Tool
this redis monitor tool only support <b>Open-Falcon</b> Agent，so we must need Open-Falcon in out environment.

## Requirement
- Develop Operation System：MacOS / Linux
- Program Language：Golang
- Dependency Package
  - gopkg.in/yaml.v2
  - gopkg.in/redis.v4

## Collect Matrics
Matrics collection setting is in ```conf/monitor.yaml```, we just using "info section" to collect system matrics.

## Deployment
### VM
1. Use make file
````
make build
make pack
````
2. move ```redis-monitor-v1.0.0.tar.gz``` file to  ```/monitor```
````
mv redis-monitor-v1.0.0.tar.gz /monitor/
````
3. decompression the gz file 
````
cd /monitor
tar xvf redis-monitor-v1.0.0.tar.gz
````
4. edit ```/monitor/conf/monitor.yaml```
5. add the record to crontab
````
#crontab -e
*/1 * * * * cd /monitor && ./redis-monitor >> /monitor/log
````

### Docker
1. just docker build
````
docker build -t redis-monitor:1.0.0 .
````

2. start app using docker run
````
docker run --name redis-monitor -d \
-e Agent_Url="127.0.0.1:1988" \
-e Redis_Endpoint="Redis Server" \
-e Redis_Port=6379 \
-e Redis_Password="" \
redis-monitor:1.0.0
````