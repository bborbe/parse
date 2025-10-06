// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parse_test

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/parse"
)

var _ = DescribeTable("ParseTime",
	func(value interface{}, format string, expectedResult string, expectError bool) {
		result, err := parse.ParseTime(context.Background(), value, format)
		if expectError {
			Expect(err).NotTo(BeNil())
			Expect(result).To(Equal(time.Time{}))
		} else {
			Expect(err).To(BeNil())
			expected, parseErr := time.Parse(format, expectedResult)
			Expect(parseErr).To(BeNil())
			Expect(result).To(Equal(expected))
		}
	},
	Entry("string RFC3339", "2023-12-25T10:30:00Z", time.RFC3339, "2023-12-25T10:30:00Z", false),
	Entry("string custom format", "25/12/2023", "02/01/2006", "25/12/2023", false),
	Entry("string date only", "2023-12-25", "2006-01-02", "2023-12-25", false),
	Entry("string time only", "10:30:00", "15:04:05", "10:30:00", false),
	Entry("small int as string", 2023, "2006", "2023", false),
	Entry("small int64 as string", int64(2024), "2006", "2024", false),
	Entry("float as string", 2025.0, "2006", "2025", false),
	Entry("invalid time string", "invalid-time", time.RFC3339, "", true),
	Entry("empty string", "", time.RFC3339, "", true),
	Entry("wrong format", "2023-12-25", time.RFC3339, "", true),
	Entry("nil value", nil, time.RFC3339, "", true),
	Entry("unsupported type", []int{1, 2, 3}, time.RFC3339, "", true),
)

var _ = DescribeTable(
	"ParseTimeDefault",
	func(value interface{}, format string, defaultValue string, expectedResult string) {
		defaultTime, parseErr := time.Parse(time.RFC3339, defaultValue)
		Expect(parseErr).To(BeNil())

		result := parse.ParseTimeDefault(context.Background(), value, format, defaultTime)

		expected, parseErr := time.Parse(time.RFC3339, expectedResult)
		Expect(parseErr).To(BeNil())
		Expect(result).To(Equal(expected))
	},
	Entry(
		"valid time",
		"2023-12-25T10:30:00Z",
		time.RFC3339,
		"2000-01-01T00:00:00Z",
		"2023-12-25T10:30:00Z",
	),
	Entry(
		"invalid time returns default",
		"invalid",
		time.RFC3339,
		"2000-01-01T00:00:00Z",
		"2000-01-01T00:00:00Z",
	),
	Entry("nil returns default", nil, time.RFC3339, "2000-01-01T00:00:00Z", "2000-01-01T00:00:00Z"),
	Entry(
		"wrong format returns default",
		"2023-12-25",
		time.RFC3339,
		"2000-01-01T00:00:00Z",
		"2000-01-01T00:00:00Z",
	),
)
