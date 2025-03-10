package common

import (
	"github.com/grafana/agent/pkg/river/token"
	"github.com/grafana/agent/pkg/river/token/builder"
	"github.com/prometheus/prometheus/storage"
)

// ConvertAppendable implements both the [builder.Tokenizer] and
// [storage.Appendable] interfaces. This allows us to set component.Arguments
// that leverage [storage.Appendable] with an implementation that can be
// tokenized as a specific string.
type ConvertAppendable struct {
	storage.Appendable

	Expr string // The specific string to return during tokenization.
}

func (f ConvertAppendable) RiverCapsule() {}
func (f ConvertAppendable) RiverTokenize() []builder.Token {
	return []builder.Token{{
		Tok: token.STRING,
		Lit: f.Expr,
	}}
}

var _ storage.Appendable = (*ConvertAppendable)(nil)
var _ builder.Tokenizer = ConvertAppendable{}
