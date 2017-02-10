package Session

import (
	"game"
	"sync"
	"time"
)

const CleanTimeerSec = 60 // 1 minute

type Sessions struct {
	sync.RWMutex
	data map[string]*Game.Game
}

var once sync.Once
var list *Sessions

func Start() *Sessions {
	onceBody := func() {
		list = &Sessions{
			data: map[string]*Game.Game{},
		}
		list.StartClean()
	}
	once.Do(onceBody)

	return list
}

func (this *Sessions) FindOrCreate(id string) *Game.Game {

	if id == "" {
		return this.Create()
	}

	user, find := this.Get(id)
	if find {
		return user
	}

	return this.Create()
}

func (this *Sessions) Get(id string) (*Game.Game, bool) {
	this.RLock()
	defer this.RUnlock()
	user, find := this.data[id]
	return user, find
}

func (this *Sessions) Create() *Game.Game {
	this.Lock()
	defer this.Unlock()
	user := Game.New()
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
		if this.data[k].Exp().Before(u) {
			delete(this.data, k)
		}
		v.RUnlock()
	}

}
