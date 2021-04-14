sudo docker ps -a|grep telegraf| awk '{print $1}' |xargs sudo docker stop
sudo docker ps -a|grep telegraf| awk '{print $1}' |xargs sudo docker rm
sudo docker images |grep telegraf| awk '{print $3}' |xargs sudo docker image rm
go build  cmd/telegraf/telegraf.go
sudo docker build -t telegraf .
sudo docker run -p 8095:8095 --network=bridge --link kafka --link influxsrv -d telegraf:latest
