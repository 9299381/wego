package idwork

import (
	"errors"
	"github.com/9299381/wego/constants"
	"strconv"
	"sync"
	"time"
)

var once sync.Once
var ins *snowflake

func getID(server int64) string {
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
func (s *snowflake) epochGen() int64 {
	start := "2010-01-01 00:00:00"
	layout := constants.YmdHis
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(layout, start, loc)
	return theTime.UnixNano() / 1000000
}

func (s *snowflake) nextId() string {
	s.mu.Lock()
	epoch := s.epochGen()
	timestamp := time.Now().UnixNano() / 1000000
	lastTime := s.lastTime
	//生成唯一序列
	if timestamp == lastTime {
		s.sequence = (s.sequence + 1) & s.sequenceMask
		if s.sequence == 0 {
			for timestamp <= lastTime {
				timestamp = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		s.sequence = 0
	}
	s.lastTime = timestamp
	r := int64((timestamp-epoch)<<s.timeShift |
		(s.server << s.serverShift) |
		(s.sequence),
	)
	s.mu.Unlock()
	return strconv.FormatInt(int64(r), 10)
}
