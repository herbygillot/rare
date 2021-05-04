package expressions

import (
	"fmt"
	"strconv"
)

const (
	ArraySeparator       rune   = '\x00'
	ArraySeparatorString string = string(ArraySeparator)
)

// KeyBuilderFunction defines a helper function at runtime
type KeyBuilderFunction func([]KeyBuilderStage) KeyBuilderStage

// KeyBuilderStage is a stage within the compiled builder
type KeyBuilderStage func(KeyBuilderContext) string

func stageLiteral(s string) KeyBuilderStage {
	return KeyBuilderStage(func(context KeyBuilderContext) string {
		return s
	})
}

func stageSimpleVariable(s string) KeyBuilderStage {
	index, err := strconv.Atoi(s)
	if err != nil {
		return KeyBuilderStage(func(context KeyBuilderContext) string {
			return context.GetKey(s)
		})
	}
	return KeyBuilderStage(func(context KeyBuilderContext) string {
		return context.GetMatch(index)
	})
}

func stageError(msg string) KeyBuilderStage {
	errMessage := fmt.Sprintf("<%s>", msg)
	return KeyBuilderStage(func(context KeyBuilderContext) string {
		return errMessage
	})
}
