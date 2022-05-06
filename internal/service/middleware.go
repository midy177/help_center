package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"runtime"
	"sync"
	"time"
)

// MidRecover 恢复中间件
func MidRecover(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				stack := make([]byte, 1<<10)
				length := runtime.Stack(stack, false)
				// stdlog.Println(string(stack[:length]))
				os.Stdout.Write(stack[:length])
				ctx.Error(err)
			}
		}()
		return next(ctx)
	}
}

var CacheMap sync.Map

type Cache struct {
	Expire time.Time
	Data   interface{}
}

func GetCache(name string) (bool, interface{}) {
	v, ok := CacheMap.Load(name)
	if ok {
		if c, isOk := v.(Cache); isOk {
			if time.Until(c.Expire) >= 0 {
				return true, c.Data
			} else {
				CacheMap.Delete(name)
				return false, nil
			}
		}
	}
	return false, nil
}
func StoreCache(name string, data interface{}) {
	expire := time.Now().Add(time.Minute * 5)
	CacheMap.Store(name, Cache{Expire: expire, Data: data})
}
func CacheMidPre(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "GET" {
			ok, data := GetCache(c.Request().RequestURI)
			if ok {
				return c.JSON(http.StatusOK, data)
			}
		}
		return next(c)
	}
}
