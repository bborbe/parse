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

var _ = DescribeTable("ParseInt64Array",
	func(input interface{}, expectedResult []int64, hasError bool) {
		result, err := parse.ParseInt64Array(context.Background(), input)
		Expect(result).To(Equal(expectedResult))
		if hasError {
			Expect(err).NotTo(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	},
	Entry("[]interface", []interface{}{1, 2, 3}, []int64{1, 2, 3}, false),
	Entry("[]int", []int{1, 2, 3}, []int64{1, 2, 3}, false),
	Entry("[]int64", []int64{1, 2, 3}, []int64{1, 2, 3}, false),
	Entry("[]float64", []float64{1, 2, 3}, []int64{1, 2, 3}, false),
)
