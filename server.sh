#/bin/sh
# stworker 服务进程管理
# author 850630887@qq.com
# date 2018-01-21

# 定义变量
SERVICE_FLAG=stworker
SYSTEM_TIME=`date '+%Y-%m-%d %T'`
ROOT_DIR=`pwd`
SHELL_LOG_PATH="${ROOT_DIR}/logs/shell"
STWORKER_PIDFILE="${SHELL_LOG_PATH}/stworker.pid"
LOG="${SHELL_LOG_PATH}/shell.log"

# 初始化环境
if [ ! -d "$SHELL_LOG_PATH" ]; then 
	mkdir "${SHELL_LOG_PATH}"
fi
if [ ! -f "$STWORKER_PIDFILE" ]; then
	touch "${STWORKER_PIDFILE}"
fi
echo $SYSTEM_TIME >> $LOG

# start
function start() {
	pid=`cat ${STWORKER_PIDFILE}`
	count=`ps -ef |grep "${pid}" |grep ${SERVICE_FLAG} |grep -v 'grep' |wc -l`	
	if [ $count -gt 0 ]; then
		echo "${SERVICE_FLAG} is running. PID is ${pid}"	
		exit
	fi
	nohup ${ROOT_DIR}/bin/${SERVICE_FLAG} > /dev/null 2>&1 &	
	if [ $? -eq 0 ]; then
		pid=`ps -ef | grep $SERVICE_FLAG | grep -v 'grep' | awk '{print $2}'`
		echo "${SERVICE_FLAG} start ok. PID is ${pid}" >> $LOG
		echo $pid > $STWORKER_PIDFILE
	else
		echo "${SERVICE_FLAG} start fiald." >> $LOG
	fi
}

# stop
function stop() {
	pid=`cat ${STWORKER_PIDFILE}`
	count=`ps -ef |grep $SERVICE_FLAG |grep -v grep | wc -l`
	if [ $count -ge 1 ]; then
		`ps -ef |grep "${SERVICE_FLAG}" |grep -v "grep" |awk '{print $2}' |xargs sudo kill -9`	
		echo "${SERVICE_FLAG} is stop." >> $LOG
	else
		echo "${SERVICE_FLAG} is not running." >> $LOG
	fi
}

# restart
function restart() {
	stop
	start
	echo "${SERVICE_FLAG} restart ok." >> $LOG
}

# reload
function reload() {
	# 判断是否启动
	pid=`ps -ef |grep "${SERVICE_FLAG}" |grep -v grep |awk '{print $2}'`	
	if [[ $? != 0 || -z $pid ]]; then
		echo "pid is null" >> $LOG
		start
		exit 0	
	fi	

	# 软重启
	`kill -12 $pid`	
	if [ $? -ge 1 ]; then
		echo "${SERVICE_FLAG} reload fiald." >> $LOG
	else
		echo $pid > $STWORKER_PIDFILE
		echo "${SERVICE_FLAG} reload ok." >> $LOG
	fi
}

case $1 in
	start)
	start
	exit;;
	
	stop)
	stop
	exit;;

	restart)
	restart
	exit;;

	reload)
	reload
	exit;;
	
	*)
	echo "Usage: server.sh {start|stop|restart|reload}";
esac



