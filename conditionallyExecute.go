package condexec

type conditionallyExecute struct {
	condition bool
	onTrue    func()
	onFalse   func()
}

func New(condition bool) *conditionallyExecute {
	return &conditionallyExecute{condition: condition}
}
func (ce *conditionallyExecute) OnTrue(onTrueFunc func()) *conditionallyExecute {
	ce.onTrue = onTrueFunc
	return ce
}
func (ce *conditionallyExecute) OnFalse(onFalseFunc func()) *conditionallyExecute {
	ce.onFalse = onFalseFunc
	return ce
}
func (ce *conditionallyExecute) Execute() *conditionallyExecute{
	if ce.condition{
		ce.onTrue()
	}else{
		ce.onFalse()
	}
	return ce
}
