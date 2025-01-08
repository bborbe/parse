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
