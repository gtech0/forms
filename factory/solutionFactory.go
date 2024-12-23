package factory

type SolutionChoiceFactory struct {
	commonMapper *CommonFieldQuestionMapper
}

func NewSolutionChoiceFactory() *SingleChoiceFactory {
	return &SingleChoiceFactory{
		commonMapper: NewCommonFieldQuestionMapper(),
	}
}
