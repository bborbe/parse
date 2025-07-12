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

var _ = DescribeTable("ParseInt64ArrayDefault",
	func(input interface{}, defaultValue []int64, expectedResult []int64) {
		result := parse.ParseInt64ArrayDefault(context.Background(), input, defaultValue)
		Expect(result).To(Equal(expectedResult))
	},
	Entry("valid []interface{}", []interface{}{1, 2, 3}, []int64{999}, []int64{1, 2, 3}),
	Entry("valid []int", []int{10, 20}, []int64{999}, []int64{10, 20}),
	Entry("valid []int64", []int64{100, 200}, []int64{999}, []int64{100, 200}),
	Entry("valid []float64", []float64{1.0, 2.0}, []int64{999}, []int64{1, 2}),
	Entry("invalid returns default", "invalid", []int64{888, 777}, []int64{888, 777}),
	Entry("nil returns default", nil, []int64{123}, []int64{123}),
)
