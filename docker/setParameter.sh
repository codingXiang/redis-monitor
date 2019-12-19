#!/bin/sh

#!/bin/sh

function log_info()
{
  local date=`date`
  local para=$1
  echo "$date $1"
  $1
  echo "log info:$date $1" &>> $SYS_LOG
}
SYS_LOG=setParameter.log

find replace.sh -a | xargs  grep -r '@' > envSpace.txt
sed 's/[[:space:]]//g' envSpace.txt > env.txt
envList=$(cat env.txt)

for tmp in ${envList}; do
    env=$(cut -d'@' -f2 <<< "$tmp")
    envReplace=${env//./_}
    log_info "sed -i s#@"$env"@#"$(eval echo \$$envReplace)"#g replace.sh"
done

