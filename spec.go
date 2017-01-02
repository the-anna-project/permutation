package permutation

// List is supposed to be permuted by a permutation service.
type List interface {
	// Indizes returns the list's current indizes.
	Indizes() []int
	// MaxGrowth returns the list's configured growth limit. The growth limit
	// is used to prevent infinite permutations. E.g. MaxGrowth set to 4 will not
	// permute up to a list of 5 raw values.
	MaxGrowth() int
	// PermutedValues returns the list's permuted values.
	PermutedValues() []interface{}
	// RawValues returns the list's unpermuted raw values.
	RawValues() []interface{}
	// SetIndizes sets the given indizes of the current permutation list.
	SetIndizes(indizes []int)
	SetMaxGrowth(maxGrowth int)
	SetRawValues(rawValues []interface{})
}

// Service creates permutations of arbitrary lists as configured.
type Service interface {
	// PermuteBy permutes the configured values by applying the given delta to
	// the currently configured indizes. Error might indicate that the configured
	// max growth is reached. It processes the essentiaal operation of permuting
	// the list's indizes using the given delta. Indizes are capped by the
	// amount of configured values. E.g. having a list of two values configured
	// will increment the indizes in a base 2 number system.
	//
	//     Imagine the following values being configured.
	//
	//         []interface{"a", "b"}
	//
	//     Here the index 0 translates to indizes []int{0} and causes the
	//     following permutation on the raw values.
	//
	//         []interface{"a"}
	//
	//     Here the index 1 translates to indizes []int{1} and causes the
	//     following permutation on the raw values.
	//
	//         []interface{"b"}
	//
	//     Here the index 00 translates to indizes []int{0, 0} and causes the
	//     following permutation on the raw values.
	//
	//         []interface{"a", "a"}
	//
	PermuteBy(list List, delta int) error
}
