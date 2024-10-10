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
})
