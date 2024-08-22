// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse_test

import (
	"context"

	"github.com/bborbe/parse"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = DescribeTable("ParseString",
	func(value interface{}, expectedResult string, expectError bool) {
		result, err := parse.ParseString(context.Background(), value)
		if expectError {
			Expect(err).NotTo(BeNil())
			Expect(result).To(Equal(""))
		} else {
			Expect(err).To(BeNil())
			Expect(result).To(Equal(expectedResult))
		}
	},
	Entry("string", "banana", "banana", false),
	Entry("int", 42, "42", false),
	Entry("int32", int32(42), "42", false),
	Entry("int64", int64(42), "42", false),
	Entry("uint", uint(42), "42", false),
	Entry("uint32", uint32(42), "42", false),
	Entry("uint64", uint64(42), "42", false),
	Entry("struct", struct{}{}, "", true),
)
