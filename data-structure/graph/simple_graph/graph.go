/*
Copyright 2017 Julian Griggs
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package simple_graph

import (
	"errors"
	"sync"
)

var (
	ErrVertexNotFound = errors.New("vertex not found")

	ErrSelfLoop = errors.New("self loops not permitted")

	ErrParallelEdge = errors.New("parallel edges are not permitted")
)

//v - number of vertices  e - number of edges
type Graph struct {
	mutex         sync.RWMutex
	adjacencyList map[interface{}]map[interface{}]struct{}
	v, e          int
}

func (g *Graph) V() int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return g.v
}

func (g *Graph) E() int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return g.e
}

func (g *Graph) AddEdge(v, w interface{}) error {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if v == w {
		return ErrSelfLoop
	}

	g.addVertex(v)
	g.addVertex(w)

	if _, ok := g.adjacencyList[v][w]; ok {
		return ErrParallelEdge
	}

	g.adjacencyList[v][w] = struct{}{}
	g.adjacencyList[w][v] = struct{}{}
	g.e++
	return nil
}

func (g *Graph) Adj(v interface{}) ([]interface{}, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	deg, err := g.Degree(v)
	if err != nil {
		return nil, ErrVertexNotFound
	}

	adj := make([]interface{}, deg)
	i := 0
	for key := range g.adjacencyList[v] {
		adj[i] = key
		i++
	}
	return adj, nil
}

func (g *Graph) Degree(v interface{}) (int, error) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	val, ok := g.adjacencyList[v]
	if !ok {
		return 0, ErrVertexNotFound
	}
	return len(val), nil
}

func (g *Graph) addVertex(v interface{}) {
	mm, ok := g.adjacencyList[v]
	if !ok {
		mm = make(map[interface{}]struct{})
		g.adjacencyList[v] = mm
		g.v++
	}
}

func New() *Graph {
	return &Graph{
		adjacencyList: make(map[interface{}]map[interface{}]struct{}),
		v:             0,
		e:             0,
	}
}
