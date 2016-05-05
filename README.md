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
* ```Quick sort``` (in progress)
* ```Merge sort``` (done)
* ```Heap sort``` (coming soon)
* ```Selection sort``` (coming soon)
* ```Quick sort 3``` (coming soon)
* ```Insertion sort``` (coming soon)
* ```Shell sort``` (coming soon)

### Resources
* https://stevenschmatz.gitbooks.io/data-structures-and-algorithms/
* https://en.wikipedia.org/wiki/Dynamic_programming
