// The MIT License (MIT)
// 
// Copyright (c) 2021 Uber Technologies, Inc.
// 
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// 
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: uber/cadence/api/v1/tasklist.proto

package apiv1

var yarpcFileDescriptorClosure216fa006947e00a0 = [][]byte{
	// uber/cadence/api/v1/tasklist.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x4e, 0x23, 0x37,
		0x14, 0xee, 0x10, 0xd8, 0x85, 0x43, 0x61, 0x67, 0xbd, 0x3f, 0x90, 0x6c, 0xb7, 0xa5, 0xb9, 0x58,
		0xa1, 0x55, 0x3b, 0x29, 0x54, 0xbd, 0xea, 0x45, 0x15, 0x12, 0xd4, 0x1d, 0x11, 0xb2, 0xd1, 0x64,
		0x96, 0x8a, 0x4a, 0x95, 0xeb, 0x8c, 0x4d, 0xb0, 0x66, 0xc6, 0x1e, 0xd9, 0x9e, 0x00, 0x0f, 0xd2,
		0x3e, 0x4c, 0x9f, 0xa8, 0x8f, 0x51, 0xd9, 0x99, 0x09, 0x29, 0xa4, 0xbd, 0xb3, 0xcf, 0x77, 0xbe,
		0xf3, 0xf3, 0x9d, 0x63, 0x43, 0xbb, 0x9c, 0x30, 0xd5, 0x49, 0x08, 0x65, 0x22, 0x61, 0x1d, 0x52,
		0xf0, 0xce, 0xec, 0xa8, 0x63, 0x88, 0x4e, 0x33, 0xae, 0x4d, 0x50, 0x28, 0x69, 0x24, 0x7a, 0x61,
		0x7d, 0x82, 0xca, 0x27, 0x20, 0x05, 0x0f, 0x66, 0x47, 0xad, 0x2f, 0xa7, 0x52, 0x4e, 0x33, 0xd6,
		0x71, 0x2e, 0x93, 0xf2, 0xaa, 0x43, 0x4b, 0x45, 0x0c, 0x97, 0x62, 0x4e, 0x6a, 0x7d, 0xf5, 0x10,
		0x37, 0x3c, 0x67, 0xda, 0x90, 0xbc, 0xa8, 0x1c, 0x1e, 0x05, 0xb8, 0x51, 0xa4, 0x28, 0x98, 0xd2,
		0x73, 0xbc, 0xfd, 0x09, 0x36, 0x63, 0xa2, 0xd3, 0x01, 0xd7, 0x06, 0x21, 0x58, 0x17, 0x24, 0x67,
		0xfb, 0xde, 0x81, 0x77, 0xb8, 0x15, 0xb9, 0x33, 0xfa, 0x01, 0xd6, 0x53, 0x2e, 0xe8, 0xfe, 0xda,
		0x81, 0x77, 0xb8, 0x7b, 0xfc, 0x75, 0xb0, 0xa2, 0xc8, 0xa0, 0x0e, 0x70, 0xc6, 0x05, 0x8d, 0x9c,
		0x7b, 0x9b, 0x80, 0x5f, 0x5b, 0xcf, 0x99, 0x21, 0x94, 0x18, 0x82, 0xce, 0xe1, 0x65, 0x4e, 0x6e,
		0xb1, 0x6d, 0x5b, 0xe3, 0x82, 0x29, 0xac, 0x59, 0x22, 0x05, 0x75, 0xe9, 0xb6, 0x8f, 0xbf, 0x08,
		0xe6, 0x95, 0x06, 0x75, 0xa5, 0x41, 0x5f, 0x96, 0x93, 0x8c, 0x5d, 0x90, 0xac, 0x64, 0xd1, 0xf3,
		0x9c, 0xdc, 0xda, 0x80, 0x7a, 0xc4, 0xd4, 0xd8, 0xd1, 0xda, 0x9f, 0xa0, 0x59, 0xa7, 0x18, 0x11,
		0x65, 0xb8, 0x55, 0x65, 0x91, 0xcb, 0x87, 0x46, 0xca, 0xee, 0xaa, 0x4e, 0xec, 0x11, 0xbd, 0x83,
		0x67, 0xf2, 0x46, 0x30, 0x85, 0xaf, 0xa5, 0x36, 0xd8, 0xf5, 0xb9, 0xe6, 0xd0, 0x1d, 0x67, 0xfe,
		0x20, 0xb5, 0x19, 0x92, 0x9c, 0xb5, 0xff, 0xf6, 0x60, 0xb7, 0x8e, 0x3b, 0x36, 0xc4, 0x94, 0x1a,
		0x7d, 0x03, 0x68, 0x42, 0x92, 0x34, 0x93, 0x53, 0x9c, 0xc8, 0x52, 0x18, 0x7c, 0xcd, 0x85, 0x71,
		0xb1, 0x1b, 0x91, 0x5f, 0x21, 0x3d, 0x0b, 0x7c, 0xe0, 0xc2, 0xa0, 0xb7, 0x00, 0x8a, 0x11, 0x8a,
		0x33, 0x36, 0x63, 0x99, 0xcb, 0xd1, 0x88, 0xb6, 0xac, 0x65, 0x60, 0x0d, 0xe8, 0x0d, 0x6c, 0x91,
		0x24, 0xad, 0xd0, 0x86, 0x43, 0x37, 0x49, 0x92, 0xce, 0xc1, 0x77, 0xf0, 0x4c, 0x11, 0xc3, 0x96,
		0xd5, 0x59, 0x3f, 0xf0, 0x0e, 0xbd, 0x68, 0xc7, 0x9a, 0x17, 0xbd, 0xa3, 0x3e, 0xec, 0x58, 0x19,
		0x31, 0xa7, 0x78, 0x92, 0xc9, 0x24, 0xdd, 0xdf, 0x70, 0x1a, 0x1e, 0xfc, 0xe7, 0x78, 0xc2, 0xfe,
		0x89, 0xf5, 0x8b, 0xb6, 0x2d, 0x2d, 0xa4, 0xee, 0xd2, 0xfe, 0x09, 0xb6, 0x97, 0x30, 0xd4, 0x84,
		0x4d, 0x6d, 0x88, 0x32, 0x98, 0xd3, 0xaa, 0xb9, 0xa7, 0xee, 0x1e, 0x52, 0xf4, 0x0a, 0x9e, 0x30,
		0x41, 0x2d, 0x30, 0xef, 0x67, 0x83, 0x09, 0x1a, 0xd2, 0xf6, 0x9f, 0x1e, 0xc0, 0x48, 0x66, 0x19,
		0x53, 0xa1, 0xb8, 0x92, 0xa8, 0x0f, 0x7e, 0x46, 0xb4, 0xc1, 0x24, 0x49, 0x98, 0xd6, 0xd8, 0xae,
		0x62, 0x35, 0xdc, 0xd6, 0xa3, 0xe1, 0xc6, 0xf5, 0x9e, 0x46, 0xbb, 0x96, 0xd3, 0x75, 0x14, 0x6b,
		0x44, 0x2d, 0xd8, 0xe4, 0x94, 0x09, 0xc3, 0xcd, 0x5d, 0x35, 0xa1, 0xc5, 0x7d, 0x95, 0x3e, 0x8d,
		0x15, 0xfa, 0xb4, 0xff, 0xf2, 0xa0, 0x39, 0x36, 0x3c, 0x49, 0xef, 0x4e, 0x6f, 0x59, 0x52, 0xda,
		0xd5, 0xe8, 0x1a, 0xa3, 0xf8, 0xa4, 0x34, 0x4c, 0xa3, 0x9f, 0xc1, 0xbf, 0x91, 0x2a, 0x65, 0xca,
		0xed, 0x22, 0xb6, 0x6f, 0xb0, 0xaa, 0xf3, 0xed, 0xff, 0xee, 0x77, 0xb4, 0x3b, 0xa7, 0x2d, 0x1e,
		0x4c, 0x0c, 0x4d, 0x9d, 0x5c, 0x33, 0x5a, 0x66, 0x0c, 0x1b, 0x89, 0xe7, 0xea, 0xd9, 0xb6, 0x65,
		0x69, 0x5c, 0xed, 0xdb, 0xc7, 0xcd, 0xc7, 0x6b, 0x5d, 0xbd, 0xe0, 0xe8, 0x75, 0xcd, 0x8d, 0xe5,
		0xd8, 0x32, 0xe3, 0x39, 0xb1, 0xfd, 0x87, 0x07, 0x7b, 0x8f, 0x36, 0xbb, 0x27, 0xc5, 0x15, 0x9f,
		0xa2, 0x7d, 0x78, 0x3a, 0x63, 0x4a, 0x73, 0x29, 0xea, 0x11, 0x55, 0x57, 0x14, 0xc0, 0x0b, 0x51,
		0xe6, 0xd8, 0xad, 0x5e, 0x51, 0xb3, 0xb4, 0xab, 0x62, 0x23, 0x7a, 0x2e, 0xca, 0x3c, 0x62, 0x84,
		0x2e, 0xc2, 0x69, 0xf4, 0x1d, 0xbc, 0xb4, 0xfe, 0x37, 0x8a, 0x5b, 0x3d, 0xef, 0x09, 0x0d, 0x47,
		0x40, 0xa2, 0xcc, 0x7f, 0xb1, 0xd0, 0x3d, 0xe3, 0xfd, 0xef, 0xf0, 0xf9, 0xf2, 0x4b, 0x47, 0x2d,
		0x78, 0x1d, 0x77, 0xc7, 0x67, 0x78, 0x10, 0x8e, 0x63, 0x7c, 0x16, 0x0e, 0xfb, 0x38, 0x1c, 0x5e,
		0x74, 0x07, 0x61, 0xdf, 0xff, 0x0c, 0x35, 0xe1, 0xd5, 0x03, 0x6c, 0xf8, 0x31, 0x3a, 0xef, 0x0e,
		0x7c, 0x6f, 0x05, 0x34, 0x8e, 0xc3, 0xde, 0xd9, 0xa5, 0xbf, 0xf6, 0x9e, 0xde, 0x67, 0x88, 0xef,
		0x0a, 0xf6, 0xef, 0x0c, 0xf1, 0xe5, 0xe8, 0x74, 0x29, 0xc3, 0x1b, 0xd8, 0x7b, 0x80, 0xf5, 0x4f,
		0x7b, 0xe1, 0x38, 0xfc, 0x38, 0xf4, 0xbd, 0x15, 0x60, 0xb7, 0x17, 0x87, 0x17, 0x61, 0x7c, 0xe9,
		0xaf, 0x9d, 0xfc, 0x06, 0x7b, 0x89, 0xcc, 0x57, 0x4d, 0xfa, 0x64, 0x67, 0xa1, 0xbb, 0x9d, 0xd6,
		0xc8, 0xfb, 0xf5, 0x68, 0xca, 0xcd, 0x75, 0x39, 0x09, 0x12, 0x99, 0x77, 0x96, 0xff, 0xf0, 0x6f,
		0x39, 0xcd, 0x3a, 0x53, 0x39, 0xff, 0x56, 0xab, 0x0f, 0xfd, 0x47, 0x52, 0xf0, 0xd9, 0xd1, 0xe4,
		0x89, 0xb3, 0x7d, 0xff, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x81, 0xda, 0x9c, 0xf6, 0xf4, 0x05,
		0x00, 0x00,
	},
	// google/protobuf/duration.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a,
		0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0x56,
		0x5c, 0x1c, 0x2e, 0x50, 0x25, 0x42, 0x12, 0x5c, 0xec, 0xc5, 0xa9, 0xc9, 0xf9, 0x79, 0x29, 0xc5,
		0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x30, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x5e, 0x62, 0x5e,
		0x7e, 0xb1, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x84, 0xe3, 0xd4, 0xcc, 0xc8, 0x25, 0x9c,
		0x9c, 0x9f, 0xab, 0x87, 0x66, 0xa6, 0x13, 0x2f, 0xcc, 0xc4, 0x00, 0x90, 0x48, 0x00, 0x63, 0x94,
		0x21, 0x54, 0x45, 0x7a, 0x7e, 0x4e, 0x62, 0x5e, 0xba, 0x5e, 0x7e, 0x51, 0x3a, 0xc2, 0x81, 0x25,
		0x95, 0x05, 0xa9, 0xc5, 0xfa, 0xd9, 0x79, 0xf9, 0xe5, 0x79, 0x70, 0xc7, 0x16, 0x24, 0xfd, 0x60,
		0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce, 0x1d, 0xa2, 0x39, 0x00, 0xaa,
		0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4, 0x35, 0x89, 0x0d, 0x6c, 0x94,
		0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8a, 0xb4, 0xc3, 0xfb, 0x00, 0x00, 0x00,
	},
	// google/protobuf/timestamp.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4f, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0xc9, 0xcc, 0x4d,
		0x2d, 0x2e, 0x49, 0xcc, 0x2d, 0xd0, 0x03, 0x0b, 0x09, 0xf1, 0x43, 0x14, 0xe8, 0xc1, 0x14, 0x28,
		0x59, 0x73, 0x71, 0x86, 0xc0, 0xd4, 0x08, 0x49, 0x70, 0xb1, 0x17, 0xa7, 0x26, 0xe7, 0xe7, 0xa5,
		0x14, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0xc1, 0xb8, 0x42, 0x22, 0x5c, 0xac, 0x79, 0x89,
		0x79, 0xf9, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x53, 0x2b, 0x23, 0x97,
		0x70, 0x72, 0x7e, 0xae, 0x1e, 0x9a, 0xa1, 0x4e, 0x7c, 0x70, 0x23, 0x03, 0x40, 0x42, 0x01, 0x8c,
		0x51, 0x46, 0x50, 0x25, 0xe9, 0xf9, 0x39, 0x89, 0x79, 0xe9, 0x7a, 0xf9, 0x45, 0xe9, 0x48, 0x6e,
		0xac, 0x2c, 0x48, 0x2d, 0xd6, 0xcf, 0xce, 0xcb, 0x2f, 0xcf, 0x43, 0xb8, 0xb7, 0x20, 0xe9, 0x07,
		0x23, 0xe3, 0x22, 0x26, 0x66, 0xf7, 0x00, 0xa7, 0x55, 0x4c, 0x72, 0xee, 0x10, 0xdd, 0x01, 0x50,
		0x2d, 0x7a, 0xe1, 0xa9, 0x39, 0x39, 0xde, 0x20, 0x0d, 0x21, 0x20, 0xbd, 0x49, 0x6c, 0x60, 0xb3,
		0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0x65, 0xce, 0x7d, 0xff, 0x00, 0x00, 0x00,
	},
	// google/protobuf/wrappers.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
		0xcf, 0x49, 0xd5, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x2f, 0x4a, 0x2c,
		0x28, 0x48, 0x2d, 0x2a, 0xd6, 0x03, 0x8b, 0x08, 0xf1, 0x43, 0xe4, 0xf5, 0x60, 0xf2, 0x4a, 0xca,
		0x5c, 0xdc, 0x2e, 0xf9, 0xa5, 0x49, 0x39, 0xa9, 0x61, 0x89, 0x39, 0xa5, 0xa9, 0x42, 0x22, 0x5c,
		0xac, 0x65, 0x20, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x63, 0x10, 0x84, 0xa3, 0xa4, 0xc4, 0xc5,
		0xe5, 0x96, 0x93, 0x9f, 0x58, 0x82, 0x45, 0x0d, 0x13, 0x92, 0x1a, 0xcf, 0xbc, 0x12, 0x33, 0x13,
		0x2c, 0x6a, 0x98, 0x61, 0x6a, 0x94, 0xb9, 0xb8, 0x43, 0x71, 0x29, 0x62, 0x41, 0x35, 0xc8, 0xd8,
		0x08, 0x8b, 0x1a, 0x56, 0x34, 0x83, 0xb0, 0x2a, 0xe2, 0x85, 0x29, 0x52, 0xe4, 0xe2, 0x74, 0xca,
		0xcf, 0xcf, 0xc1, 0xa2, 0x84, 0x03, 0xc9, 0x9c, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x2c, 0x8a,
		0x38, 0x91, 0x1c, 0xe4, 0x54, 0x59, 0x92, 0x5a, 0x8c, 0x45, 0x0d, 0x0f, 0x54, 0x8d, 0x53, 0x33,
		0x23, 0x97, 0x70, 0x72, 0x7e, 0xae, 0x1e, 0x5a, 0xf0, 0x3a, 0xf1, 0x86, 0x43, 0xc3, 0x3f, 0x00,
		0x24, 0x12, 0xc0, 0x18, 0x65, 0x08, 0x55, 0x91, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f,
		0x94, 0x8e, 0x88, 0xab, 0x92, 0xca, 0x82, 0xd4, 0x62, 0xfd, 0xec, 0xbc, 0xfc, 0xf2, 0x3c, 0x78,
		0xbc, 0x15, 0x24, 0xfd, 0x60, 0x64, 0x5c, 0xc4, 0xc4, 0xec, 0x1e, 0xe0, 0xb4, 0x8a, 0x49, 0xce,
		0x1d, 0xa2, 0x39, 0x00, 0xaa, 0x43, 0x2f, 0x3c, 0x35, 0x27, 0xc7, 0x1b, 0xa4, 0x3e, 0x04, 0xa4,
		0x35, 0x89, 0x0d, 0x6c, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x92, 0x48, 0x30, 0x06,
		0x02, 0x00, 0x00,
	},
}
