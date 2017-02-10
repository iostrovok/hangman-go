package Game

/*

"Game struct" is a state machine which has 4 states: “”, “is_play_now”, “winner”, “loser”.
“” is default state.

Existing events:

“loser” -> NewWord -> “is_play_now”
“winner” -> NewWord -> “is_play_now”
“” -> NewWord -> “is_play_now”
“is_play_now” -> Move -> “is_play_now”
“is_play_now” -> Move -> “winner” // player found all letters
“is_play_now” -> Move -> “loser” // count of attempts > 9

Each input (Move, NewWord) returns a data, which is returned to player.

*/

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"
)

var clenerReg = regexp.MustCompile(`[^a-zA-Z]`)

const CountAttempt = 10
const KeyLength = 10
const LiveTimeSec = 60 * 60 * 24 // 1 day

type Game struct {
	sync.RWMutex

	// currently state. Can be “”, “is_play_now”, “winner”, “loser”
	last_state string

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
}

// Constructor
func New() *Game {
	id := RandString(KeyLength)

	user := &Game{
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
	Interface functions: Info(), NewWord(word string), Move(s string).
	HAVE TO USE lock/unlock into these functions.
	Only interface functions can use lock/unlock other side we get a race condition.
*/
// Info() just returns user info.
func (this *Game) Info() map[string]interface{} {
	this.Lock()
	defer this.Unlock()
	this.UpTime() // Player is active, give him more time for session.
	return this.outData(nil)
}

// Starts new game
func (this *Game) NewWord(word string) map[string]interface{} {
	this.Lock()
	defer this.Unlock()
	this.UpTime() // Player is active, give him more time for session.

	// check that player state IS NOT "is_play_now”
	if this.isPlayNow() {
		// TODO a cheater or old cookie
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

func (this *Game) Move(s string) map[string]interface{} {
	this.Lock()
	defer this.Unlock()
	this.UpTime() // Player is active, give him more time for session.

	// ----> check params start and conditions
	this.last_letter = clenerReg.ReplaceAllString(strings.ToLower(s), "")

	// check that player state IS "is_play_now”
	if !this.isPlayNow() {
		// TODO a cheater or old cookie
		return this.outData(errors.New("not_in_the_game"))
	}

	// check correct params.
	// We also may check it in index.go.
	if len(this.last_letter) != 1 {
		return this.outData(errors.New("bad_move"))
	}

	// check duplicate letter
	if this.check_letters[this.last_letter] {
		return this.outData(errors.New("duplicate"))
	}
	this.check_letters[this.last_letter] = true

	// <---- check params start and conditions finish

	// Fix moves data for drawing UI
	this.last_letters = append(this.last_letters, this.last_letter)

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

// Exp() returns when the game is expires by timeout
func (this *Game) Exp() time.Time {
	this.RLock()
	defer this.RUnlock()
	return this.exp
}

/*
	Internal functions. DON'T USE lock/unlock into these functions!
*/

// DON'T USE lock/unlock this function
func (this *Game) outData(err error) map[string]interface{} {
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
		"last_letters": this.last_letters,
	}
}

// DON'T USE lock/unlock this function
func (this *Game) UpTime() {
	this.exp = time.Now().Add(LiveTimeSec * time.Second)
}

// DON'T USE lock/unlock this function
func (this *Game) checkWin() {
	for i := range this.view_word {
		if this.view_word[i] != this.word[i] {
			return
		}
	}
	// If all ok, user win!
	this.win()
}

// DON'T USE lock/unlock this function
func (this *Game) clean() {
	this.last_word = this.word
	this.word = []string{}
	this.check_letters = map[string]bool{}
	this.last_state = ""
}

// DON'T USE lock/unlock this function
func (this *Game) win() {
	this.clean()
	this.wins++
	this.last_state = "winner"
}

// DON'T USE lock/unlock this function
func (this *Game) defeat() {
	this.clean()
	this.losings++
	this.last_state = "loser"
}

func (this *Game) isPlayNow() bool {
	return this.last_state == "is_play_now"
}
