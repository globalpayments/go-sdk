package builders

type BaseBuilder struct {
	Validations Validations
}

func NewBaseBuilder() *BaseBuilder {
	return &BaseBuilder{
		Validations: *NewValidations(),
	}
}

func (b *BaseBuilder) GetValidations() Validations {
	return b.Validations
}

func (b *BaseBuilder) SetValidations(validations Validations) {
	b.Validations = validations
}

func (b *BaseBuilder) Execute(configName string) error {
	err := b.Validations.Validate(*b)
	if err != nil {
		return err
	}
	return nil
}
