package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

/********************************************************************
created:    2022-03-19
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

func main() {

}

type Service struct {
	redisClient *RedisClient
	once        sync.Once

	redisClient2 unsafe.Pointer
	lock         sync.Mutex
}

func NewService() *Service {
	var service = &Service{
		redisClient: &RedisClient{},
	}

	return service
}

// todo singleton延迟初始化逻辑, 不应该掺杂在业务逻辑中 -> 单一职责 & DRY
// todo if 仅单线程调用setData1() then NOP; if 存在并发访问 then lost update
func (service *Service) setData1(key string, value string) {
	// singleton
	if service.redisClient == nil {
		service.redisClient, _ = createRedis()
	}

	// business code
	service.redisClient.set(key, value)
}

func (service *Service) setData2(key string, value string) {
	var redis = service.getRedis1()
	redis.set(key, value)
}

func (service *Service) getRedis1() *RedisClient {
	if service.redisClient == nil {
		service.redisClient, _ = createRedis()
	}

	return service.redisClient
}

func (service *Service) getRedis2() *RedisClient {
	service.once.Do(func() {
		service.redisClient, _ = createRedis()
	})

	return service.redisClient
}

func (service *Service) getRedis3() *RedisClient {
	var redis = (*RedisClient)(atomic.LoadPointer(&service.redisClient2))
	if redis == nil {
		service.lock.Lock()
		defer service.lock.Unlock()

		redis = (*RedisClient)(atomic.LoadPointer(&service.redisClient2))
		if redis == nil { // double check
			if redis1, err := createRedis(); err == nil {
				redis = redis1
				atomic.StorePointer(&service.redisClient2, unsafe.Pointer(redis1))
			}
		}
	}

	return redis
}

func incrementSyncMap(m *sync.Map, key int) {
	if m != nil {
		if v, ok := m.Load(key); ok {
			var next = v.(int) + 1
			m.Store(key, next)
		}
	}
}
