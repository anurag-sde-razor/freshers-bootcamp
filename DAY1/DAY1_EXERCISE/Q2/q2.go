package main

import (
    "fmt"
)


type Node struct {
    Value string
    Left  *Node
    Right *Node
}


func preorderTraversal(node *Node) {
    if node == nil {
        return
    }
    fmt.Print(node.Value, " ")
    preorderTraversal(node.Left)
    preorderTraversal(node.Right)
}


func postorderTraversal(node *Node) {
    if node == nil {
        return
    }
    postorderTraversal(node.Left)
    postorderTraversal(node.Right)
    fmt.Print(node.Value, " ")
}

func main() {
    
    root := &Node{
        Value: "+",
        Left: &Node{
            Value: "a",
        },
        Right: &Node{
            Value: "-",
            Left: &Node{
                Value: "b",
            },
            Right: &Node{
                Value: "c",
            },
        },
    }

    fmt.Println("Preorder traversal:")
    preorderTraversal(root)
    fmt.Println()

    fmt.Println("Postorder traversal:")
    postorderTraversal(root)
    fmt.Println()
}