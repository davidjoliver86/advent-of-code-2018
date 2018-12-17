package main

import (
	"fmt"
	"sort"

	util "github.com/davidjoliver86/advent-of-code-2018"
)

type Plan map[rune]map[rune]byte
type Queue []rune

type Instruction struct {
	Step      rune
	DependsOn rune
}

func CreatePlan(instructions []Instruction, upto rune) Plan {
	dependencies := make(map[rune]map[rune]byte)
	for r := 'A'; r <= upto; r++ {
		dependencies[r] = make(map[rune]byte)
	}
	for _, instruction := range instructions {
		dependencies[instruction.Step][instruction.DependsOn] = 1
	}
	return dependencies
}

func (p Plan) String() string {
	plan := ""
	for i, charmap := range p {
		plan += fmt.Sprintf("%s: [", string(i))
		for j, _ := range charmap {
			plan += string(j)
		}
		plan += string("] ")
	}
	return plan
}

func (q Queue) String() string {
	letters := ""
	for _, letter := range q {
		letters += string(letter)
	}
	return fmt.Sprintf("[%s]", letters)
}

func (q *Queue) Add(p Plan) {
	for step, dependencies := range p {
		if len(dependencies) == 0 && !q.Find(step) {
			*q = append(*q, step)
		}
	}
}

func (q *Queue) Remove() rune {
	r := (*q)[0]
	(*q) = (*q)[1:]
	return r
}

func (q Queue) Find(ch rune) bool {
	for _, c := range q {
		if ch == c {
			return true
		}
	}
	return false
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

func (q *Queue) WorkCycle(p Plan) Queue {
	complete := make(Queue, 0)

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
	return complete
}

func DoWork(instructions []Instruction) Queue {
	max := instructions[0].Step
	for _, i := range instructions {
		if i.Step > max {
			max = i.Step
		}
		if i.DependsOn > max {
			max = i.DependsOn
		}
	}
	plan := CreatePlan(instructions, max)
	queue := make(Queue, 0)
	complete := make(Queue, 0)
	queue.Add(plan)
	for len(queue) != 0 {
		iteration := queue.WorkCycle(plan)
		for _, ch := range iteration {
			complete = append(complete, ch)
		}
		queue.Add(plan)
	}
	return complete
}

func ParseInstructions(path string) []Instruction {
	instructions := make([]Instruction, 0)
	lines := util.FileLines(path)
	for _, line := range lines {
		instruction := Instruction{rune(line[36]), rune(line[5])}
		instructions = append(instructions, instruction)
	}
	return instructions
}

func main() {
	instructions := ParseInstructions("input.txt")
	fmt.Println(DoWork(instructions))
}
