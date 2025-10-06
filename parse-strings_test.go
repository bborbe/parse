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
)
