package main

import (
	"fmt"
	"github.com/gogf/gf/database/gredis"
	"github.com/gogf/gf/util/gconv"
)

type RedisMap struct {
	redis *gredis.Redis
}
func NewRedisMap(host string, port int) *RedisMap{
	config := &gredis.Config{
		Host:host,
		Port:port,
	}
	group := "default"
	gredis.SetConfig(config,group)
	redisMap := &RedisMap{}
	redisMap.redis = gredis.Instance(group)
	return redisMap
}
func (this *RedisMap) Get(cf string, key string) (string,error) {
	realKey := cf + key
	value,err := this.redis.Do("Get",realKey)
	if err != nil {
		return "",err
	} else {
		return gconv.String(value),nil
	}
}
func (this *RedisMap) Set(cf string, key string, value string) error {
	realKey := cf + key
	_,err := this.redis.Do("Set",realKey,value)
	return err
}
func main() {
	host := "127.0.0.1"
	port := 6379
	redisMap := NewRedisMap(host,port)
	cf := "num_to_url"
	err := redisMap.Set(cf,"1234","testurl")
	if err != nil {
		fmt.Println(err)
	}
	value,_ := redisMap.Get(cf,"1234")
	fmt.Println(value)
}-
