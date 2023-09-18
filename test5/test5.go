package exercices

import (
  "sort"
  "testing"
  "reflect"
  "fmt"
)

type trie struct {
  words [3]int
  wordsLen int
  children map[string]*trie
}

func newTrie () *trie {
  var t trie
  t.children = make(map[string]*trie)
  return &t
}

func (t *trie) insertWord (word string, index int, pos int) int{
  if t.wordsLen < 3 {
    t.words[t.wordsLen] = index
    t.wordsLen += 1
  }

  if pos == len(word) {
    return 0
  }

  ch := string(word[pos])
  if t.children[ch] == nil {
    t.children[ch] = newTrie()
  }
  t.children[ch].insertWord(word, index, pos+1)
  return 0
}

func suggestedProducts(products []string, searchWord string) [][]string {
  sort.Strings(products)

  root := *newTrie()
  for i, word := range products {
    root.insertWord(word, i, 0)
  }

  var ret [][]string
  mismatched := false
  t := root
  for _, ch := range searchWord {
    ch := string(ch)
    if t.children[ch] == nil || mismatched {
      list := make([]string, 0)
      ret = append(ret, list)
      mismatched = true
      continue
    }

    t = *t.children[ch]
    var list []string
    var i int
    for i=0; i<t.wordsLen; i+=1 {
      list = append(list, products[t.words[i]])
    }
    ret = append(ret, list)
  }
  return ret
}

func TestSuggestedProducts(t *testing.T) {
  products := []string{"mobile","mouse","moneypot","monitor","mousepad"}
  searchWord := "mouse"
  output := [][]string{
    {"mobile","moneypot","monitor"}, 
    {"mobile","moneypot","monitor"}, 
    {"mouse","mousepad"}, 
    {"mouse","mousepad"}, 
    {"mouse","mousepad"},
  }
  fmt.Println(output)
  if !reflect.DeepEqual(suggestedProducts(products, searchWord), output) {
    t.Fatalf("Failed testcase #1")
  }

  products = []string{"havana"}
  searchWord = "havana"
  output = [][]string{
    {"havana"}, 
    {"havana"}, 
    {"havana"}, 
    {"havana"}, 
    {"havana"}, 
    {"havana"},
  }
  fmt.Println(output)
  if !reflect.DeepEqual(suggestedProducts(products, searchWord), output) {
    t.Fatalf("Failed testcase #2")
  }

  products = []string{"bags","baggage","banner","box","cloths"}
  searchWord = "bags"
  output = [][]string{
    {"baggage","bags","banner"}, 
    {"baggage","bags","banner"}, 
    {"baggage","bags"}, 
    {"bags"},
  }
  fmt.Println(output)
  if !reflect.DeepEqual(suggestedProducts(products, searchWord), output) {
    t.Fatalf("Failed testcase #3")
  }
}