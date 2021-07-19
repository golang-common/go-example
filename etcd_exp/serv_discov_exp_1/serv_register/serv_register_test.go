// @Author: Perry
// @Date  : 2020/1/2
// @Desc  : 服务注册
/*
1.创建一个client连到etcd。
2.匹配到所有相同前缀的key。把值存到serverList这个map里面。
3.watch这个key前缀，当有增加或者删除的时候就修改这个map。
4.所以这个map就是实时的服务列表
*/

package serv_register

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
	"testing"
	"time"
)

func NewClientDis(addr []string) (*ClientDis, error) {
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	if client, err := clientv3.New(conf); err == nil {
		return &ClientDis{
			client:     client,
			serverList: make(map[string]string),
		}, nil
	} else {
		return nil, err
	}
}

type ClientDis struct {
	client     *clientv3.Client
	serverList map[string]string /*用来存储服务内容的字典*/
	lock       sync.Mutex
}

func (s *ClientDis) GetService(prefix string) ([]string, error) {
	resp, err := s.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	addrs := s.extractAddrs(resp)

	go s.watcher(prefix)
	return addrs, nil
}

/*监听etcd服务的变化,监听到变化后给*/
func (s *ClientDis) watcher(prefix string) {
	watchChan := s.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range watchChan {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case clientv3.EventTypeDelete:
				s.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

/*传入etcd应答,设置结构的serviceList值,返回服务列表*/
func (s *ClientDis) extractAddrs(resp *clientv3.GetResponse) []string {
	addrs := make([]string, 0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			s.SetServiceList(string(resp.Kvs[i].Key), string(resp.Kvs[i].Value))
			addrs = append(addrs, string(v))
		}
	}
	return addrs
}

/*设置本地服务map的值*/
func (s *ClientDis) SetServiceList(key, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.serverList[key] = val
	log.Println("set data key:", key, "val:", val)
}

/*删除本地服务map的项*/
func (s *ClientDis) DelServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.serverList, key)
	log.Println("del data key:", key)
}

/*返回本地的所有服务列表*/
func (s *ClientDis) SerList2Array() []string {
	s.lock.Lock()
	defer s.lock.Unlock()
	addrs := make([]string, 0)
	for _, v := range s.serverList {
		addrs = append(addrs, v)
	}
	return addrs
}

func Test1(t *testing.T) {
	fmt.Println("starting")
	cli, _ := NewClientDis([]string{"127.0.0.1:2379"})
	cli.GetService("/node")
	select {}
}
