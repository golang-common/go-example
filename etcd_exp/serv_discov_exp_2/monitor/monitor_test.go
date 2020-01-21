// @Author: Perry
// @Date  : 2020/1/2
// @Desc  : 服务监听，watch服务的变化并实时更新

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"testing"
	"time"
)

type Monitor struct {
	Path   string
	Nodes  map[string]*Node
	Client *clientv3.Client
}

type Node struct {
	State bool
	Key   string
	Info  ServiceInfo
}

type ServiceInfo struct {
	IP string
}

func NewMonitor(endpoints []string, watchPath string) (*Monitor, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	monitor := &Monitor{
		Path:   watchPath,
		Nodes:  make(map[string]*Node),
		Client: cli,
	}

	go monitor.WatchNodes()
	return monitor, err
}

func (s *Monitor) AddNode(key string, info *ServiceInfo) {
	node := &Node{
		State: true,
		Key:   key,
		Info:  *info,
	}
	s.Nodes[node.Key] = node
}

func (s *Monitor) GetServiceInfo(ev *clientv3.Event) *ServiceInfo {
	info := &ServiceInfo{}
	err := json.Unmarshal([]byte(ev.Kv.Value), info)
	if err != nil {
		log.Println(err)
	}
	return info
}

func (s *Monitor) WatchNodes() {
	rch := s.Client.Watch(context.Background(), s.Path, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				log.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				info := s.GetServiceInfo(ev)
				s.AddNode(string(ev.Kv.Key), info)
			case clientv3.EventTypeDelete:
				log.Printf("[%s] %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				delete(s.Nodes, string(ev.Kv.Key))
			}
		}
	}
}

func Test1(t *testing.T) {
	m, err := NewMonitor([]string{
		"http://127.0.0.1:2379",
	}, "/node")
	if err != nil {
		log.Fatal(err)
	}

	for {
		for k, v := range m.Nodes {
			fmt.Printf("node:%s, ip=%s\n", k, v.Info.IP)
		}
		fmt.Printf("nodes num = %d\n", len(m.Nodes))
		time.Sleep(time.Second * 5)
	}
}
