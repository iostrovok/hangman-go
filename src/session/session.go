package Session

import (
	"sync"
	"time"
)

const KeyLength = 10
const CleanTimeerSec = 60        // 1 minute
const LiveTimeSec = 60 * 60 * 24 // 1 day

type Sessions struct {
	sync.RWMutex
	data map[string]*UserGame
}

var once sync.Once
var list *Sessions

func Start() *Sessions {
	onceBody := func() {
		list = &Sessions{
			data: map[string]*UserGame{},
		}
		list.StartClean()
	}
	once.Do(onceBody)

	return list
}

func (this *Sessions) FindOrCreate(id string) *UserGame {

	if id == "" {
		return this.Create()
	}

	user, find := this.Get(id)
	if find {
		return user
	}

	return this.Create()
}

func (this *Sessions) Get(id string) (*UserGame, bool) {
	this.RLock()
	defer this.RUnlock()
	user, find := this.data[id]
	return user, find
}

func (this *Sessions) Create() *UserGame {
	this.Lock()
	defer this.Unlock()
	user := newUserGame()
	this.data[user.ID] = user
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
