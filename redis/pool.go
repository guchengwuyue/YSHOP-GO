package redis

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/gomodule/redigo/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Pool *redis.Pool
)

func init() {
	host, _ := beego.AppConfig.String("redis_host")
	pwd, _ := beego.AppConfig.String("redis_pwd")
	Pool = newPool(host,pwd)
	cleanupHook()
}

func newPool(server string, pwd string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			//处理密码
			if pwd != "" {
				_, err = c.Do("AUTH",pwd)
				if err != nil {
					return nil, err
				}
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}