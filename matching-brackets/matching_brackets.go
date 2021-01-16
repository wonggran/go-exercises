package brackets

import (
	"errors"
	"regexp"
)

type stack struct {
	arr []byte
}

func (s *stack) push(b byte) {
	s.arr = append(s.arr, b)
}

func (s *stack) emptyCheck() error {
	if len(s.arr) == 0 {
		return errors.New("Stack is empty.")
	}
	return nil
}

func (s *stack) peak() (byte, error) {
	if err := s.emptyCheck(); err != nil {
		return 0, err
	}
	return s.arr[len(s.arr)-1], nil
}

func (s *stack) pop() (byte, error) {
	ret, err := s.peak()
	if err != nil {
		return 0, err
	}
	s.arr = s.arr[:len(s.arr)-1]
	return ret, nil
}

func replaceLatexBrackets(input string) string {
	left := regexp.MustCompile(`\\left`)
	right := regexp.MustCompile(`\\right`)

	byteInput := []byte(input)
	leftParen, rightParen := []byte{'('}, []byte{')'}
	result := left.ReplaceAll(byteInput, leftParen)
	result = right.ReplaceAll(result, rightParen)

	return string(result)
}

func emptyInputCheck(input string) bool {
	if len(input) == 0 {
		return true
	}
	return false
}

func singleRuneInputCheck(input string) bool {
	if len(input) == 1 {
		return true
	}
	return false
}

func Bracket(input string) bool {
	// Push left brackets onto a stack and when a right bracket is encountered check if its corresponding
	// left bracket is on the top of the stack, if it is pop it, if it isn't then we have an invalid matching.
	// If there are brackets left on the stack after the entire input is processed then we have an excess of
	// brackets and therefore it is an invalid input.
	// Replace LaTeX \left and \right with ( ) for easier parsing.
	// The empty input is valid.
	if emptyInputCheck(input) {
		return true
	}

	if singleRuneInputCheck(input) {
		return false
	}

	cleanedInput := replaceLatexBrackets(input)
	byteInput := []byte(cleanedInput)

	leftBrackets := map[byte]bool{
		'(': true,
		'[': true,
		'{': true,
	}

	rightBrackets := map[byte]bool{
		')': true,
		']': true,
		'}': true,
	}

	correspondingLeft := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	s := &stack{}
	for _, b := range byteInput {
		if isLeft := leftBrackets[b]; isLeft {
			s.push(b)
		} else if isRight := rightBrackets[b]; isRight {
			if peakedByte, _ := s.peak(); peakedByte == correspondingLeft[b] {
				s.pop()
			} else {
				return false
			}
		} else {
			continue
		}
	}

	if err := s.emptyCheck(); err != nil {
		return true
	}

	return false
}
