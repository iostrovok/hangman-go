package Game

import (
	// "fmt"
	. "gopkg.in/check.v1"
	"testing"
)

/*
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
*/

func TestUser(t *testing.T) {
	TestingT(t)
}

type TestsSuite struct{}

var _ = Suite(&TestsSuite{})

func (s TestsSuite) Test_New(c *C) {
	//c.Skip("Not now")
	mygame := New()
	c.Assert(mygame, NotNil)
	c.Assert(mygame.last_state, Not(Equals), "is_play_now")
}

func (s TestsSuite) Test_checkID(c *C) {
	//c.Skip("Not now")
	mygame1 := New()
	mygame2 := New()
	c.Assert(mygame1.ID, Not(Equals), mygame2.ID)
}

func (s TestsSuite) Test_NewWord_1(c *C) {
	//c.Skip("Not now")
	word := "abc"
	mygame := New()
	mygame.NewWord(word)

	c.Assert(mygame.word, DeepEquals, []string{"a", "b", "c"})
	c.Assert(mygame.view_word, DeepEquals, []string{"*", "*", "*"})

	c.Assert(mygame.last_state, Equals, "is_play_now")
}

func (s TestsSuite) Test_Move_1(c *C) {
	//c.Skip("Not now")
	word := "abc"
	mygame := New()
	mygame.NewWord(word)
	res := mygame.Move("a")
	c.Assert(mygame.view_word, DeepEquals, []string{"a", "*", "*"})

	// Check return
	c.Assert(res["state"], Equals, "is_play_now")
	c.Assert(res["view_word"], DeepEquals, "a**")
}

// “loser” -> NewWord -> “is_play_now”
func (s TestsSuite) Test_loser_NewWord_is_play_now(c *C) {
	//c.Skip("Not now")
	word := "abc"
	mygame := New()
	mygame.NewWord(word)

	mygame.last_state = "loser"
	c.Assert(mygame.last_state, Equals, "loser")

	mygame.NewWord(word)
	c.Assert(mygame.word, DeepEquals, []string{"a", "b", "c"})
	c.Assert(mygame.last_state, Equals, "is_play_now")
}

// “winner” -> NewWord -> “is_play_now”
func (s TestsSuite) Test_winner_NewWord_is_play_now(c *C) {
	//c.Skip("Not now")
	word := "abc"
	mygame := New()
	mygame.NewWord("abcde")

	mygame.last_state = "winner"
	c.Assert(mygame.last_state, Equals, "winner")

	mygame.NewWord(word)
	c.Assert(mygame.word, DeepEquals, []string{"a", "b", "c"})
	c.Assert(mygame.last_state, Equals, "is_play_now")
}

// “” -> NewWord -> “is_play_now”
func (s TestsSuite) Test__NewWord_is_play_now(c *C) {
	//c.Skip("Not now")
	word := "abc"
	mygame := New()

	c.Assert(mygame.last_state, Equals, "")

	mygame.NewWord(word)
	c.Assert(mygame.word, DeepEquals, []string{"a", "b", "c"})
	c.Assert(mygame.last_state, Equals, "is_play_now")
}

// “is_play_now” -> Move -> “is_play_now”
func (s TestsSuite) Test_is_play_now_Move_is_play_now(c *C) {
	//c.Skip("Not now")
	word := "abc"
	mygame := New()
	mygame.NewWord(word)

	c.Assert(mygame.last_state, Equals, "is_play_now")

	mygame.Move("a")
	c.Assert(mygame.last_state, Equals, "is_play_now")
}

// “is_play_now” -> Move -> “winner”
func (s TestsSuite) Test_is_play_now_Move_winner(c *C) {
	//c.Skip("Not now")
	word := "abc"
	mygame := New()
	mygame.NewWord(word)

	c.Assert(mygame.last_state, Equals, "is_play_now")
	mygame.view_word = []string{"a", "b", "*"}

	c.Assert(mygame.attempt < 10, Equals, true)

	mygame.Move("c")
	c.Assert(mygame.last_state, Equals, "winner")
}

// “is_play_now” -> Move -> “loser”
func (s TestsSuite) Test_is_play_now_Move_loser(c *C) {
	//c.Skip("Not now")
	word := "abc"
	mygame := New()
	mygame.NewWord(word)

	c.Assert(mygame.last_state, Equals, "is_play_now")
	mygame.view_word = []string{"*", "*", "*"}
	mygame.attempt = 9

	mygame.Move("t")
	c.Assert(mygame.last_state, Equals, "loser")
}

func (s TestsSuite) Test_correct_loser(c *C) {
	//c.Skip("Not now")
	word := "abc"
	mygame := New()
	mygame.NewWord(word)

	// 9 bad attemps and 2 goods
	attempts := []string{"q", "w", "e", "r", "t", "y", "a", "f", "i", "o", "b"}
	for _, letter := range attempts {
		mygame.Move(letter)
		c.Assert(mygame.last_state, Equals, "is_play_now")
	}

	mygame.Move("z")
	c.Assert(mygame.last_state, Equals, "loser")
}

func (s TestsSuite) Test_correct_winner(c *C) {
	//c.Skip("Not now")
	word := "abc"
	mygame := New()
	mygame.NewWord(word)

	// 9 bad attemps and 2 goods
	attempts := []string{"q", "w", "e", "r", "t", "y", "a", "f", "i", "o", "b"}
	for _, letter := range attempts {
		mygame.Move(letter)
		c.Assert(mygame.last_state, Equals, "is_play_now")
	}

	// Success move
	mygame.Move("c")
	c.Assert(mygame.last_state, Equals, "winner")
}

func (s TestsSuite) Test_correct_winner_fast(c *C) {
	//c.Skip("Not now")
	word := "abc"
	mygame := New()
	mygame.NewWord(word)

	// Success move
	mygame.Move("a")
	c.Assert(mygame.last_state, Equals, "is_play_now")
	mygame.Move("b")
	c.Assert(mygame.last_state, Equals, "is_play_now")
	mygame.Move("c")
	c.Assert(mygame.last_state, Equals, "winner")
}
