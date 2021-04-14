FROM nimmis/golang
RUN apt-get update
RUN mkdir /telegraf
RUN mkdir /telegraf/log

COPY ca.key /telegraf/
COPY ca.pem /telegraf/
COPY telegraf /telegraf/
COPY telegraf.conf /telegraf/
WORKDIR /telegraf
CMD ./telegraf  -v=0 --config ./telegraf.conf -log_dir=./log