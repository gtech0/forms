package util

import (
	"hedgehog-forms/model/form/pattern"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"slices"
)

func Zero[T any]() T {
	return *new(T)
}

func FindKeyByValue(m map[string]int, value int) (key string, ok bool) {
	for k, v := range m {
		if v == value {
			key = k
			ok = true
			return
		}
	}
	return
}

func ExtractQuestionObjs(formPattern pattern.FormPattern) []question.IQuestion {
	questions := make([]question.IQuestion, 0)
	for _, patternSection := range formPattern.Sections {
		for _, sectionBlock := range patternSection.Blocks {
			switch assertedBlock := sectionBlock.(type) {
			case *block.DynamicBlock:
				questions = slices.Concat(questions, assertedBlock.Questions)
			case *block.StaticBlock:
				variants := assertedBlock.Variants
				for _, variant := range variants {
					questions = slices.Concat(questions, variant.Questions)
				}
			}
		}
	}
	return questions
}
