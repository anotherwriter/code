/**
https://leetcode.com/problems/min-stack/
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:
 * MinStack() initializes the stack object.
 * void push(int val) pushes the element val onto the stack.
 * void pop() removes the element on the top of the stack.
 * int top() gets the top element of the stack.
 * int getMin() retrieves the minimum element in the stack.

You must implement a solution with O(1) time complexity for each function.

Constraints:
 * -2^31 <= val <= 2^31 - 1
 * Methods pop, top and getMin operations will always be called on non-empty stacks.
 * At most 3 * 10^4 calls will be made to push, pop, top, and getMin.
*/
package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // return -3

	minStack.Pop()
	fmt.Println(minStack.Top())    // return 0
	fmt.Println(minStack.GetMin()) // return -2
}

type MinStack struct {
	data      []int
	monoStack []int // monotonous decreasing stack
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)

	if len(this.monoStack) == 0 || this.monoStack[len(this.monoStack)-1] >= val {
		this.monoStack = append(this.monoStack, val)
	}
}

func (this *MinStack) Pop() {
	top := this.Top()
	this.data = this.data[:len(this.data)-1]

	if top == this.monoStack[len(this.monoStack)-1] {
		this.monoStack = this.monoStack[:len(this.monoStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.monoStack[len(this.monoStack)-1]
}
