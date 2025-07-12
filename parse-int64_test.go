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

type MyInt64 int64

var _ = DescribeTable("ParseInt64",
	func(value interface{}, expectedResult int64, expectError bool) {
		result, err := parse.ParseInt64(context.Background(), value)
		if expectError {
			Expect(err).NotTo(BeNil())
			Expect(result).To(Equal(int64(0)))
		} else {
			Expect(err).To(BeNil())
			Expect(result).To(Equal(expectedResult))
		}
	},
	Entry("string", "1337", int64(1337), false),
	Entry("stringer", MyStringer("1337"), int64(1337), false),
	Entry("custom int64", MyInt64(1337), int64(1337), false),
	Entry("int64", 1337, int64(1337), false),
	Entry("float32", float32(1337), int64(1337), false),
	Entry("int64", 1337, int64(1337), false),
	Entry("invalid", "banana", int64(0), true),
)

var _ = DescribeTable("ParseInt64Default",
	func(value interface{}, defaultValue int64, expectedResult int64) {
		result := parse.ParseInt64Default(context.Background(), value, defaultValue)
		Expect(result).To(Equal(expectedResult))
	},
	Entry("valid string", "1337", int64(999), int64(1337)),
	Entry("valid int", 42, int64(999), int64(42)),
	Entry("valid int64", int64(100), int64(999), int64(100)),
	Entry("valid float32", float32(25), int64(999), int64(25)),
	Entry("custom type", MyInt64(200), int64(999), int64(200)),
	Entry("stringer", MyStringer("75"), int64(999), int64(75)),
	Entry("invalid returns default", "banana", int64(999), int64(999)),
	Entry("nil returns default", nil, int64(123), int64(123)),
	Entry("unsupported type returns default", []int{1, 2}, int64(888), int64(888)),
)
