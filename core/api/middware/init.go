package middware

import (
	"context"
	"coredx/log"
	"github.com/patrickmn/go-cache"
	"time"
)

var LimitCache *cache.Cache
var blackList *BlackList

type BlackList struct {
	blackList map[string]*Black
}
type Black struct {
	Ip       string
	ForbidTs int64
}

func (b *BlackList) find(ip string) bool {
	return b.blackList[ip] != nil
}

func (b *BlackList) add(ip string) {
	b.blackList[ip] = &Black{}
	b.blackList[ip].Ip = ip
	b.blackList[ip].ForbidTs = time.Now().Unix()
	log.Debugf("add black:%s", ip)
}
func (b *BlackList) clear() {
	for k, v := range blackList.blackList {
		log.Debug(k, v)
		log.Debug(time.Now().Unix() - v.ForbidTs)
		if time.Now().Unix()-v.ForbidTs >= 60*3 {
			log.Debugf("del black ip:%s", k)
			delete(b.blackList, k)
		}
	}
}

//var BlackList *cache.Cache

func Init(ctx context.Context) {
	LimitCache = cache.New(300*time.Second, time.Minute*10)
	blackList = &BlackList{blackList: make(map[string]*Black, 10)}
	go func(cx context.Context, bl *BlackList) {
		for {
			select {
			case <-cx.Done():
				log.Info("limiter recv cancel")
				return
			case <-time.After(time.Second * 10):
				//删除限流ip
				bl.clear()
			}
		}
	}(ctx, blackList)
}
