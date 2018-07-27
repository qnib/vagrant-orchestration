# Docker CE multi-networks

This vagrant setup create three nodes with multiple network interfaces:

```bash
ip -o -4 add
1: lo    inet 127.0.0.1/8 scope host lo\       valid_lft forever preferred_lft forever
2: eth0    inet 10.0.2.15/24 brd 10.0.2.255 scope global eth0\       valid_lft forever preferred_lft forever
3: eth1    inet 192.168.100.10/24 brd 192.168.100.255 scope global eth1\       valid_lft forever preferred_lft forever
4: eth2    inet 192.168.101.10/24 brd 192.168.101.255 scope global eth2\       valid_lft forever preferred_lft forever
5: docker0    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0\       valid_lft forever preferred_lft forever
10: docker_gwbridge    inet 172.18.0.1/16 brd 172.18.255.255 scope global docker_gwbridge\       valid_lft forever preferred_lft forever
```

`eth1` and `eth2` are shared between the nodes.

By using `docker swarm init --advertise-addr=192.168.100.10 --data-path-addr=192.168.101.10` the management and the data plane are separated.

### pinger

This simple stack creates a set of services pinging each other.

```
vagrant@swarm0:/vagrant$ sudo docker stack deploy -c docker-compose.yml test
Creating network test_testnet
Creating service test_srv1
Creating service test_srv2
Creating service test_srv0
```
