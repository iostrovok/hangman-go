package Session

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"
)

var clenerReg = regexp.MustCompile(`[^a-zA-Z]`)

const CountAttempt = 9

type UserGame struct {
	sync.RWMutex

	ID            string
	exp           time.Time
	wins          int
	losings       int
	word          []string
	check_letters map[string]bool
	view_word     []string
	last_word     []string
	last_letters  []string
	attempt       int
	last_attempt  bool
	last_letter   string
	last_state    string
}

func newUserGame() *UserGame {
	id := RandString(KeyLength)

	user := &UserGame{
		ID:            id,
		exp:           time.Now().Add(LiveTimeSec * time.Second),
		word:          []string{},
		last_word:     []string{},
		view_word:     []string{},
		check_letters: map[string]bool{},
		last_letters:  []string{},
	}

	return user
}

/*
	Interface functions. HAVE TO USE lock/unlock into these functions
*/

func (this *UserGame) Info() map[string]interface{} {
	this.Lock()
	defer this.Unlock()
	this.UpTime()
	return this.outData(nil)
}

func (this *UserGame) NewWord(word string) map[string]interface{} {
	this.Lock()
	defer this.Unlock()
	this.UpTime()

	if this.IsPlayNow() {
		// TODO a cheat or old cookie
		return this.outData(errors.New("Already in the game"))
	}

	this.attempt = 0
	this.check_letters = map[string]bool{}
	this.last_letters = []string{}
	this.word = strings.Split(word, "")
	this.view_word = strings.Split(word, "")

	for v := range this.view_word {
		this.view_word[v] = "*"
	}

	this.last_state = "is_play_now"

	return this.outData(nil)
}

func (this *UserGame) Move(s string) map[string]interface{} {
	this.Lock()
	defer this.Unlock()
	this.UpTime()

	// ----> check params start
	this.last_letter = clenerReg.ReplaceAllString(strings.ToLower(s), "")

	if !this.IsPlayNow() {
		// TODO a cheat or old cookie
		return this.outData(errors.New("not_in_the_game"))
	}

	if len(this.last_letter) != 1 {
		return this.outData(errors.New("bad_move"))
	}

	if this.check_letters[this.last_letter] {
		return this.outData(errors.New("duplicate"))
	}
	this.last_letters = append(this.last_letters, this.last_letter)
	this.check_letters[this.last_letter] = true
	// <---- check params finish

	// Check letter in word
	find := 0
	for k, v := range this.word {
		if v == this.last_letter {
			this.view_word[k] = v
			find++
		}
	}

	// Unsuccessful turn
	this.last_attempt = true
	if find == 0 {
		this.last_attempt = false
		this.attempt++
	}

	// Check "game over"
	if this.attempt >= CountAttempt {
		this.defeat()
	} else {
		this.checkWin()
	}

	return this.outData(nil)
}

/*
	Internal functions. DON'T USE lock/unlock into these functions!
*/

// DON'T USE lock/unlock this function
func (this *UserGame) outData(err error) map[string]interface{} {
	errstr := ""
	if err != nil {
		errstr = fmt.Sprintf("%s", err)
	}

	return map[string]interface{}{
		"wins":         this.wins,
		"losings":      this.losings,
		"view_word":    strings.Join(this.view_word, ""),
		"last_word":    strings.Join(this.last_word, ""),
		"attempt":      this.attempt,
		"last_attempt": this.last_attempt,
		"last_letter":  this.last_letter,
		"state":        this.last_state,
		"error":        errstr,
		"last_letters": strings.Join(this.last_letters, ", "),
	}
}

// DON'T USE lock/unlock this function
func (this *UserGame) UpTime() {
	this.exp = time.Now().Add(LiveTimeSec * time.Second)
}

// DON'T USE lock/unlock this function
func (this *UserGame) checkWin() {
	for i := range this.view_word {
		if this.view_word[i] != this.word[i] {
			return
		}
	}
	// If all ok, user win!
	this.win()
}

// DON'T USE lock/unlock this function
func (this *UserGame) clean() {
	this.last_word = this.word
	this.word = []string{}
	this.check_letters = map[string]bool{}
	this.last_state = ""
}

// DON'T USE lock/unlock this function
func (this *UserGame) win() {
	this.clean()
	this.wins++
	this.last_state = "winner"
}

// DON'T USE lock/unlock this function
func (this *UserGame) defeat() {
	this.clean()
	this.losings++
	this.last_state = "loser"
}

func (this *UserGame) IsPlayNow() bool {
	return this.last_state == "is_play_now"
}
