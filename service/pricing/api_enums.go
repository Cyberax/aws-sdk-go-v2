// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pricing

type FilterType string

// Enum values for FilterType
const (
	FilterTypeTermMatch FilterType = "TERM_MATCH"
)

func (enum FilterType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FilterType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
