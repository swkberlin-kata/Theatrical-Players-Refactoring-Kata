package theatre

type Performance struct {
	PlayID   string
	Audience int
}

type Play struct {
	Name string
	Type string
}

type Invoice struct {
	Customer     string
	Performances []Performance
}

type ResultPerPerformance struct {
	price float32;
	playName string
	playAudience int
}

type TotalResult struct {
	customerName string
	volumeCredits int
	results [] ResultPerPerformance
}