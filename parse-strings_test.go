// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/parse"
)

// Test types for string subtype slice support
type Direction string

// Test type implementing String() string for HasString interface support
type DirectionWithString struct {
	val string
}

func (d DirectionWithString) String() string {
	return d.val
}

// Test type implementing Strings() []string for HasStrings interface support
type DirectionsWithStrings struct {
	vals []string
}

func (d DirectionsWithStrings) Strings() []string {
	return d.vals
}

var _ = DescribeTable(
	"ParseStrings",
	func(value interface{}, expectedResult []string, expectError bool) {
		result, err := parse.ParseStrings(context.Background(), value)
		if expectError {
			Expect(err).NotTo(BeNil())
			Expect(result).To(BeNil())
		} else {
			Expect(err).To(BeNil())
			Expect(result).To(Equal(expectedResult))
		}
	},
	Entry("nil", nil, nil, false),
	Entry("strings", []string{"banana", "apple"}, []string{"banana", "apple"}, false),
	Entry("single string", "banana", []string{"banana"}, false),
	Entry(
		"interface slice",
		[]interface{}{"banana", 42, true},
		[]string{"banana", "42", "true"},
		false,
	),
	Entry("float64 slice", []float64{1.2, 3.4}, []string{"1.2", "3.4"}, false),
	Entry("bool slice", []bool{true, false}, []string{"true", "false"}, false),
	Entry("int slice", []int{1, 2, 3}, []string{"1", "2", "3"}, false),
	Entry("int32 slice", []int32{10, 20}, []string{"10", "20"}, false),
	Entry("int64 slice", []int64{100, 200}, []string{"100", "200"}, false),
	Entry("unsupported type", 42, nil, true),
	Entry("unsupported slice type", []float32{1.0, 2.0}, nil, true),
	// String subtype slice tests
	Entry("slice of string subtype", []Direction{"up", "down"}, []string{"up", "down"}, false),
	Entry("empty slice of string subtype", []Direction{}, []string{}, false),
	// HasStrings interface tests
	Entry(
		"HasStrings interface",
		DirectionsWithStrings{vals: []string{"north", "south"}},
		[]string{"north", "south"},
		false,
	),
	Entry(
		"HasStrings interface empty",
		DirectionsWithStrings{vals: []string{}},
		[]string{},
		false,
	),
	// HasString interface tests (single value converted to slice)
	Entry(
		"HasString interface",
		DirectionWithString{val: "east"},
		[]string{"east"},
		false,
	),
	// Slice of HasString tests
	Entry(
		"slice of HasString",
		[]DirectionWithString{{val: "west"}, {val: "northwest"}},
		[]string{"west", "northwest"},
		false,
	),
	Entry(
		"empty slice of HasString",
		[]DirectionWithString{},
		[]string{},
		false,
	),
	// Slice of interface type ([]parse.HasString)
	Entry(
		"slice of HasString interface type",
		[]parse.HasString{
			DirectionWithString{val: "interface1"},
			DirectionWithString{val: "interface2"},
		},
		[]string{"interface1", "interface2"},
		false,
	),
)

var _ = DescribeTable(
	"ParseStringsDefault",
	func(value interface{}, defaultValue []string, expectedResult []string) {
		result := parse.ParseStringsDefault(context.Background(), value, defaultValue)
		Expect(result).To(Equal(expectedResult))
	},
	Entry("valid nil", nil, []string{"default"}, nil),
	Entry(
		"valid strings",
		[]string{"banana", "apple"},
		[]string{"default"},
		[]string{"banana", "apple"},
	),
	Entry("valid single string", "test", []string{"default"}, []string{"test"}),
	Entry("valid int slice", []int{1, 2}, []string{"default"}, []string{"1", "2"}),
	Entry(
		"invalid returns default",
		42,
		[]string{"default1", "default2"},
		[]string{"default1", "default2"},
	),
	Entry(
		"unsupported type returns default",
		[]float32{1.0},
		[]string{"fallback"},
		[]string{"fallback"},
	),
	// New type support tests
	Entry(
		"string subtype slice",
		[]Direction{"left", "right"},
		[]string{"default"},
		[]string{"left", "right"},
	),
	Entry(
		"HasStrings interface",
		DirectionsWithStrings{vals: []string{"up"}},
		[]string{"default"},
		[]string{"up"},
	),
	Entry(
		"HasString interface",
		DirectionWithString{val: "down"},
		[]string{"default"},
		[]string{"down"},
	),
)
