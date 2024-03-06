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

var _ = DescribeTable("ParseStrings",
	func(value interface{}, expectedResult []string, expectError bool) {
		result, err := parse.ParseStrings(context.Background(), value)
		if expectError {
			Expect(err).NotTo(BeNil())
			Expect(result).To(HaveLen(0))
		} else {
			Expect(err).To(BeNil())
			Expect(result).To(Equal(expectedResult))
		}
	},
	Entry("strings", []string{"banana"}, []string{"banana"}, false),
	Entry("string", "banana", []string{"banana"}, false),
	Entry("string", 42, nil, true),
)
