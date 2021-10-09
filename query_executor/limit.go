package query_executor

type Limit struct {
	LimitValue int
	Count      int
	ScanPtr    *Scan
}

func (l *Limit) Init(limitVal int) {
	l.ScanPtr = &Scan{}
	l.ScanPtr.Init(&data, 2)
	l.LimitValue = limitVal
}

func (l *Limit) Next() *Movie {
	if l.Count >= l.LimitValue {
		return nil
	}
	res := l.ScanPtr.Next()
	l.Count++
	return res
}
