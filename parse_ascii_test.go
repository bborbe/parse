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

var _ = Describe("ParseAscii", func() {
	var ctx context.Context
	BeforeEach(func() {
		ctx = context.Background()
	})
	It("remove umlaute", func() {
		Expect(parse.ParseAscii(ctx, "žůžoüÄÅ")).To(Equal("zuzouAA"))
	})
	It("leave normal chars untouched", func() {
		Expect(parse.ParseAscii(ctx, "abc0123")).To(Equal("abc0123"))
	})
	It("handle empty string", func() {
		Expect(parse.ParseAscii(ctx, "")).To(Equal(""))
	})
	It("handle int input", func() {
		Expect(parse.ParseAscii(ctx, 123)).To(Equal("123"))
	})
	It("handle float input", func() {
		Expect(parse.ParseAscii(ctx, 12.34)).To(Equal("12.34"))
	})
	It("handle bool input", func() {
		Expect(parse.ParseAscii(ctx, true)).To(Equal("true"))
	})
	It("handle complex unicode", func() {
		Expect(parse.ParseAscii(ctx, "Ñoël Ümlaut")).To(Equal("Noel Umlaut"))
	})
	It("handle stringer input", func() {
		Expect(parse.ParseAscii(ctx, MyStringer("tëst"))).To(Equal("test"))
	})
})
