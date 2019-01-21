package strategy

type strategy interface{
	analyze() strategy.Result
}

