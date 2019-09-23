package snowflake

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

var once sync.Once
var ins *snowflake

func GetID(server int64) string {
	n := getIns(server)
	return n.nextId()

}
func getIns(server int64) *snowflake {
	once.Do(func() {
		var err error
		ins, err = newNode(server)
		if err != nil {
			panic(err)
		}
	})
	return ins
}

type snowflake struct {
	mu       sync.Mutex
	lastTime int64

	server    int64
	serverMax int64

	sequence     int64
	sequenceMask int64

	timeShift   uint8
	serverShift uint8
}

func newNode(server int64) (*snowflake, error) {
	var serverBits uint8 = 10
	var sequenceBits uint8 = 12
	it := snowflake{}

	it.server = server
	it.serverMax = -1 ^ (-1 << serverBits)
	it.sequenceMask = -1 ^ (-1 << sequenceBits)

	it.serverShift = sequenceBits
	it.timeShift = serverBits + sequenceBits

	if it.server < 0 || it.server > it.serverMax {
		return nil, errors.New("server number must be between 0 and " + strconv.FormatInt(it.serverMax, 10))
	}

	return &it, nil
}
func (it *snowflake) epochGen() int64 {
	start := "2010-01-01 00:00:00"
	layout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(layout, start, loc)
	return theTime.UnixNano() / 1000000
}

func (it *snowflake) nextId() string {
	it.mu.Lock()
	epoch := it.epochGen()
	timestamp := time.Now().UnixNano() / 1000000
	lastTime := it.lastTime
	//生成唯一序列
	if timestamp == lastTime {
		it.sequence = (it.sequence + 1) & it.sequenceMask
		if it.sequence == 0 {
			for timestamp <= lastTime {
				timestamp = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		it.sequence = 0
	}
	it.lastTime = timestamp
	r := int64((timestamp-epoch)<<it.timeShift |
		(it.server << it.serverShift) |
		(it.sequence),
	)
	it.mu.Unlock()
	return strconv.FormatInt(int64(r), 10)
}
