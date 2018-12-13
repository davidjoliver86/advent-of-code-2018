package main

import (
	"fmt"
	"sort"
)

type Plan map[rune]map[rune]byte
type Queue []rune

type Instruction struct {
	Step      rune
	DependsOn rune
}

func CreatePlan(instructions []Instruction, max rune) Plan {
	dependencies := make(map[rune]map[rune]byte)
	for r := 'A'; r <= max; r++ {
		dependencies[r] = make(map[rune]byte)
	}
	for _, instruction := range instructions {
		dependencies[instruction.Step][instruction.DependsOn] = 1
	}
	return dependencies
}

func (q *Queue) Add(p Plan) {
	for step, dependencies := range p {
		if len(dependencies) == 0 {
			*q = append(*q, step)
		}
	}
}

func (q *Queue) Remove() rune {
	r := (*q)[0]
	(*q) = (*q)[1:]
	return r
}

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i] < q[j]
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) WorkCycle(p Plan) string {
	complete := make([]rune, 0)

	// sort queue
	sort.Sort(q)

	// pick off lowest letter
	step := q.Remove()

	// if that step's dependencies are empty > delete from plan, add to return value
	if len(p[step]) == 0 {
		delete(p, step)
		complete = append(complete, step)
	}

	// else if that step is another step's dependency -> delete from that list
	for _, dependencies := range p {
		_, ok := dependencies[step]
		if ok {
			delete(dependencies, step)
		}
	}
	return string(complete)
}

func main() {
	instructions := []Instruction{
		Instruction{'A', 'C'},
		Instruction{'F', 'C'},
		Instruction{'B', 'A'},
		Instruction{'D', 'A'},
		Instruction{'E', 'B'},
		Instruction{'E', 'D'},
		Instruction{'E', 'F'},
	}
	plan := CreatePlan(instructions, 'F')
	queue := make(Queue, 0)
	queue.Add(plan)
	fmt.Println(queue)
	fmt.Println(plan)
	omg := queue.WorkCycle(plan)
	fmt.Println(omg)
	fmt.Println(queue)
	fmt.Println(plan)
	queue.Add(plan)
	omg2 := queue.WorkCycle(plan)
	fmt.Println(omg2)
	fmt.Println(queue)
	fmt.Println(plan)
	omg3 := queue.WorkCycle(plan)
	fmt.Println(omg3)
	fmt.Println(queue)
	fmt.Println(plan)
}
