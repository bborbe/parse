// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse_test

import (
	"bytes"
	"context"

	"github.com/bborbe/parse"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ParseInt", func() {
	var ctx context.Context
	var err error
	var input interface{}
	var result int
	BeforeEach(func() {
		ctx = context.Background()
	})
	JustBeforeEach(func() {
		result, err = parse.ParseInt(ctx, input)
	})
	Context("int", func() {
		BeforeEach(func() {
			input = 123
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("returns correct result", func() {
			Expect(result).To(Equal(123))
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
			Expect(result).To(Equal(123))
		})
	})
	Context("string", func() {
		type Bar int
		BeforeEach(func() {
			input = "123"
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("returns correct result", func() {
			Expect(result).To(Equal(123))
		})
	})
	Context("Stringer", func() {
		type Bar int
		BeforeEach(func() {
			input = bytes.NewBufferString("123")
		})
		It("returns no error", func() {
			Expect(err).To(BeNil())
		})
		It("returns correct result", func() {
			Expect(result).To(Equal(123))
		})
	})
})
