hugeadm --pool-pages-min 2MB:4096
cat /proc/sys/vm/nr_hugepages 
ls /sys/kernel/mm/transparent_hugepage
cat /proc/meminfo
ls -lah /dev/hugepages/

docker run -itd --device=/dev/hugepages:/dev/hugepages --privileged -v `pwd`/docker/vpp-centos-userspace-cni/:/etc/vpp/ --name vpp ligato/vpp-base
sleep 130
docker logs  vpp

sleep 10
docker ps -a
