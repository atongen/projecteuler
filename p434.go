package main

/*
http://projecteuler.net/problem=434

Recall that a graph is a collection of vertices and edges connecting the vertices, and that two vertices connected by an edge are called adjacent.
Graphs can be embedded in Euclidean space by associating each vertex with a point in the Euclidean space.
A flexible graph is an embedding of a graph where it is possible to move one or more vertices continuously so that the distance between at least two nonadjacent vertices is altered while the distances between each pair of adjacent vertices is kept constant.
A rigid graph is an embedding of a graph which is not flexible.
Informally, a graph is rigid if by replacing the vertices with fully rotating hinges and the edges with rods that are unbending and inelastic, no parts of the graph can be moved independently from the rest of the graph.

The grid graphs embedded in the Euclidean plane are not rigid, as the following animation demonstrates:

However, one can make them rigid by adding diagonal edges to the cells. For example, for the 2x3 grid graph, there are 19 ways to make the graph rigid:

Note that for the purposes of this problem, we do not consider changing the orientation of a diagonal edge or adding both diagonal edges to a cell as a different way of making a grid graph rigid.

Let R(m,n) be the number of ways to make the m  n grid graph rigid.
E.g. R(2,3) = 19 and R(5,5) = 23679901

Define S(N) as R(i,j) for 1  i, j  N.
E.g. S(5) = 25021721.
Find S(100), give your answer modulo 1000000033

http://www.cs.sunysb.edu/~jgao/CSE590-fall05/notes/lecture3.pdf
http://www.calstatela.edu/faculty/sheubac/papers/Flopgrid.PDF
http://users.wpi.edu/~bservat/umap.pdf
http://en.wikipedia.org/wiki/Kirchhoff%27s_theorem
http://www.wpi.edu/Pubs/ETD/Available/etd-050106-132458/unrestricted/thesis7.pdf
*/

import (
  "fmt"
  "strconv"
  "math"
  matrix "github.com/skelterjohn/go.matrix"
)

type GridKey struct {
  m, n int
}
var memo = make(map[GridKey]int)

func p434naiveCountRigid(m int, n int) int {
  if m > n {
    // swap
    m, n = n, m
  }
  key := GridKey{m,n}
  if cache, found := memo[key]; found {
    fmt.Printf("%dx%d: %d\n", m,n,cache)
    return cache
  }

  vertices := make(map[int]int)

  data := make([]float64, m*n)
  grid := matrix.MakeDenseMatrix(data, m, n)
  num := 0
  var min int = m+n-1
  var max int = int(math.Pow(float64(2), float64(m*n)))
  for i := min; i < max; i++ {
    //if i % 10000 == 0 {
    //  fmt.Println(i)
    //}
    val := strconv.FormatInt(int64(i), 2)
    num_vert := 0
    for idx, char := range val {
      if char == '0' {
        data[idx] = 0
      } else {
        data[idx] = 1
        num_vert += 1
      }
    }
    //fmt.Println(grid)
    if p434NaiveIsGridRigid(grid) {
      vertices[num_vert] += 1
      num++
      if num_vert == min {
        fmt.Println(grid)
      }
    }
  }
  fmt.Println(vertices)
  memo[key] = num
  fmt.Printf("%dx%d: %d\n", m,n,num)
  return num
}

func p434NaiveIsGridRigid(grid matrix.Matrix) bool {
  //m, n := grid.Rows(), grid.Cols()
  //if m > n {
  //  m, n = n, m
  //}
  m := grid.Rows()
  trans := matrix.Transpose(grid)
  prod := matrix.Product(grid, trans)
  result := prod
  for i := 0; i < 2*(m-1); i++ {
    result = matrix.Product(result, prod)
  }
  result = matrix.Product(result, grid)
  //fmt.Println(result)
  for i := 0; i < result.NumElements(); i++ {
    if result.Array()[i] == 0 {
      return false
    }
  }
  return true
}

func p434countRigid(m int, n int) int {
  num := 0
  min := m+n-1
  max := m*n
  for i := min; i <= max; i++ {
    if i == max {
      // all diagonals braced
      num += 1
    } else if i == max - 1 {
      // one diagonal not braced
      num += max
    } else if i == min {
      num += p434countMinCombo(m, n)
    } else {
      // TODO
      num += 1
    }
  }
  return num
}

func p434countMinCombo(m int, n int) int {
  square := m+n
  squ_len := square*square
  adj_data := make([]float64, squ_len)
  d_data := make([]float64, squ_len)
  for i := 0; i < square; i++ {
    n := 0
    for j:= 0; j < square; j++ {
      if (!((i < m && j < m) ||
            (i >= m && j >= m))) {
        adj_data[i*square+j] = 1
        n += 1
      }
    }
    d_data[i*square+i] = float64(n)
  }
  adj := matrix.MakeDenseMatrix(adj_data, square, square)
  fmt.Println(adj)

  d := matrix.MakeDenseMatrix(d_data, square, square)
  fmt.Println(d)

  c := matrix.Difference(d,adj)
  fmt.Println(c)

  sub := c.GetMatrix(1, 1, square-1, square-1)
  fmt.Println(sub)

  return int(sub.Det())
}

func p434sumRigidCombos(n int) int {
  sum := 0
  for i := 1; i <= n; i++ {
    for j := 1; j <= n; j++ {
      //sum += p434countRigid(i, j)
      sum += p434naiveCountRigid(i, j)
    }
  }
  return sum
}

func runP434() {
  //p434sumRigidCombos(4)
  //p434naiveCountRigid(2,3)
  //p434naiveCountRigid(4,4)
  fmt.Println(p434countMinCombo(4,4))
}
