# hangman-go

1) Install GO (golang) from https://golang.org/doc/install

2) Run console and make command
```
git clone https://github.com/iostrovok/hangman-go.git
cd ./hangman-go/
make test
make
```

3) Open the http://localhost:19720 into your browser

## Small description
Game logic is realized in package Game. Package Game contents 2 work files and *_test.go file. 
```
rnd_string.go	- internal function:
game.go - game logic
game_test.go - tests
```

### Core game logic
There is a Game object in "game.go". It’s a state machine which has 4 states: “”, “is_play_now”, “winner”, “loser”. “” is default state.

Existing events / Current State -> Input (function) -> Next State /:
```
“loser” -> NewWord -> “is_play_now”
“winner” -> NewWord -> “is_play_now”
“” -> NewWord -> “is_play_now”
“is_play_now” -> Move -> “is_play_now”
“is_play_now” -> Move -> “winner” // player found all letters
“is_play_now” -> Move -> “loser” // count of attempts > 9
```

Each input (Move, NewWord) returns a data, which is returned to player.

#### Function Move
func (this *Game) Move(s string) map[string]interface{}
```
1)	Checks input data “s”. It has to be /[a-zA-Z]/ symbol.
2)	Checks game's state, it has to be “is_play_now”.
3)	Checks duplicate data “s”.
4)	Find symbol in guessed word.
5)	Checks possible events for next states “winner” and “loser” and sets new state if it’s necessary.
```

#### Function NewWord
func (this *Game) NewWord(word string) map[string]interface{}
```
1)	Checks game's state, it has not to be “is_play_now”.
2)	Initialize data for new game.
3)	Sets state to “is_play_now”.
```
