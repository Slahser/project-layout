package etcd_wrapper

import (
	"sync"
	"time"

	etcd "go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

var (
	once       sync.Once
	etcdClient *etcd.Client

	etcdEndpoints = []string{"http://8.210.134.212:2379"}
	//etcdUsername    = "tt"
	//etcdPassword    = "tt"
	etcdDialTimeout = 5 * time.Second

	zapDevConfig = zap.NewDevelopmentConfig()
)

func init() {
	InitEtcdClient()
}

func GetEtcdKv() etcd.KV {
	return etcd.NewKV(GetEtcdClient())
}

func GetEtcdClient() *etcd.Client {
	once.Do(func() {
		etcdClient = InitEtcdClient()
	})
	return etcdClient
}

func InitEtcdClient() *etcd.Client {

	cli, err := etcd.New(etcd.Config{
		Endpoints:   etcdEndpoints,
		DialTimeout: etcdDialTimeout,
		//Username:    etcdUsername,
		//Password:    etcdPassword,
		LogConfig: &zapDevConfig,
	})

	if err != nil {
		zap.S().Errorw("etcd init error")
	}

	defer cli.Close()

	zap.S().Info("etcd init succeed")
	return cli
}
