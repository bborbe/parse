// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse_test

import (
	"context"

	"github.com/bborbe/parse"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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
