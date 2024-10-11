package domain

type Formatter interface {
	Format(node *CustomNode) (string, error)
}
