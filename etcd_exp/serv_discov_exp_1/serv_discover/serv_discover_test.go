// @Author: Perry
// @Date  : 2020/1/2
// @Desc  : 服务注册的简单实现
/*
etcd的租约模式:客户端申请一个租约并设置过期时间，每隔一段时间就要请求etcd申请续租。
客户端可以通过租约存key。如果不续租，过期了，etcd会删除这个租约上的所有key-value。
类似于心跳模式。

一般相同的服务存的key的前缀是一样的比如“server/001"=>"127.0.0.1:1212"
和”server/002"=>"127.0.0.1:1313"这种模式，然后客户端就直接匹配“server/”这个key
*/

package serv_discover

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"testing"
	"time"
)

type ServiceReg struct {
	client        *clientv3.Client             /*客户端对象*/
	lease         clientv3.Lease               /*租约对象*/
	leaseResp     *clientv3.LeaseGrantResponse /*租约应答对象*/
	cancelFunc    func()
	keepAliveChan <-chan *clientv3.LeaseKeepAliveResponse /*保活应答*/
}

func NewServiceReg(addr []string, timeNum int64) (*ServiceReg, error) {
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	var client *clientv3.Client

	if clientTem, err := clientv3.New(conf); err != nil {
		return nil, err
	} else {
		client = clientTem
	}

	ser := &ServiceReg{
		client: client,
	}

	if err := ser.setLease(timeNum); err != nil {
		return nil, err
	}
	go ser.ListenLeaseRespChan()
	return ser, nil
}

/*设置租约*/
func (s *ServiceReg) setLease(timeNum int64) error {
	lease := clientv3.NewLease(s.client)

	/*设置租约时间*/
	leaseResp, err := lease.Grant(context.TODO(), timeNum)
	if err != nil {
		return err
	}
	/*设置续租*/
	ctx, cancelFunc := context.WithCancel(context.TODO())
	leaseRespChan, err := lease.KeepAlive(ctx, leaseResp.ID)
	if err != nil {
		return err
	}

	s.lease = lease
	s.leaseResp = leaseResp
	s.cancelFunc = cancelFunc
	s.keepAliveChan = leaseRespChan
	return nil
}

/*监听续租情况*/
func (s *ServiceReg) ListenLeaseRespChan() {
	for {
		select {
		case leaseKeepResp := <-s.keepAliveChan:
			if leaseKeepResp == nil {
				log.Printf("已经关闭续租功能\n")
				return
			} else {
				log.Printf("续租成功\n")
			}
		}
	}
}

/*通过租约注册服务*/
func (s *ServiceReg) PutService(key, val string) error {
	kv := clientv3.NewKV(s.client)
	_, err := kv.Put(context.TODO(), key, val, clientv3.WithLease(s.leaseResp.ID))
	return err
}

/*撤销租约*/
func (s *ServiceReg) RevokeLease() error {
	s.cancelFunc()
	time.Sleep(2 * time.Second)
	_, err := s.lease.Revoke(context.TODO(), s.leaseResp.ID)
	return err
}

func Test1(t *testing.T) {
	fmt.Println("starting")
	ser, _ := NewServiceReg([]string{"127.0.0.1:2379"}, 5)
	_ = ser.PutService("/node/111", "daipengyuan")
	_ = ser.PutService("/node/112", "daipengyuan2")
	select {}
}
