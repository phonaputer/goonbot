package executer

type Token interface {
	Apply(current RollingResult) RollingResult
}
