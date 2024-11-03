package redis

import "errors"

var (
	ErrorNoneToken = errors.New("token in redis not exist")
	ErrorNoneCap   = errors.New("capability in redis not exist")
)
