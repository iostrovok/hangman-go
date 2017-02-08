package Session

import (
	"fmt"
	"sync"
	"time"
)

const KeyLength = 10
const CleanTimeerSec = 60        // 1 minute
const LiveTimeSec = 60 * 60 * 24 // 1 day

type User struct {
	sync.RWMutex

	ID   string
	exp  time.Time
	data map[string]interface{}
}

type Sessions struct {
	sync.RWMutex
	data map[string]*User
}

var once sync.Once
var list *Sessions

func Start() *Sessions {
	onceBody := func() {
		fmt.Println("Only once")
		list = &Sessions{
			data: map[string]*User{},
		}
		list.StartClean()
	}
	once.Do(onceBody)

	return list
}

func (this *Sessions) FindOrCreate(id string) *User {

	if id == "" {
		return this.Create()
	}

	user, find := this.Get(id)
	if find {
		return user
	}

	return this.Create()
}

func (this *Sessions) Get(id string) (*User, bool) {
	this.RLock()
	defer this.RUnlock()
	user, find := this.data[id]
	return user, find
}

func (this *Sessions) Create() *User {

	id := RandString(KeyLength)

	this.Lock()
	defer this.Unlock()
	user := &User{
		ID:   id,
		exp:  time.Now().Add(LiveTimeSec * time.Second),
		data: map[string]interface{}{},
	}

	this.data[id] = user
	return user
}

func (this *Sessions) StartClean() {
	go func() {
		for {
			time.Sleep(CleanTimeerSec * time.Second)
			this.Clean()
		}
	}()
}

func (this *Sessions) Clean() {

	u := time.Now()

	this.Lock()
	defer this.Unlock()

	for k, v := range this.data {
		v.RLock()
		if this.data[k].exp.Before(u) {
			delete(this.data, k)
		}
		v.RUnlock()
	}

}
