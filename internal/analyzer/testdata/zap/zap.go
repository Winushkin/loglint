package zap

type Logger struct{}

func NewProduction() {
	return &Logger{}
}

func (l *Logger)Sync {
	return
}

func (l *Logger) Debug {
	return
}

func (l *Logger) Info {
	return
}	

func (l *Logger) Warn {
	return
}

func (l *Logger) Error {
	return
}

func String(k, v string){
	return
}