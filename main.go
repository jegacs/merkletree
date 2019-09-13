package main

import (
	"crypto/sha256"
	"fmt"
)

type Node struct {
	Label string
	Left  *Node
	Right *Node
}

func main() {
	messages := [][]byte{[]byte("uno"), []byte("dos"), []byte("tres"), []byte("cuatro")}

	for _, message := range messages {
		fmt.Printf("%x\n", sha256.Sum256(message))
	}

	leafs := []Node{
		Node{
			Label: fmt.Sprintf("%x", sha256.Sum256(messages[0])),
		},
		Node{
			Label: fmt.Sprintf("%x", sha256.Sum256(messages[1])),
		},
		Node{
			Label: fmt.Sprintf("%x", sha256.Sum256(messages[2])),
		},
		Node{
			Label: fmt.Sprintf("%x", sha256.Sum256(messages[3])),
		},
	}

	parents := []Node{
		Node{
			Label: fmt.Sprintf("%x", sha256.Sum256([]byte(leafs[0].Label+leafs[1].Label))),
			Left:  &leafs[0],
			Right: &leafs[1],
		},
		Node{
			Label: fmt.Sprintf("%x", sha256.Sum256([]byte(leafs[2].Label+leafs[3].Label))),
			Left:  &leafs[2],
			Right: &leafs[3],
		},
	}

	fmt.Println(parents)

	root := Node{
		Label: fmt.Sprintf("%x", sha256.Sum256([]byte(parents[0].Label+parents[1].Label))),
		Left:  &parents[0],
		Right: &parents[1],
	}

	fmt.Println("root ", root)

}
