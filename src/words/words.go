package Words

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"sync"
)

var clenerReg = regexp.MustCompile(`[^a-zA-Z]`)

type Words struct {
	sync.RWMutex
	data []string
}

var once sync.Once
var list *Words

func Init(file string) *Words {
	onceBody := func() {
		fmt.Printf("Words file: %s\n", file)

		wordList, err := loadFile(file)
		if err != nil {
			log.Fatalf("Error words.txt file[%s]: %s\n", file, err)
			os.Exit(1)
		}

		fmt.Printf("Words count: %s\n", len(wordList))

		list = &Words{
			data: wordList,
		}
	}
	once.Do(onceBody)

	return list
}

func Get() string {
	return list.Get()
}

func (this *Words) Get() string {
	return this.data[rand.Intn(len(this.data))]
}

func loadFile(file string) ([]string, error) {
	words, err := ioutil.ReadFile(file) // just pass the file name
	if err != nil {
		return nil, err
	}

	out := []string{}
	uniq := map[string]bool{}
	list := strings.Split(string(words), "\n")
	for _, w := range list {
		w := clenerReg.ReplaceAllString(strings.ToLower(w), "")
		if len(w) > 2 && !uniq[w] {
			out = append(out, w)
			uniq[w] = true
		}
	}

	if len(out) == 0 {
		err = errors.New("File is empty")
	}

	return out, err
}
