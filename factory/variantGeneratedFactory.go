package factory

type VariantGeneratedFactory struct {
	questionGeneratedFactory *QuestionGeneratedFactory
}

func NewVariantGeneratedFactory() *VariantGeneratedFactory {
	return &VariantGeneratedFactory{
		questionGeneratedFactory: NewQuestionGeneratedFactory(),
	}
}
