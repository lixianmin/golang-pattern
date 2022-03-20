package main

/********************************************************************
created:    2022-03-19
author:     lixianmin

Copyright (C) - All Rights Reserved
*********************************************************************/

type RedisClient struct {
}

func createRedis() (*RedisClient, error) {
	return &RedisClient{}, nil
}

func (my *RedisClient) set(key string, value string) {

}
