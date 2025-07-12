// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/parse"
)

type MyFloat64 float64

var _ = DescribeTable("ParseFloat64",
	func(value interface{}, expectedResult float64, expectError bool) {
		result, err := parse.ParseFloat64(context.Background(), value)
		if expectError {
			Expect(err).NotTo(BeNil())
			Expect(result).To(Equal(0.0))
		} else {
			Expect(err).To(BeNil())
			Expect(result).To(Equal(expectedResult))
		}
	},
	Entry("string", "1.2", 1.2, false),
	Entry("stringer", MyStringer("1.2"), 1.2, false),
	Entry("custom float64", MyFloat64(1.2), 1.2, false),
	Entry("int", 1, 1.0, false),
	Entry("float32", float32(100), float64(100), false),
	Entry("float64", 1.2, 1.2, false),
	Entry("invalid", "banana", 0.0, true),
)

var _ = DescribeTable("ParseFloat64Default",
	func(value interface{}, defaultValue float64, expectedResult float64) {
		result := parse.ParseFloat64Default(context.Background(), value, defaultValue)
		Expect(result).To(Equal(expectedResult))
	},
	Entry("valid string", "1.2", 999.9, 1.2),
	Entry("valid int", 42, 999.9, 42.0),
	Entry("valid float64", 3.14, 999.9, 3.14),
	Entry("valid float32", float32(2.5), 999.9, 2.5),
	Entry("custom type", MyFloat64(1.5), 999.9, 1.5),
	Entry("stringer", MyStringer("2.7"), 999.9, 2.7),
	Entry("invalid returns default", "banana", 999.9, 999.9),
	Entry("nil returns default", nil, 123.45, 123.45),
	Entry("unsupported type returns default", []int{1, 2}, 88.88, 88.88),
)
