package permutation

// ListConfig represents the configuration used to create a new list.
type ListConfig struct {
}

// DefaultListConfig provides a default configuration to create a new list by
// best effort.
func DefaultListConfig() ListConfig {
	return ListConfig{}
}

// NewList creates a new configured list.
func NewList(config ListConfig) (List, error) {
	newList := &list{
		// Settings.
		indizes:        []int{},
		maxGrowth:      5,
		minGrowth:      1,
		permutedValues: []interface{}{},
		rawValues:      []interface{}{},
	}

	return newList, nil
}

type list struct {
	// Settings.

	// indizes represents an ordered list where each index represents a raw value
	// position.
	indizes []int
	// maxGrowth represents the maximum length PermutedValues is allowed to grow.
	maxGrowth int
	// minGrowth represents the minimum length PermutedValues is allowed to have.
	minGrowth int
	// permutedValues represents the permuted list of RawValues. Initially this is
	// the zero value []interface{}{}.
	permutedValues []interface{}
	// rawValues represents the values being used to permute PermutedValues.
	rawValues []interface{}
}

func (l *list) Indizes() []int {
	return l.indizes
}

func (l *list) MaxGrowth() int {
	return l.maxGrowth
}

func (l *list) MinGrowth() int {
	return l.minGrowth
}

func (l *list) PermutedValues() []interface{} {
	return l.permutedValues
}

func (l *list) RawValues() []interface{} {
	return l.rawValues
}

func (l *list) SetIndizes(indizes []int) {
	l.indizes = indizes
	l.permutedValues = permuteValues(l)
}

func (l *list) SetMaxGrowth(maxGrowth int) {
	if maxGrowth <= 0 {
		panic(maskAnyf(invalidConfigError, "max growth must be 1 or greater"))
	}
	l.maxGrowth = maxGrowth
}

func (l *list) SetMinGrowth(minGrowth int) {
	if minGrowth > l.MaxGrowth() {
		panic(maskAnyf(invalidConfigError, "min growth must not be greater than max growth"))
	}
	if minGrowth <= 0 {
		panic(maskAnyf(invalidConfigError, "min growth must be 1 or greater"))
	}
	l.minGrowth = minGrowth
}

func (l *list) SetRawValues(rawValues []interface{}) {
	l.rawValues = rawValues
}
