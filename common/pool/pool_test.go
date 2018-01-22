package pool

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const (
	//模拟的最大goroutine
	maxGoroutine = 5
	//资源池的大小
	poolRes = 2
)

func TestPool(t *testing.T) {
	//等待任务完成
	var wg sync.WaitGroup
	wg.Add(maxGoroutine)
	p, err := NewPool(createConnection, poolRes)
	if err != nil {
		log.Println(err)
		return
	}
	//模拟好几个goroutine同时使用资源池查询数据
	for query := 0; query < maxGoroutine; query++ {
		go func(q int) {
			dbQuery(q, p)
			wg.Done()
		}(query)
	}
	wg.Wait()
	log.Println("开始关闭资源池")
	p.Close()
}

//模拟数据库查询
func dbQuery(query int, pool *Pool) {
	conn, err := pool.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	defer pool.Release(conn)
	//模拟查询
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("第%d个查询，使用的是ID为%d的数据库连接", query, conn.(*dbConnection).ID)
}

//数据库连接
type dbConnection struct {
	ID int32 //连接的标志
}

//实现io.Closer接口
func (db *dbConnection) Close() error {
	log.Println("关闭连接", db.ID)
	return nil
}

var idCounter int32

//生成数据库连接的方法，以供资源池使用
func createConnection() (io.Closer, error) {
	//并发安全，给数据库连接生成唯一标志
	id := atomic.AddInt32(&idCounter, 1)
	return &dbConnection{id}, nil
}
