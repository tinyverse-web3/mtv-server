#！/bin/bash

log_path=/home/logs

current_date=`date -d "-1 day" "+%Y%m%d"`
echo "$current_date"

cp $log_path/mtv.log $log_path/mtv_$current_date.log
cp /dev/null $log_path/mtv.log
find $log_path -mtime +14 -name 'mtv_*.log' -exec rm -rf {} \;

cp $log_path/qasks.log $log_path/qasks_$current_date.log
cp /dev/null $log_path/qasks.log

find $log_path -mtime +14 -name 'qasks_*.log' -exec rm -rf {} \;
