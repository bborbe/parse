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
)
