package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("No input files passed!")
		return
	}

	for _, filePath := range os.Args[1:] {
		variablesStack := NewStack()
		depth := 0

		inputFile, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Input file \"%s\" can't be opened!\n", filePath)
			continue
		}
		defer inputFile.Close()

		scanner := bufio.NewScanner(inputFile)
		scanner.Split(bufio.ScanRunes)

		var buf strings.Builder

		for scanner.Scan() {
			c := scanner.Text()

			switch c {
			case "{":
				depth++
			case "}", ";":
				if c == "}" {
					for {
						item, err := variablesStack.Peek()
						if err != nil {
							break
						}
						if depth != item.Depth {
							break
						}
						_, _ = variablesStack.Pop()
					}
					depth--
				}

				if c == ";" {
					line := buf.String()
					parts := strings.Split(line, "=")
					if len(parts) == 2 {
						variableName := strings.TrimSpace(parts[0])
						variableValueStr := strings.TrimSpace(parts[1])
						variableValue, err := strconv.Atoi(variableValueStr)
						if err != nil {
							fmt.Println("Error parsing variable value:", variableValueStr)
							continue
						}
						variableInfo := codeBlockVariableInfo{
							Info:  variableInfo{Name: variableName, Value: variableValue},
							Depth: depth,
						}
						_ = variablesStack.Push(variableInfo)
					} else if len(parts) == 1 {
						fmt.Println("Showing variables:")
						toPrint := variablesStack
						for {
							item, err := toPrint.Pop()
							if err != nil {
								break
							}
							fmt.Printf("\tName: \"%s\", Value: %d\n", item.Info.Name, item.Info.Value)
						}
					}
					buf.Reset()
				}
			case "=":
				buf.WriteString(c)
			default:
				if strings.TrimSpace(c) != "" {
					buf.WriteString(c)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
	}
}
