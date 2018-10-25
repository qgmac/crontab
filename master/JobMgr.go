package master

import (
	"fmt"
	"github.com/etcd-io/etcd/clientv3"
	"time"
)

type JobMgr struct {
	client * clientv3.Client
	kv clientv3.KV
	lease clientv3.Lease
}

var(
	//单例
	G_jobMgr *JobMgr
)

//初始化管理器
func InitJobMgr()(err error){
	var(
		config clientv3.Config
		client *clientv3.Client
		kv clientv3.KV
		lease clientv3.Lease
	)
	fmt.Println(G_config.EtcdEndpoints)
	//初始化配置
	config = clientv3.Config{
		Endpoints:G_config.EtcdEndpoints,//集群地址
		DialTimeout:time.Duration(G_config.EtcdDialTimeout),//连接超时
	}
	//建立连接
	if client,err = clientv3.New(config);err != nil {
		return
	}

	//得到kv和Lease的API子集
	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)


	G_jobMgr = &JobMgr{
		client:client,
		kv:kv,
		lease:lease,
	}
	return
}

