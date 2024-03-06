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

var _ = Describe("ParseFloat64", func() {
	var ctx context.Context
	var err error
	var input interface{}
	var result float64
	BeforeEach(func() {
		ctx = context.Background()
	})
	JustBeforeEach(func() {
		result, err = parse.ParseFloat64(ctx, input)
	})
	Context("float64", func() {
		BeforeEach(func() {
			input = 123.45
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("returns correct result", func() {
			Expect(result).To(Equal(123.45))
		})
	})
	Context("float32", func() {
		BeforeEach(func() {
			input = float32(123)
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("returns correct result", func() {
			Expect(result).To(Equal(123.0))
		})
	})
	Context("string", func() {
		type Bar float64
		BeforeEach(func() {
			input = "123.45"
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("returns correct result", func() {
			Expect(result).To(Equal(123.45))
		})
	})
})
