PROJECT="redis cluster"

default:
	echo ${PROJECT}

start_all:
	./script/run.sh start_master
	./script/run.sh start_slave1
	./script/run.sh start_slave2
	./script/run.sh start_sentinel1
	./script/run.sh start_sentinel2
	./script/run.sh start_sentinel3

start_master:
	./script/run.sh start_master

start_slave1:
	./script/run.sh start_slave1

start_sentinel1:
	./script/run.sh start_sentinel1

start_sentinel2:
	./script/run.sh start_sentinel2

start_sentinel3:
	./script/run.sh start_sentinel3

start_slave2:
	./script/run.sh start_slave2

restart_all:
	./script/run.sh kill_master
	./script/run.sh kill_slave1
	./script/run.sh kill_slave2
	./script/run.sh kill_sentinel1
	./script/run.sh kill_sentinel2
	./script/run.sh kill_sentinel3

	./script/run.sh start_master
	./script/run.sh start_slave1
	./script/run.sh start_slave2
	./script/run.sh start_sentinel1
	./script/run.sh start_sentinel2
	./script/run.sh start_sentinel3

restart_master:
	./script/run.sh kill_master
	./script/run.sh start_master

restart_slave1:
	./script/run.sh kill_slave1
	./script/run.sh start_slave1

restart_slave2:
	./script/run.sh kill_slave2
	./script/run.sh start_slave2

restart_sentinel1:
	./script/run.sh kill_sentinel1
	./script/run.sh start_sentinel1

restart_sentinel2:
	./script/run.sh kill_sentinel2
	./script/run.sh start_sentinel2

restart_sentinel3:
	./script/run.sh kill_sentinel3
	./script/run.sh start_sentinel3