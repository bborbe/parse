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

type MyInt int

var _ = DescribeTable("ParseInt",
	func(value interface{}, expectedResult int, expectError bool) {
		result, err := parse.ParseInt(context.Background(), value)
		if expectError {
			Expect(err).NotTo(BeNil())
			Expect(result).To(Equal(0))
		} else {
			Expect(err).To(BeNil())
			Expect(result).To(Equal(expectedResult))
		}
	},
	Entry("string", "1337", 1337, false),
	Entry("stringer", MyStringer("1337"), 1337, false),
	Entry("custom int", MyInt(1337), 1337, false),
	Entry("int", 1337, 1337, false),
	Entry("float32", float32(1337), 1337, false),
	Entry("int", 1337, 1337, false),
	Entry("invalid", "banana", 0, true),
)

var _ = DescribeTable("ParseIntDefault",
	func(value interface{}, defaultValue int, expectedResult int) {
		result := parse.ParseIntDefault(context.Background(), value, defaultValue)
		Expect(result).To(Equal(expectedResult))
	},
	Entry("valid string", "1337", 999, 1337),
	Entry("valid int", 42, 999, 42),
	Entry("valid float32", float32(25), 999, 25),
	Entry("custom type", MyInt(100), 999, 100),
	Entry("stringer", MyStringer("75"), 999, 75),
	Entry("invalid returns default", "banana", 999, 999),
	Entry("nil returns default", nil, 123, 123),
	Entry("unsupported type returns default", []int{1, 2}, 888, 888),
)
