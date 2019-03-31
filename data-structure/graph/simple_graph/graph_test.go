package simple_graph

import "testing"

func TestSimpleGraph(t *testing.T) {
	g := New()
	if g.V() != 0 {
		t.Error()
	}
	if g.E() != 0 {
		t.Error()
	}

	g.AddEdge("A", "B")
	if g.V() != 2 {
		t.Error()
	}
	if g.E() != 1 {
		t.Error()
	}

	g.AddEdge("B", "C")
	if g.V() != 3 {
		t.Error()
	}
	if g.E() != 2 {
		t.Error()
	}

	g.AddEdge("A", "C")
	if g.V() != 3 {
		t.Error()
	}
	if g.E() != 3 {
		t.Error()
	}

	g.AddEdge("C", "A")
	if g.V() != 3 {
		t.Error()
	}
	if g.E() != 3 {
		t.Error()
	}

	g.AddEdge("C", "C")
	if g.V() != 3 {
		t.Error()
	}

	_, err := g.Adj("A")
	if err != nil {
		t.Error(err)
	}
}
