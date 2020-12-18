package main

import (
	"aoc"
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func printList(l list.List) {
	f := l.Front()
	for f != nil {
		fmt.Print(f.Value, " ")
		f = f.Next()
	}
	fmt.Println("")
}

func sumInParen(fields []string) int {
	valStack := list.New()
	opStack := list.New()
	for _, field := range fields {
		if field == "+" {
			opStack.PushFront("+")
		} else if field == "*" {
			opStack.PushFront("*")
		} else {
			val, _ := strconv.Atoi(field)
			valStack.PushFront(val)
		}
	}

	for valStack.Len() > 1 {
		val1 := valStack.Remove(valStack.Back()).(int)
		val2 := valStack.Remove(valStack.Back()).(int)
		op := opStack.Remove(opStack.Back())
		if op == "+" {
			valStack.PushBack(val1 + val2)
		} else {
			valStack.PushBack(val1 * val2)
		}
	}
	return valStack.Remove(valStack.Back()).(int)
}

func findParen(fields []string) (list.List, list.List) {
	openParenStack := list.New()
	closeParenStack := list.New()

	for i, field := range fields {
		if field == "(" {
			openParenStack.PushBack(i)
		} else if field == ")" {
			closeParenStack.PushFront(i)
			break
		}
	}
	return *openParenStack, *closeParenStack
}

func part1(lines []string) int {
	sum := 0
	for _, line := range lines {
		newLine := strings.ReplaceAll(line, "(", " ( ")
		newLine = strings.ReplaceAll(newLine, ")", " ) ")
		newLine = "( " + newLine + " )"
		fields := strings.Fields(newLine)
		lineSum := 0

		openParenStack, closeParenStack := findParen(fields)

		for openParenStack.Len() > 0 {
			openIndex := openParenStack.Remove(openParenStack.Back()).(int)
			closeIndex := closeParenStack.Remove(closeParenStack.Back()).(int)
			lineSum = sumInParen(fields[openIndex+1 : closeIndex])
			newFields := fields[:openIndex]
			newFields = append(newFields, strconv.Itoa(lineSum))
			newFields = append(newFields, fields[closeIndex+1:]...)
			fields = newFields
			openParenStack, closeParenStack = findParen(fields)
		}
		sum += lineSum
	}
	return sum
}

func add(fields []string) {
	for i := 1; i < len(fields)-1; i++ {
		if fields[i] == "+" {
			val, err := strconv.Atoi(fields[i-1])
			if err != nil {
				continue
			}
			val2, err := strconv.Atoi(fields[i+1])
			if err != nil {
				continue
			}
			newFields := fields[:i-1]
			newFields = append(newFields, strconv.Itoa(val+val2))
			newFields = append(newFields, fields[i+2:]...)
			fields = newFields
			i = -1
		}
	}
}

func part2(lines []string) uint64 {
	var sum uint64 = 0

	for _, line := range lines {
		newLine := strings.ReplaceAll(line, "(", " ( ")
		newLine = strings.ReplaceAll(newLine, ")", " ) ")
		newLine = "( " + newLine + " )"
		lineSum := 0
		fields := strings.Fields(newLine)
		add(fields)
		openParenStack, closeParenStack := findParen(fields)
		for openParenStack.Len() > 0 {
			openIndex := openParenStack.Remove(openParenStack.Back()).(int)
			closeIndex := closeParenStack.Remove(closeParenStack.Back()).(int)
			lineSum = sumInParen(fields[openIndex+1 : closeIndex])
			newFields := fields[:openIndex]
			newFields = append(newFields, strconv.Itoa(lineSum))
			newFields = append(newFields, fields[closeIndex+1:]...)
			fields = newFields
			add(fields)
			openParenStack, closeParenStack = findParen(fields)
		}
		sum += uint64(lineSum)
	}
	return sum
}

func main() {
	lines := aoc.GetStdin()
	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}
