package ds

import . "gopkg.in/check.v1"

func (s *DSTestSuite) TestBTreePut(c *C) {
	var b BTree
	b.Put("abc", "hello world!!")
	b.Put("abcd", "cats are cool")
	b.Put("user_123", "Jeff Goldbloom")
	b.Put("112", "Rogue Dead Guy Ale")
	b.Put("user_421", "Thor")
	b.Put("user_671", "Odin")
	b.Put("harry", "you're a wizard")
	b.Put("star", "wars")

	// b.Put("10", "a")
	// b.Put("20", "b")
	// b.Put("30", "c")
	// b.Put("40", "d")
	// b.Put("50", "e")
	// b.Put("15", "f")
	// b.Put("25", "g")
	// b.Put("18", "h")
	// b.Put("45", "i")

	b.Debug()
}
