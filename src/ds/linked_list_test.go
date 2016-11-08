package ds

import . "gopkg.in/check.v1"

func (s *DSTestSuite) TestLinkedListInsert(c *C) {
	list := LinkedList{}
	n1 := LLNode{Data: "hello world"}
	n2 := LLNode{Data: "here be dragons"}
	n3 := LLNode{Data: "I cast magic missle!"}

	list.Insert(&n1)
	list.Insert(&n2)
	list.Insert(&n3)

	c.Assert(list.Root, Equals, &n1)
	c.Assert(list.Root.Next, Equals, &n2)
	c.Assert(list.Root.Next.Next, Equals, &n3)
}

func (s *DSTestSuite) TestLinkedListInsertAtSimple(c *C) {
	list := LinkedList{}
	n1 := LLNode{Data: "hello world"}
	n2 := LLNode{Data: "here be dragons"}

	// n3 will start as the root node
	list.Insert(&n2)
	c.Assert(list.Root, Equals, &n2)

	// now the root should become n1
	list.InsertAt(&n1, &n2, true)

	c.Assert(list.Root, Equals, &n1)
	c.Assert(list.Root.Next, Equals, &n2)
}

func (s *DSTestSuite) TestLinkedListInsertAtFull(c *C) {
	list := LinkedList{}
	n1 := LLNode{Data: "hello world"}
	n2 := LLNode{Data: "here be dragons"}
	n3 := LLNode{Data: "I cast magic missle!"}
	n4 := LLNode{Data: "I don't say blah blah blah!"}

	// n3 will start as the root node
	list.Insert(&n3)
	c.Assert(list.Root, Equals, &n3)

	// now the root should become n1
	list.InsertAt(&n1, &n3, true)

	// insert n4 after n3
	list.InsertAt(&n4, &n3, false)

	// insert n2 before n3
	list.InsertAt(&n2, &n3, true)

	c.Assert(list.Root, Equals, &n1)
	c.Assert(list.Root.Next, Equals, &n2)
	c.Assert(list.Root.Next.Next, Equals, &n3)
	c.Assert(list.Root.Next.Next.Next, Equals, &n4)
}

func (s *DSTestSuite) TestLinkedListRemoveRoot(c *C) {
	list := LinkedList{}
	n1 := LLNode{Data: "hello world"}
	n2 := LLNode{Data: "here be dragons"}
	n3 := LLNode{Data: "I cast magic missle!"}

	// insert n1 (root)
	list.Insert(&n1)

	// insert n2
	list.Insert(&n2)

	// insert n3
	list.Insert(&n3)

	// remove the root
	list.Remove(&n1)

	c.Assert(list.Root, Equals, &n2)
	c.Assert(list.Root.Next, Equals, &n3)
	c.Assert(list.Root.Next.Next, IsNil)
}

func (s *DSTestSuite) TestLinkedListRemoveMiddle(c *C) {
	list := LinkedList{}
	n1 := LLNode{Data: "hello world"}
	n2 := LLNode{Data: "here be dragons"}
	n3 := LLNode{Data: "I cast magic missle!"}
	n4 := LLNode{Data: "I don't say blah blah blah!"}

	// insert n1
	list.Insert(&n1)

	// insert n2
	list.Insert(&n2)

	// insert n3
	list.Insert(&n3)

	// insert n4
	list.Insert(&n4)

	// now remove n2
	list.Remove(&n2)

	c.Assert(list.Root, Equals, &n1)
	c.Assert(list.Root.Next, Equals, &n3)
	c.Assert(list.Root.Next.Next, Equals, &n4)

	// make extra sure nothing was accidentally overwritten giving us a false positive...
	c.Assert(list.Root.Next.Next.Data, Equals, "I don't say blah blah blah!")
}

func (s *DSTestSuite) TestLinkedListRemoveLast(c *C) {
	list := LinkedList{}
	n1 := LLNode{Data: "hello world"}
	n2 := LLNode{Data: "here be dragons"}
	n3 := LLNode{Data: "I cast magic missle!"}

	// insert n1 (root)
	list.Insert(&n1)

	// insert n2
	list.Insert(&n2)

	// insert n3
	list.Insert(&n3)

	// remove the root
	list.Remove(&n3)

	c.Assert(list.Root, Equals, &n1)
	c.Assert(list.Root.Next, Equals, &n2)
	c.Assert(list.Root.Next.Next, IsNil)
}

func (s *DSTestSuite) TestLinkedListFind(c *C) {
	list := LinkedList{}
	n1 := LLNode{Data: "hello world"}
	n2 := LLNode{Data: "here be dragons"}
	n3 := LLNode{Data: "I cast magic missle!"}
	n4 := LLNode{Data: "I don't say blah blah blah!"}

	list.Insert(&n1)
	list.Insert(&n2)
	list.Insert(&n3)
	list.Insert(&n4)

	r := list.Find(&n3)
	c.Assert(r.Data, Equals, n3.Data)

	r = list.Find(&n1)
	c.Assert(r.Data, Equals, n1.Data)

	r = list.Find(&n4)
	c.Assert(r.Data, Equals, n4.Data)
}

func (s *DSTestSuite) TestLinkedListFindByValue(c *C) {
	list := LinkedList{}
	n1 := LLNode{Data: "hello world"}
	n2 := LLNode{Data: "here be dragons"}
	n3 := LLNode{Data: "I cast magic missle!"}
	n4 := LLNode{Data: "I don't say blah blah blah!"}

	list.Insert(&n1)
	list.Insert(&n2)
	list.Insert(&n3)
	list.Insert(&n4)

	r := list.FindByValue(n3.Data)
	c.Assert(r.Data, Equals, n3.Data)

	r = list.FindByValue(n1.Data)
	c.Assert(r.Data, Equals, n1.Data)

	r = list.FindByValue(n4.Data)
	c.Assert(r.Data, Equals, n4.Data)
}
