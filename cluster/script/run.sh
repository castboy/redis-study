SERVER="/opt/homebrew/opt/redis/bin/redis-server"
SENTIENL="/opt/homebrew/opt/redis/bin/redis-sentinel"

PID_DIR="/usr/local/redis/var"
PID_MASTER="$PID_DIR/redis-master.PID"
PID_SLAVE1="$PID_DIR/redis-slave1.PID"
PID_SLAVE2="$PID_DIR/redis-slave2.PID"
PID_SENTINEL1="$PID_DIR/redis-sentinel1.PID"
PID_SENTINEL2="$PID_DIR/redis-sentinel2.PID"
PID_SENTINEL3="$PID_DIR/redis-sentinel3.PID"

LOG_DIR="/usr/local/redis/var"
LOG_MASTER="$LOG_DIR/redis-master.log"
LOG_SLAVE1="$LOG_DIR/redis-slave1.log"
LOG_SLAVE2="$LOG_DIR/redis-slave2.log"
LOG_SENTINEL1="$LOG_DIR/redis-sentinel1.log"
LOG_SENTINEL2="$LOG_DIR/redis-sentinel2.log"
LOG_SENTINEL3="$LOG_DIR/redis-sentinel3.log"

DATA_DIR="/usr/local/redis/data"
DATA_MASTER="$DATA_DIR/master"
DATA_SLAVE1="$DATA_DIR/slave1"
DATA_SLAVE2="$DATA_DIR/slave2"
DATA_SENTINEL1="$DATA_DIR/redis-sentinel1"
DATA_SENTINEL2="$DATA_DIR/redis-sentinel2"
DATA_SENTINEL3="$DATA_DIR/redis-sentinel3"

PORT_MASTER=6380
PORT_SLAVE1=6381
PORT_SLAVE2=6382
PORT_SENTINEL1=26380
PORT_SENTINEL2=26381
PORT_SENTINEL3=26382

start() {
  if [ ! -d $1 ]; then
      mkdir -p $1
  fi

  if [ ! -f $2 ]; then
      touch $2
  fi

  if [ ! -d $3 ]; then
      mkdir -p $3
  fi

  if [ ! -f $4 ]; then
      touch $4
  fi

  if [ ! -d $5 ]; then
      mkdir -p $5
  fi

  $7 $6
}

kill() {
  ps aux | grep -w $1 | grep -v grep | awk '{print $2}' | xargs kill -9
}

case "$1" in
start_master)
  echo "start redis-master ..."
  start $PID_DIR $PID_MASTER $LOG_DIR $LOG_MASTER $DATA_MASTER ./master/redis.master.conf $SERVER
  echo "start redis-master end"
  ;;
start_slave1)
  echo "start redis-slave1 ..."
  start $PID_DIR $PID_SLAVE1 $LOG_DIR $LOG_SLAVE1 $DATA_SLAVE1 ./slave/redis.slave1.conf $SERVER
  echo "start redis-slave1 end"
  ;;
start_slave2)
  echo "start redis-slave2 ..."
  start $PID_DIR $PID_SLAVE2 $LOG_DIR $LOG_SLAVE2 $DATA_SLAVE2 ./slave/redis.slave2.conf $SERVER
  echo "start redis-slave2 end"
  ;;
start_sentinel1)
  echo "start redis-sentinel1 ..."
  start $PID_DIR $PID_SENTINEL1 $LOG_DIR $LOG_SENTINEL1 $DATA_SENTINEL1 ./sentinel/sentinel1.conf $SENTIENL
  echo "start redis-sentinel1 end"
  ;;
start_sentinel2)
  echo "start redis-sentinel2 ..."
  start $PID_DIR $PID_SENTINEL2 $LOG_DIR $LOG_SENTINEL2 $DATA_SENTINEL2 ./sentinel/sentinel2.conf $SENTIENL
  echo "start redis-sentinel2 end"
  ;;
start_sentinel3)
  echo "start redis-sentinel3 ..."
  start $PID_DIR $PID_SENTINEL3 $LOG_DIR $LOG_SENTINEL3 $DATA_SENTINEL3 ./sentinel/sentinel3.conf $SENTIENL
  echo "start redis-sentinel3 end"
  ;;
kill_master)
  echo "kill master ..."
  kill $PORT_MASTER
  echo "kill master down"
  ;;
kill_slave1)
  echo "kill slave1 ..."
  kill $PORT_SLAVE1
  echo "kill slave1 end"
  ;;
kill_slave2)
  echo "kill slave2 ..."
  kill $PORT_SLAVE2
  echo "kill slave2 end"
  ;;
kill_sentinel1)
  echo "kill sentinel1 ..."
  kill $PORT_SENTINEL1
  echo "kill sentinel1 end"
  ;;
kill_sentinel2)
  echo "kill sentinel2 ..."
  kill $PORT_SENTINEL2
  echo "kill sentinel2 end"
  ;;
kill_sentinel3)
  echo "kill sentinel3 ..."
  kill $PORT_SENTINEL3
  echo "kill sentinel3 end"
  ;;
esac