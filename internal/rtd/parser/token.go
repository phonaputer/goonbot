package parser

type Token interface {
	Apply(currentTotal int) int
}
