package etcd

import (
	"go.etcd.io/etcd/clientv3"
	"gogw-server/logger"
	"sync"
)

var etcdClient *clientv3.Client
var onceEtcdClient sync.Once

func GetEtcdClient() *clientv3.Client {
	if etcdClient == nil {
		onceEtcdClient.Do(func() {
			cfg := clientv3.Config{
				Endpoints: []string{"http://47.110.37.178:2379"},
			}
			var err error
			etcdClient, err = clientv3.New(cfg)
			if err != nil {
				logger.Error("new etcd err:", err)
			}
		})
	}
	return etcdClient
}
