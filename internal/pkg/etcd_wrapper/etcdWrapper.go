package etcdwrapper

import (
	"sync"
	"time"

	etcd "github.com/coreos/etcd/clientv3"
	"go.uber.org/zap"
)

var (
	once       sync.Once
	EtcdClient *etcd.Client
	EtcdKV     etcd.KV

	etcdEndpoints = []string{"8.210.134.212:2379"}
	//etcdUsername    = "tt"
	//etcdPassword    = "tt"
	etcdDialTimeout = 5 * time.Second

	zapDevConfig = zap.NewDevelopmentConfig()
)

func init() {
	InitEtcdClient()
}

func InitEtcdClient() *etcd.Client {

	EtcdClient, err := etcd.New(etcd.Config{
		Endpoints:   etcdEndpoints,
		DialTimeout: etcdDialTimeout,
		//Username:    etcdUsername,
		//Password:    etcdPassword,
		LogConfig: &zapDevConfig,
	})

	if err != nil {
		zap.S().Errorw("etcd init error")
	}

	EtcdKV = etcd.NewKV(EtcdClient)
	zap.S().Info("etcd init succeed")
	return EtcdClient
}

func GetEtcdKv() etcd.KV {
	return etcd.NewKV(GetEtcdClient())
}

func GetEtcdClient() *etcd.Client {
	once.Do(func() {
		EtcdClient = InitEtcdClient()
	})
	return EtcdClient
}
