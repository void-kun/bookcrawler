echo "Dum sql from docker"

$time = Get-Date -Format "yyyyMMddHHmm"
docker exec -it bw-database pg_dumpall -c -U zrik > ./docker/postgres/dump_$time.sql