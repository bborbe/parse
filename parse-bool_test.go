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

type MyBool bool

var _ = DescribeTable("ParseBool",
	func(value interface{}, expectedResult bool, expectError bool) {
		result, err := parse.ParseBool(context.Background(), value)
		if expectError {
			Expect(err).NotTo(BeNil())
			Expect(result).To(Equal(false))
		} else {
			Expect(err).To(BeNil())
			Expect(result).To(Equal(expectedResult))
		}
	},
	Entry("true", true, true, false),
	Entry("false", false, false, false),
	Entry("error", "error", false, true),
	Entry("true string", "true", true, false),
	Entry("false string", "false", false, false),
	Entry("TRUE string", "TRUE", true, false),
	Entry("FALSE string", "FALSE", false, false),
	Entry("True string", "True", true, false),
	Entry("False string", "False", false, false),
	Entry("stringer true", MyStringer("true"), true, false),
	Entry("stringer false", MyStringer("false"), false, false),
	Entry("MyBool false", MyBool(true), true, false),
)
