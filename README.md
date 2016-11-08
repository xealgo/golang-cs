# golang-cs
An educational project for practicing algorithms and data structures in Golang.

### Testing
Tests are written using the Gocheck library which can be found at
https://labix.org/gocheck. The following example will run all of the sorting tests and run the benchmarks for each.

```
go test ./src/sorting/ -check.f SortingTestSuite -v -check.b 1000
```

### Sorting
* ```Bubble sort``` (done)
* ```Merge sort``` (done)
* ```Quick sort``` (in progress)
* ```Heap sort``` (coming soon)
* ```Selection sort``` (coming soon)
* ```Quick sort 3``` (coming soon)
* ```Insertion sort``` (coming soon)
* ```Shell sort``` (coming soon)

### Data Structures
* ```Self-balancing Binary Tree``` (in progress)
* ```Trie``` (coming soon)
* ```Linked List``` (done)
* ```Double Linked List``` (coming soon)
* ```Stack``` (coming soon)
* ```Queue``` (coming soon)
* ```Dequeue``` (coming soon)
* ```Red Black Tree``` (near future)

### Resources
* https://stevenschmatz.gitbooks.io/data-structures-and-algorithms/
* https://en.wikipedia.org/wiki/Dynamic_programming
* https://en.wikipedia.org/wiki/List_of_terms_relating_to_algorithms_and_data_structures
* http://blog.hackerearth.com/2015/05/top-7-algorithms-and-data-structures-every-programmer-should-know-about.html
