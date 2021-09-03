package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"
)

type KeyValue struct {
	Key   string      `json:"key"`
	Sense interface{} `json:"value"`
	Ttl   int64       `json:"ttl"`
}

/*
CreateSec – time of object creation in UNIX timestamp in seconds.
TtlSec – lifetime of the object. If TtlSec is 0, then object is "immortal".
*/

type Vitals struct {
	CreateSec int64
	TtlSec    int64
}

type Value struct {
	Value interface{}
	Timer Vitals
}

type RequestValue struct {
	Key string `json:"key"`
}

type PatternValue struct {
	Pattern string `json:"key"`
}

type Storage interface {
	Set(kv *KeyValue) error
	Get(gv *RequestValue) (interface{}, error)
	Delete(gv *RequestValue) error
	Keys(pv *PatternValue) ([]string, error)
}

type MemoryStorage struct {
	data map[string]Value
	sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]Value),
	}
}

func (s *MemoryStorage) Set(kv *KeyValue) error {
	if kv.Ttl < 0 {
		return errors.New("Check the ttl.")
	}

	if err := checkKey(kv.Sense); err != nil {
		return err
	}

	s.Lock()
	s.data[kv.Key] = Value{
		Value: kv.Sense,
		Timer: Vitals{
			CreateSec: time.Now().Unix(),
			TtlSec:    kv.Ttl,
		},
	}
	s.Unlock()
	return nil
}

func (s *MemoryStorage) Get(rv *RequestValue) (interface{}, error) {
	s.RLock()
	currentSec := time.Now().Unix()
	value, exist := s.data[rv.Key]
	if !exist {
		s.RUnlock()
		return nil, errors.New("Does not exist.")
	}
	if value.Timer.TtlSec == 0 || value.Timer.CreateSec+value.Timer.TtlSec >= currentSec {
		s.RUnlock()
		return value.Value, nil
	}
	s.RUnlock()
	go s.Delete(rv)
	return nil, errors.New("The key has run out of life time.")
}

func (s *MemoryStorage) Delete(rv *RequestValue) error {
	s.Lock()
	_, exist := s.data[rv.Key]
	if exist {
		delete(s.data, rv.Key)
	}
	s.Unlock()
	return nil
}

func (s *MemoryStorage) Keys(pv *PatternValue) ([]string, error) {
	s.RLock()
	currentSec := time.Now().Unix()
	var keys []string
	for k, value := range s.data {
		if checkPattern(pv.Pattern, k) && (value.Timer.TtlSec == 0 || value.Timer.CreateSec+value.Timer.TtlSec >= currentSec) {
			keys = append(keys, k)
		}
	}
	s.RUnlock()
	return keys, nil
}

func checkPattern(search string, key string) bool {
	pattern := `^` + strings.Replace(strings.Replace(search, `*`, `.*`, -1), `?`, `.`, -1) + `$`
	fmt.Println(pattern)
	matched, _ := regexp.MatchString(pattern, key)
	return matched
}

func checkKey(i interface{}) error {
	_, convertString := i.(string)
	_, convertList := i.([]interface{})
	_, convertMapString := i.(map[string]interface{})
	/*Last convertion – convert in "True" map. In this realization must be false, but in some vertion may be true.*/
	_, convertMap := i.(map[interface{}]interface{})
	if convertString || convertList || convertMapString || convertMap {
		return nil
	}
	return errors.New("Something wrong. Check your value.")
}
