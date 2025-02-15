
# step1: start etcd cluster in the background

nohup etcd --name "$NODENAME" --initial-advertise-peer-urls http://"$PRIVATEIPV4":2380 \
      --listen-peer-urls http://"$PRIVATEIPV4":2380 --listen-client-urls http://"$PRIVATEIPV4":2379,http://127.0.0.1:2379 \
      --advertise-client-urls http://"$PRIVATEIPV4":2379 --initial-cluster-token etcd-cluster \
      --initial-cluster k8sdev-master=http://172.31.22.213:2380,k8sdev-worker1=http://172.31.30.34:2380,k8sdev-worker2=http://172.31.22.64:2380 \
      --initial-cluster-state new > /tmp/etcd.log 2>&1 &

# step2: start flanneld

nohup sudo ./flanneld -iface="$PRIVATEIPV4" > /tmp/flanneld.log 2>&1 &

# step3: adjust iptables

sudo iptables -P FORWARD ACCEPT
