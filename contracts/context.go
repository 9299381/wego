package contracts

import (
	"context"
	"strings"
	"time"
)

type Context struct {
	context.Context
	Keys map[string]interface{}
	Log  ILogger
}

func (c *Context) reset() {
	c.Keys = nil
}

func (c *Context) Copy() *Context {
	var cp = *c
	cp.Keys = make(map[string]interface{})
	for k, v := range c.Keys {
		cp.Keys[k] = v
	}
	return &cp
}

// Set is used to store a new key/value pair exclusively for this contexts.
// It also lazy initializes  c.Keys if it was not used previously.
func (c *Context) Set(key string, value interface{}) {
	if c.Keys == nil {
		c.Keys = make(map[string]interface{})
	}
	c.Keys[key] = value
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exists it returns (nil, false)
func (c *Context) Get(key string) (value interface{}, exists bool) {
	value, exists = c.Keys[key]
	return
}

// MustGet returns the value for the given key if it exists, otherwise it panics.
func (c *Context) MustGet(key string) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}

// GetString returns the value associated with the key as a string.
func (c *Context) GetString(key string) (s string) {
	if val, ok := c.Get(key); ok && val != nil {
		s, _ = val.(string)
	}
	return
}

// GetBool returns the value associated with the key as a boolean.
func (c *Context) GetBool(key string) (b bool) {
	if val, ok := c.Get(key); ok && val != nil {
		b, _ = val.(bool)
	}
	return
}

// GetInt returns the value associated with the key as an integer.
func (c *Context) GetInt(key string) (i int) {
	if val, ok := c.Get(key); ok && val != nil {
		i, _ = val.(int)
	}
	return
}

// GetInt64 returns the value associated with the key as an integer.
func (c *Context) GetInt64(key string) (i64 int64) {
	if val, ok := c.Get(key); ok && val != nil {
		i64, _ = val.(int64)
	}
	return
}

// GetFloat64 returns the value associated with the key as a float64.
func (c *Context) GetFloat64(key string) (f64 float64) {
	if val, ok := c.Get(key); ok && val != nil {
		f64, _ = val.(float64)
	}
	return
}

// GetTime returns the value associated with the key as time.
func (c *Context) GetTime(key string) (t time.Time) {
	if val, ok := c.Get(key); ok && val != nil {
		t, _ = val.(time.Time)
	}
	return
}

// GetDuration returns the value associated with the key as a duration.
func (c *Context) GetDuration(key string) (d time.Duration) {
	if val, ok := c.Get(key); ok && val != nil {
		d, _ = val.(time.Duration)
	}
	return
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (c *Context) GetStringSlice(key string) (ss []string) {
	if val, ok := c.Get(key); ok && val != nil {
		ss, _ = val.([]string)
	}
	return
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (c *Context) GetStringMap(key string) (sm map[string]interface{}) {
	if val, ok := c.Get(key); ok && val != nil {
		sm, _ = val.(map[string]interface{})
	}
	return
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (c *Context) GetStringMapString(key string) (sms map[string]string) {
	if val, ok := c.Get(key); ok && val != nil {
		sms, _ = val.(map[string]string)
	}
	return
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func (c *Context) GetStringMapStringSlice(key string) (smss map[string][]string) {
	if val, ok := c.Get(key); ok && val != nil {
		smss, _ = val.(map[string][]string)
	}
	return
}

func (c *Context) SetValue(key string, value interface{}) {
	keyMap := strings.Split(key, ".")
	if len(keyMap) == 1 {
		c.Set(keyMap[0], value)
	} else {
		pos := c.getPos(keyMap)
		if pos > 0 {
			c.modifyValue(keyMap, pos, value)
		} else {
			ret := c.setKeyToMap(keyMap[1:], value)
			c.Set(keyMap[0], ret)
		}
	}
}

func (c *Context) GetValue(key string) (ret interface{}) {
	keyMap := strings.Split(key, ".")
	ret = c.getKeyFromMap(keyMap, c.Keys)
	return
}

func (c *Context) Request(key ...string) (ret interface{}) {
	var k string
	if key == nil {
		k = "request"
	} else {
		k = "request." + key[0]
	}
	return c.GetValue(k)
}
func (c *Context) Response(key string, value interface{}) {
	k := "response." + key
	c.SetValue(k, value)
}

//------私有

func (c *Context) setKeyToMap(keyMap []string, value interface{}) (ret map[string]interface{}) {
	//pos为
	ret = make(map[string]interface{})
	if len(keyMap) == 1 {
		ret[keyMap[0]] = value
	} else if len(keyMap) > 1 {
		ret[keyMap[0]] = c.setKeyToMap(keyMap[1:], value)
	}
	return
}

/**
获取有值的位置
*/
func (c *Context) getPos(keyMap []string) (pos int) {
	l := len(keyMap)
	pos = 0
	//是否有值,如果有值应该叠加上,从最后开始
	for i := 0; i < l; i++ {
		newMap := keyMap[:l-i]
		old := c.getKeyFromMap(newMap, c.Keys)
		if old != nil {
			pos = l - i
			break
		}
	}
	return
}

func (c *Context) getKeyFromMap(keyMap []string, valueMap interface{}) (ret interface{}) {
	m, ok := valueMap.(map[string]interface{})
	if ok {
		if len(keyMap) == 0 {
			ret = nil
		} else if len(keyMap) == 1 {
			ret = m[keyMap[0]]
		} else {
			ret = c.getKeyFromMap(keyMap[1:], m[keyMap[0]])
		}
	}
	return
}
func (c *Context) modifyValue(keyMap []string, pos int, value interface{}) {
	//在函数调用时，像切片（slice）、字典（map）、
	// 接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递
	// （即使没有显式的指出指针）。
	ret, ok := c.getKeyFromMap(keyMap[:pos], c.Keys).(map[string]interface{})
	if ok {
		for v, k := range c.setKeyToMap(keyMap[pos:], value) {
			//注意这里的ret为指针,修改其值则c.key中值发生变化
			ret[v] = k
		}
	} else {
		if pos > 1 {
			c.modifyValue(keyMap, pos-1, value)
		}
	}
}
