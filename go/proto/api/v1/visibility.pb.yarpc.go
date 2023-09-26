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
// source: uber/cadence/api/v1/visibility.proto

package apiv1

var yarpcFileDescriptorClosurec1b2132ef24217c8 = [][]byte{
	// uber/cadence/api/v1/visibility.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcd, 0x6e, 0xd3, 0x40,
		0x14, 0x85, 0x71, 0x69, 0x23, 0x71, 0x53, 0xa8, 0x35, 0x08, 0x02, 0xae, 0x20, 0xc8, 0x62, 0x51,
		0x21, 0x31, 0x96, 0xcb, 0xb2, 0x0b, 0x94, 0x60, 0x83, 0x46, 0x84, 0x24, 0x38, 0x6e, 0x4a, 0xd8,
		0x58, 0x63, 0x7b, 0x1a, 0x46, 0x8c, 0x3d, 0x96, 0x3d, 0x4e, 0xdb, 0xa7, 0xe0, 0x3d, 0x79, 0x0a,
		0xe4, 0xbf, 0x0a, 0x09, 0x57, 0xec, 0xec, 0x73, 0xcf, 0xf9, 0x34, 0xf7, 0x07, 0x5e, 0x97, 0x21,
		0xcb, 0xad, 0x88, 0xc6, 0x2c, 0x8d, 0x98, 0x45, 0x33, 0x6e, 0xed, 0x6c, 0x6b, 0xc7, 0x0b, 0x1e,
		0x72, 0xc1, 0xd5, 0x0d, 0xce, 0x72, 0xa9, 0x24, 0x7a, 0x5c, 0xb9, 0x70, 0xeb, 0xc2, 0x34, 0xe3,
		0x78, 0x67, 0x1b, 0xe3, 0xad, 0x94, 0x5b, 0xc1, 0xac, 0xda, 0x12, 0x96, 0x97, 0x96, 0xe2, 0x09,
		0x2b, 0x14, 0x4d, 0xb2, 0x26, 0x65, 0x98, 0x7d, 0xec, 0x2b, 0x99, 0xff, 0xbc, 0x14, 0xf2, 0xaa,
		0xf1, 0x98, 0x5f, 0x61, 0x74, 0xd1, 0x2a, 0xee, 0x35, 0x8b, 0x4a, 0xc5, 0x65, 0xfa, 0x91, 0x0b,
		0xc5, 0x72, 0x34, 0x86, 0x61, 0x67, 0x0e, 0x78, 0xfc, 0x4c, 0x7b, 0xa5, 0x9d, 0x3c, 0xf0, 0xa0,
		0x93, 0x48, 0x8c, 0x9e, 0xc0, 0x20, 0x2f, 0xd3, 0xaa, 0xb6, 0x57, 0xd7, 0x0e, 0xf2, 0x32, 0x25,
		0xb1, 0x79, 0x02, 0xa8, 0x43, 0xfa, 0x37, 0x19, 0x6b, 0x69, 0x08, 0xf6, 0x53, 0x9a, 0xb0, 0x16,
		0x53, 0x7f, 0x9b, 0xbf, 0x34, 0x38, 0x5a, 0x29, 0x9a, 0x2b, 0x9f, 0x27, 0x9d, 0xef, 0x3d, 0x3c,
		0x64, 0x34, 0x17, 0x9c, 0x15, 0x2a, 0xa8, 0x1a, 0xaa, 0x03, 0xc3, 0x53, 0x03, 0x37, 0xdd, 0xe2,
		0xae, 0x5b, 0xec, 0x77, 0xdd, 0x7a, 0x87, 0x5d, 0xa0, 0x92, 0xd0, 0x19, 0x0c, 0x05, 0x55, 0xb7,
		0xf1, 0xbd, 0xff, 0xc6, 0xa1, 0xb1, 0x57, 0x82, 0xb9, 0x81, 0xc3, 0x95, 0xa2, 0xaa, 0x2c, 0xda,
		0xd7, 0x10, 0x18, 0x14, 0xf5, 0x7f, 0xfd, 0x8c, 0x47, 0xa7, 0x36, 0xee, 0xd9, 0x04, 0xfe, 0x67,
		0x82, 0x1f, 0x84, 0x2c, 0x58, 0x03, 0xf2, 0x5a, 0xc0, 0x9b, 0xdf, 0x1a, 0xe8, 0x24, 0x8d, 0xd9,
		0x35, 0x8b, 0xd7, 0x54, 0x94, 0xac, 0x9a, 0x0d, 0x7a, 0x09, 0x06, 0x99, 0x3b, 0xee, 0x37, 0xd7,
		0x09, 0xd6, 0x93, 0xd9, 0xb9, 0x1b, 0xf8, 0x9b, 0xa5, 0x1b, 0x90, 0xf9, 0x7a, 0x32, 0x23, 0x8e,
		0x7e, 0x0f, 0xbd, 0x80, 0xe7, 0x3d, 0xf5, 0x95, 0xef, 0x91, 0xf9, 0x27, 0x5d, 0xbb, 0x23, 0xfe,
		0xd9, 0xdd, 0x5c, 0x2c, 0x3c, 0x47, 0xdf, 0x43, 0x06, 0x3c, 0xed, 0xc5, 0xfb, 0xfa, 0xfd, 0x3b,
		0xd0, 0xce, 0xe2, 0x7c, 0x3a, 0x73, 0xf5, 0x7d, 0x74, 0x0c, 0xa3, 0x9e, 0xf2, 0x74, 0xb1, 0x98,
		0xe9, 0x07, 0x68, 0x0c, 0xc7, 0x7d, 0xd9, 0x89, 0xef, 0xfa, 0xe4, 0x8b, 0xab, 0x0f, 0xa6, 0x01,
		0x8c, 0x22, 0x99, 0xf4, 0x0d, 0x6b, 0x7a, 0xb4, 0xbe, 0xbd, 0xee, 0x65, 0xb5, 0x8c, 0xa5, 0xf6,
		0xdd, 0xde, 0x72, 0xf5, 0xa3, 0x0c, 0x71, 0x24, 0x13, 0xeb, 0xef, 0x9b, 0x7d, 0xcb, 0x63, 0x61,
		0x6d, 0x65, 0x73, 0xe1, 0xed, 0x01, 0x9f, 0xd1, 0x8c, 0xef, 0xec, 0x70, 0x50, 0x6b, 0xef, 0xfe,
		0x04, 0x00, 0x00, 0xff, 0xff, 0x46, 0x40, 0x75, 0x5d, 0x40, 0x03, 0x00, 0x00,
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
	// uber/cadence/api/v1/workflow.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x59, 0x4d, 0x73, 0xdb, 0xc8,
		0xd1, 0x7e, 0x41, 0x4a, 0xb2, 0xd4, 0xd4, 0x07, 0x34, 0x92, 0x2c, 0x5a, 0xde, 0xb5, 0x65, 0xee,
		0xda, 0x2b, 0xf3, 0x5d, 0x49, 0x2b, 0xef, 0xda, 0x5e, 0x5b, 0x71, 0x1c, 0x08, 0x84, 0x2c, 0xd8,
		0x14, 0xc8, 0x0c, 0x41, 0xcb, 0xda, 0x4a, 0x82, 0x82, 0xc8, 0x91, 0x84, 0x98, 0x04, 0x58, 0xc0,
		0xd0, 0xb6, 0xee, 0xa9, 0xca, 0x39, 0x97, 0x54, 0x2a, 0xa7, 0xfc, 0x80, 0xa4, 0x52, 0xa9, 0x9c,
		0x53, 0xa9, 0xca, 0x21, 0xb7, 0x5c, 0xf3, 0x1f, 0xf2, 0x2f, 0x52, 0x33, 0x18, 0x90, 0xe0, 0x27,
		0xe8, 0xa4, 0x6a, 0x73, 0x13, 0x7a, 0x9e, 0xa7, 0xd1, 0xd3, 0xd3, 0xfd, 0xf4, 0x80, 0x82, 0x5c,
		0xfb, 0x8c, 0xf8, 0xbb, 0x35, 0xbb, 0x4e, 0xdc, 0x1a, 0xd9, 0xb5, 0x5b, 0xce, 0xee, 0xbb, 0xbd,
		0xdd, 0xf7, 0x9e, 0xff, 0xf6, 0xbc, 0xe1, 0xbd, 0xdf, 0x69, 0xf9, 0x1e, 0xf5, 0xd0, 0x0a, 0xc3,
		0xec, 0x08, 0xcc, 0x8e, 0xdd, 0x72, 0x76, 0xde, 0xed, 0x6d, 0xdc, 0xba, 0xf0, 0xbc, 0x8b, 0x06,
		0xd9, 0xe5, 0x90, 0xb3, 0xf6, 0xf9, 0x6e, 0xbd, 0xed, 0xdb, 0xd4, 0xf1, 0xdc, 0x90, 0xb4, 0x71,
		0xbb, 0x7f, 0x9d, 0x3a, 0x4d, 0x12, 0x50, 0xbb, 0xd9, 0x12, 0x80, 0xcd, 0x61, 0x6f, 0xae, 0x79,
		0xcd, 0x66, 0xc7, 0xc5, 0xd0, 0xd8, 0xa8, 0x1d, 0xbc, 0x6d, 0x38, 0x01, 0x0d, 0x31, 0xb9, 0x3f,
		0xcc, 0xc2, 0xda, 0x89, 0x08, 0x57, 0xfb, 0x40, 0x6a, 0x6d, 0x16, 0x82, 0xee, 0x9e, 0x7b, 0xa8,
		0x0a, 0x28, 0xda, 0x87, 0x45, 0xa2, 0x95, 0xac, 0xb4, 0x29, 0x6d, 0x65, 0x1e, 0xdc, 0xdb, 0x19,
		0xb2, 0xa5, 0x9d, 0x01, 0x3f, 0x78, 0xf9, 0x7d, 0xbf, 0x09, 0x3d, 0x84, 0x29, 0x7a, 0xd5, 0x22,
		0xd9, 0x14, 0x77, 0x74, 0x67, 0xac, 0x23, 0xf3, 0xaa, 0x45, 0x30, 0x87, 0xa3, 0x27, 0x00, 0x01,
		0xb5, 0x7d, 0x6a, 0xb1, 0x34, 0x64, 0xd3, 0x9c, 0xbc, 0xb1, 0x13, 0xe6, 0x68, 0x27, 0xca, 0xd1,
		0x8e, 0x19, 0xe5, 0x08, 0xcf, 0x71, 0x34, 0x7b, 0x66, 0xd4, 0x5a, 0xc3, 0x0b, 0x48, 0x48, 0x9d,
		0x4a, 0xa6, 0x72, 0x34, 0xa7, 0x9a, 0x30, 0x1f, 0x52, 0x03, 0x6a, 0xd3, 0x76, 0x90, 0x9d, 0xde,
		0x94, 0xb6, 0x16, 0x1f, 0xec, 0x4d, 0xb6, 0x7b, 0x95, 0x31, 0x2b, 0x9c, 0x88, 0x33, 0xb5, 0xee,
		0x03, 0xba, 0x0b, 0x8b, 0x97, 0x4e, 0x40, 0x3d, 0xff, 0xca, 0x6a, 0x10, 0xf7, 0x82, 0x5e, 0x66,
		0x67, 0x36, 0xa5, 0xad, 0x34, 0x5e, 0x10, 0xd6, 0x22, 0x37, 0xa2, 0x9f, 0xc0, 0x5a, 0xcb, 0xf6,
		0x89, 0x4b, 0xbb, 0xe9, 0xb7, 0x1c, 0xf7, 0xdc, 0xcb, 0x5e, 0xe3, 0x5b, 0xd8, 0x1a, 0x1a, 0x45,
		0x99, 0x33, 0x7a, 0x4e, 0x12, 0xaf, 0xb4, 0x06, 0x8d, 0x48, 0x81, 0xc5, 0xae, 0x5b, 0x9e, 0x99,
		0xd9, 0xc4, 0xcc, 0x2c, 0x74, 0x18, 0x3c, 0x3b, 0xdb, 0x30, 0xd5, 0x24, 0x4d, 0x2f, 0x3b, 0xc7,
		0x89, 0x37, 0x86, 0xc6, 0x73, 0x4c, 0x9a, 0x1e, 0xe6, 0x30, 0x84, 0x61, 0x39, 0x20, 0xb6, 0x5f,
		0xbb, 0xb4, 0x6c, 0x4a, 0x7d, 0xe7, 0xac, 0x4d, 0x49, 0x90, 0x05, 0xce, 0xbd, 0x3b, 0x94, 0x5b,
		0xe1, 0x68, 0xa5, 0x03, 0xc6, 0x72, 0xd0, 0x67, 0x41, 0x45, 0x58, 0xb6, 0xdb, 0xd4, 0xb3, 0x7c,
		0x12, 0x10, 0x6a, 0xb5, 0x3c, 0xc7, 0xa5, 0x41, 0x36, 0xc3, 0x7d, 0x6e, 0x0e, 0xf5, 0x89, 0x19,
		0xb0, 0xcc, 0x71, 0x78, 0x89, 0x51, 0x63, 0x06, 0x74, 0x13, 0xe6, 0x58, 0x7b, 0x58, 0xac, 0x3f,
		0xb2, 0xf3, 0x9b, 0xd2, 0xd6, 0x1c, 0x9e, 0x65, 0x86, 0xa2, 0x13, 0x50, 0xb4, 0x0e, 0xd7, 0x9c,
		0xc0, 0xaa, 0xf9, 0x9e, 0x9b, 0x5d, 0xd8, 0x94, 0xb6, 0x66, 0xf1, 0x8c, 0x13, 0xa8, 0xbe, 0xe7,
		0xa2, 0x7d, 0xc8, 0xb4, 0x5b, 0x75, 0x9b, 0x8a, 0x02, 0x5b, 0x4c, 0x4c, 0x23, 0x84, 0x70, 0x9e,
		0xc3, 0x9f, 0x83, 0xdc, 0xb2, 0x7d, 0xea, 0xf0, 0x63, 0xa8, 0x79, 0xee, 0xb9, 0x73, 0x91, 0x5d,
		0xda, 0x4c, 0x6f, 0x65, 0x1e, 0x3c, 0x9f, 0xac, 0xca, 0xd8, 0x61, 0xb2, 0x53, 0x0f, 0x5d, 0xa8,
		0xdc, 0x83, 0xe6, 0x52, 0xff, 0x0a, 0x2f, 0xb5, 0x7a, 0xad, 0x1b, 0x07, 0xb0, 0x3a, 0x0c, 0x88,
		0x64, 0x48, 0xbf, 0x25, 0x57, 0xbc, 0xb5, 0xe7, 0x30, 0xfb, 0x13, 0xad, 0xc2, 0xf4, 0x3b, 0xbb,
		0xd1, 0x0e, 0xbb, 0x74, 0x0e, 0x87, 0x0f, 0x4f, 0x53, 0xdf, 0x4a, 0xb9, 0xdf, 0xa4, 0xe0, 0xd6,
		0x60, 0xa5, 0x73, 0x67, 0x42, 0xbf, 0xd0, 0xd3, 0x78, 0x16, 0x43, 0xbd, 0xf8, 0x74, 0xe8, 0x5e,
		0x4c, 0x91, 0xda, 0x58, 0x92, 0x6d, 0xd8, 0xec, 0x56, 0xa5, 0x68, 0x78, 0xcf, 0xea, 0xb6, 0xaf,
		0xd7, 0xa6, 0x42, 0x39, 0x6e, 0x0c, 0x24, 0xb8, 0x20, 0x02, 0xc0, 0x9f, 0x74, 0x5c, 0x54, 0xb8,
		0x08, 0x78, 0x6a, 0xd4, 0xd0, 0x5e, 0x9b, 0xa2, 0x13, 0xb8, 0xc9, 0xc3, 0x1b, 0xe1, 0x3d, 0x9d,
		0xe4, 0x7d, 0x9d, 0xb1, 0x87, 0x38, 0xce, 0xfd, 0x43, 0x82, 0x95, 0x21, 0xed, 0xc7, 0xaa, 0xaa,
		0xee, 0x35, 0x6d, 0xc7, 0xb5, 0x9c, 0xba, 0x48, 0xf2, 0x6c, 0x68, 0xd0, 0xeb, 0xe8, 0x36, 0x64,
		0xc4, 0xa2, 0x6b, 0x37, 0xa3, 0x7c, 0x43, 0x68, 0x32, 0xec, 0x26, 0x19, 0x21, 0xc3, 0xe9, 0xff,
		0x56, 0x86, 0xef, 0xc0, 0xbc, 0xe3, 0x3a, 0xd4, 0xb1, 0x29, 0xa9, 0xb3, 0xb8, 0xa6, 0xb8, 0x02,
		0x65, 0x3a, 0x36, 0xbd, 0x9e, 0xfb, 0x95, 0x04, 0x6b, 0xda, 0x07, 0x4a, 0x7c, 0xd7, 0x6e, 0x7c,
		0x2f, 0xa3, 0xa1, 0x3f, 0xa6, 0xd4, 0x60, 0x4c, 0xbf, 0x9e, 0x81, 0x95, 0x32, 0x71, 0xeb, 0x8e,
		0x7b, 0xa1, 0xd4, 0xa8, 0xf3, 0xce, 0xa1, 0x57, 0x3c, 0xa2, 0xdb, 0x90, 0xb1, 0xc5, 0x73, 0x37,
		0xcb, 0x10, 0x99, 0xf4, 0x3a, 0x3a, 0x84, 0x85, 0x0e, 0x20, 0x71, 0xfe, 0x44, 0xae, 0xf9, 0xfc,
		0x99, 0xb7, 0x63, 0x4f, 0xe8, 0x39, 0x4c, 0xb3, 0x59, 0x10, 0x8e, 0xa0, 0xc5, 0x07, 0xf7, 0x87,
		0x8b, 0x70, 0x6f, 0x84, 0x4c, 0xf6, 0x09, 0x0e, 0x79, 0x48, 0x87, 0xe5, 0x4b, 0x62, 0xfb, 0xf4,
		0x8c, 0xd8, 0xd4, 0xaa, 0x13, 0x6a, 0x3b, 0x8d, 0x40, 0x0c, 0xa5, 0x4f, 0x46, 0x28, 0xfa, 0x55,
		0xc3, 0xb3, 0xeb, 0x58, 0xee, 0xd0, 0x0a, 0x21, 0x0b, 0xbd, 0x84, 0x95, 0x86, 0x1d, 0x50, 0xab,
		0xeb, 0x8f, 0x0b, 0xd0, 0x74, 0xa2, 0x00, 0x2d, 0x33, 0xda, 0x51, 0xc4, 0xe2, 0x3a, 0x74, 0x08,
		0xdc, 0x18, 0x76, 0x05, 0xa9, 0x87, 0x9e, 0x66, 0x12, 0x3d, 0x2d, 0x31, 0x52, 0x25, 0xe4, 0x70,
		0x3f, 0x59, 0xb8, 0x66, 0x53, 0x4a, 0x9a, 0x2d, 0xca, 0xc7, 0xd4, 0x34, 0x8e, 0x1e, 0xd1, 0x7d,
		0x90, 0x9b, 0xf6, 0x07, 0xa7, 0xd9, 0x6e, 0x5a, 0xc2, 0x14, 0xf0, 0x91, 0x33, 0x8d, 0x97, 0x84,
		0x5d, 0x11, 0x66, 0x36, 0x9b, 0x82, 0xda, 0x25, 0xa9, 0xb7, 0x1b, 0x51, 0x24, 0x73, 0xc9, 0xb3,
		0xa9, 0xc3, 0xe0, 0x71, 0xa8, 0xb0, 0x44, 0x3e, 0xb4, 0x9c, 0xb0, 0x67, 0x43, 0x1f, 0x90, 0xe8,
		0x63, 0xb1, 0x4b, 0xe1, 0x4e, 0x9e, 0xc3, 0x3c, 0x4f, 0xca, 0xb9, 0xed, 0x34, 0xda, 0x3e, 0x11,
		0x83, 0x65, 0xf8, 0x31, 0x1d, 0x86, 0x18, 0x9c, 0x61, 0x0c, 0xf1, 0x80, 0xbe, 0x82, 0x55, 0xee,
		0x80, 0xd5, 0x3a, 0xf1, 0x2d, 0xa7, 0x4e, 0x5c, 0xea, 0xd0, 0x2b, 0x31, 0x5b, 0x10, 0x5b, 0x3b,
		0xe1, 0x4b, 0xba, 0x58, 0x41, 0x8f, 0x60, 0x3d, 0x3a, 0x82, 0x7e, 0xd2, 0x02, 0x27, 0xad, 0x89,
		0xe5, 0x5e, 0x5e, 0xee, 0xcf, 0x29, 0xb8, 0x21, 0xca, 0x4e, 0xbd, 0x74, 0x1a, 0xf5, 0xef, 0xa5,
		0x61, 0xbf, 0x8c, 0xb9, 0x65, 0x4d, 0x15, 0xd7, 0x30, 0xf9, 0x7d, 0xec, 0x12, 0xc7, 0x95, 0xac,
		0xbf, 0xbd, 0xd3, 0x03, 0xed, 0x8d, 0x5e, 0x83, 0xb8, 0xab, 0x08, 0x51, 0x6e, 0x79, 0x0d, 0xa7,
		0x76, 0xc5, 0xdb, 0x63, 0x71, 0x44, 0xa0, 0xa1, 0xe2, 0x72, 0x21, 0x2e, 0x73, 0x34, 0x5e, 0x6e,
		0xf5, 0x9b, 0xd0, 0x75, 0x98, 0x09, 0x25, 0x95, 0x37, 0xc7, 0x1c, 0x16, 0x4f, 0xb9, 0xbf, 0xa7,
		0x3a, 0x72, 0x52, 0x20, 0x35, 0x27, 0x88, 0xf2, 0xd5, 0xe9, 0x72, 0x29, 0xb9, 0xcb, 0x23, 0x62,
		0x4f, 0x97, 0x0f, 0x56, 0x70, 0xea, 0x63, 0x2b, 0xf8, 0x19, 0xcc, 0xf7, 0x34, 0x63, 0xf2, 0x9d,
		0x37, 0x13, 0x0c, 0x6f, 0xc4, 0xa9, 0xde, 0x46, 0xc4, 0xb0, 0xee, 0xf9, 0xce, 0x85, 0xe3, 0xda,
		0x0d, 0xab, 0x2f, 0xc8, 0x64, 0xe9, 0x58, 0x8b, 0xa8, 0x95, 0x78, 0xb0, 0xb9, 0xbf, 0xa4, 0xe0,
		0x46, 0x24, 0x77, 0x45, 0xaf, 0x66, 0x37, 0x0a, 0x4e, 0xd0, 0xb2, 0x69, 0xed, 0x72, 0x32, 0x75,
		0xfe, 0xdf, 0xa7, 0xeb, 0x67, 0x70, 0xab, 0x37, 0x02, 0xcb, 0x3b, 0xb7, 0xe8, 0xa5, 0x13, 0x58,
		0xf1, 0x2c, 0x8e, 0x77, 0xb8, 0xd1, 0x13, 0x51, 0xe9, 0xdc, 0xbc, 0x74, 0x02, 0xa1, 0x69, 0xe8,
		0x53, 0x00, 0x7e, 0xeb, 0xa0, 0xde, 0x5b, 0x12, 0x56, 0xe1, 0x3c, 0xe6, 0xd7, 0x24, 0x93, 0x19,
		0x72, 0x2f, 0x21, 0x13, 0xbf, 0x88, 0xee, 0xc3, 0x8c, 0xb8, 0xcb, 0x4a, 0xfc, 0x2e, 0xf8, 0x59,
		0xc2, 0x5d, 0x96, 0x5f, 0xf3, 0x05, 0x25, 0xf7, 0xc7, 0x14, 0x2c, 0xf6, 0x2e, 0xa1, 0x2f, 0x60,
		0xe9, 0xcc, 0x71, 0x6d, 0xff, 0xca, 0xaa, 0x5d, 0x92, 0xda, 0xdb, 0xa0, 0xdd, 0x14, 0x87, 0xb0,
		0x18, 0x9a, 0x55, 0x61, 0x45, 0x6b, 0x30, 0xe3, 0xb7, 0xdd, 0x68, 0xf8, 0xce, 0xe1, 0x69, 0xbf,
		0xcd, 0x6e, 0x29, 0xcf, 0xe0, 0xe6, 0xb9, 0xe3, 0x07, 0x6c, 0x60, 0x85, 0xc5, 0x6e, 0xd5, 0xbc,
		0x66, 0xab, 0x41, 0x7a, 0x3a, 0x39, 0xcb, 0x21, 0x51, 0x3b, 0xa8, 0x11, 0x80, 0xd3, 0xe7, 0x6b,
		0x3e, 0xb1, 0x3b, 0x67, 0x93, 0x9c, 0xca, 0x8c, 0xc0, 0x0b, 0x19, 0x5e, 0xe0, 0xc2, 0xec, 0xb8,
		0x17, 0x93, 0x96, 0xe9, 0x7c, 0x44, 0xe0, 0x0e, 0x6e, 0x01, 0xf0, 0x0f, 0x04, 0x6a, 0x9f, 0x35,
		0xc2, 0xa9, 0x36, 0x8b, 0x63, 0x96, 0xfc, 0x9f, 0x24, 0x58, 0x1d, 0x36, 0xb3, 0x51, 0x0e, 0x6e,
		0x95, 0x35, 0xa3, 0xa0, 0x1b, 0x2f, 0x2c, 0x45, 0x35, 0xf5, 0xd7, 0xba, 0x79, 0x6a, 0x55, 0x4c,
		0xc5, 0xd4, 0x2c, 0xdd, 0x78, 0xad, 0x14, 0xf5, 0x82, 0xfc, 0x7f, 0xe8, 0x73, 0xd8, 0x1c, 0x81,
		0xa9, 0xa8, 0x47, 0x5a, 0xa1, 0x5a, 0xd4, 0x0a, 0xb2, 0x34, 0xc6, 0x53, 0xc5, 0x54, 0xb0, 0xa9,
		0x15, 0xe4, 0x14, 0xfa, 0x7f, 0xf8, 0x62, 0x04, 0x46, 0x55, 0x0c, 0x55, 0x2b, 0x5a, 0x58, 0xfb,
		0x71, 0x55, 0xab, 0x30, 0x70, 0x3a, 0xff, 0x8b, 0x6e, 0xcc, 0x3d, 0x0a, 0x14, 0x7f, 0x53, 0x41,
		0x53, 0xf5, 0x8a, 0x5e, 0x32, 0xc6, 0xc5, 0xdc, 0x87, 0x19, 0x11, 0x73, 0x3f, 0x2a, 0x8a, 0x39,
		0xff, 0xcb, 0x54, 0xf7, 0xf7, 0x03, 0xbd, 0x8e, 0x49, 0xbb, 0xa3, 0xb9, 0x9f, 0xc3, 0xe6, 0x49,
		0x09, 0xbf, 0x3a, 0x2c, 0x96, 0x4e, 0x2c, 0xbd, 0x60, 0x61, 0xad, 0x5a, 0xd1, 0xac, 0x72, 0xa9,
		0xa8, 0xab, 0xa7, 0xb1, 0x48, 0xbe, 0x85, 0x6f, 0x46, 0xa2, 0x94, 0x22, 0xb3, 0x16, 0xaa, 0xe5,
		0xa2, 0xae, 0xb2, 0xb7, 0x1e, 0x2a, 0x7a, 0x51, 0x2b, 0x58, 0x25, 0xa3, 0x78, 0x2a, 0x4b, 0xe8,
		0x4b, 0xd8, 0x9a, 0x94, 0x29, 0xa7, 0xd0, 0x36, 0xdc, 0x1f, 0x89, 0xc6, 0xda, 0x4b, 0x4d, 0x35,
		0x63, 0xf0, 0x34, 0xda, 0x83, 0xed, 0x91, 0x70, 0x53, 0xc3, 0xc7, 0xba, 0xc1, 0x13, 0x7a, 0x68,
		0xe1, 0xaa, 0x61, 0xe8, 0xc6, 0x0b, 0x79, 0x2a, 0xff, 0x3b, 0x09, 0x96, 0x07, 0x86, 0x11, 0xba,
		0x0d, 0x37, 0xcb, 0x0a, 0xd6, 0x0c, 0xd3, 0x52, 0x8b, 0xa5, 0x61, 0x09, 0x18, 0x01, 0x50, 0x0e,
		0x14, 0xa3, 0x50, 0x32, 0x64, 0x09, 0xdd, 0x83, 0xdc, 0x30, 0x80, 0xa8, 0x05, 0x51, 0x1a, 0x72,
		0x0a, 0xdd, 0x81, 0x4f, 0x87, 0xe1, 0x3a, 0xd1, 0xca, 0xe9, 0xfc, 0xbf, 0x52, 0xf0, 0xc9, 0xb8,
		0x9f, 0x29, 0x58, 0x05, 0x76, 0xb6, 0xad, 0xbd, 0xd1, 0xd4, 0xaa, 0xc9, 0xce, 0x3c, 0xf4, 0xc7,
		0x4e, 0xbe, 0x5a, 0x89, 0x45, 0x1e, 0x4f, 0xe9, 0x08, 0xb0, 0x5a, 0x3a, 0x2e, 0x17, 0x35, 0x93,
		0x57, 0x53, 0x1e, 0xee, 0x25, 0xc1, 0xc3, 0x03, 0x96, 0x53, 0x3d, 0x67, 0x3b, 0xca, 0x35, 0xdf,
		0x37, 0x6b, 0x05, 0xb4, 0x03, 0xf9, 0x24, 0x74, 0x27, 0x0b, 0x05, 0x79, 0x0a, 0x7d, 0x03, 0x5f,
		0x25, 0x07, 0x6e, 0x98, 0xba, 0x51, 0xd5, 0x0a, 0x96, 0x52, 0xb1, 0x0c, 0xed, 0x44, 0x9e, 0x9e,
		0x64, 0xbb, 0xa6, 0x7e, 0xcc, 0xea, 0xb3, 0x6a, 0xca, 0x33, 0xf9, 0xbf, 0x4a, 0x70, 0x5d, 0xf5,
		0x5c, 0xea, 0xb8, 0x6d, 0xa2, 0x04, 0x06, 0x79, 0xaf, 0x87, 0xf7, 0x1c, 0xcf, 0x47, 0x77, 0xe1,
		0x4e, 0xe4, 0x5f, 0xb8, 0xb7, 0x74, 0x43, 0x37, 0x75, 0xc5, 0x2c, 0xe1, 0x58, 0x7e, 0xc7, 0xc2,
		0x58, 0x43, 0x16, 0x34, 0x1c, 0xe6, 0x75, 0x34, 0x0c, 0x6b, 0x26, 0x3e, 0x15, 0xa5, 0x10, 0x2a,
		0xcc, 0x68, 0xac, 0x8a, 0x59, 0x7f, 0x8b, 0xfe, 0x97, 0xd3, 0xf9, 0xdf, 0x4b, 0x90, 0x11, 0xdf,
		0xb6, 0xfc, 0xd3, 0x27, 0x0b, 0xab, 0x6c, 0x83, 0xa5, 0xaa, 0x69, 0x99, 0xa7, 0x65, 0xad, 0xb7,
		0x86, 0x7b, 0x56, 0xb8, 0x3c, 0x58, 0x66, 0x29, 0xcc, 0x4e, 0xa8, 0x24, 0xbd, 0x00, 0xf1, 0x16,
		0x86, 0xe1, 0x60, 0x39, 0x35, 0x16, 0x13, 0xfa, 0x49, 0xa3, 0x0d, 0xb8, 0xde, 0x83, 0x39, 0xd2,
		0x14, 0x6c, 0x1e, 0x68, 0x8a, 0x29, 0x4f, 0xe5, 0x7f, 0x2b, 0xc1, 0x8d, 0x48, 0x09, 0x4d, 0x36,
		0x58, 0x9d, 0x26, 0xa9, 0x97, 0xda, 0x54, 0xb5, 0xdb, 0x01, 0x41, 0xf7, 0xe1, 0x6e, 0x47, 0xc3,
		0x4c, 0xa5, 0xf2, 0xaa, 0x7b, 0x56, 0x96, 0xaa, 0xb0, 0xe6, 0xee, 0xee, 0x26, 0x11, 0x2a, 0x42,
		0x90, 0x25, 0xf4, 0x05, 0x7c, 0x36, 0x1e, 0x8a, 0xb5, 0x8a, 0x66, 0xca, 0xa9, 0xfc, 0x3f, 0x33,
		0xb0, 0x1e, 0x0f, 0x8e, 0x7d, 0x20, 0x90, 0x7a, 0x18, 0xda, 0x3d, 0xc8, 0xf5, 0x3a, 0x11, 0x3a,
		0xd7, 0x1f, 0xd7, 0x1e, 0x6c, 0x8f, 0xc1, 0x55, 0x8d, 0x23, 0xc5, 0x28, 0xb0, 0xe7, 0x08, 0x24,
		0x4b, 0xe8, 0x39, 0xec, 0x8f, 0xa1, 0x1c, 0x28, 0x85, 0x6e, 0x96, 0x3b, 0x13, 0x47, 0x31, 0x4d,
		0xac, 0x1f, 0x54, 0x4d, 0xad, 0x22, 0xa7, 0x90, 0x06, 0x4a, 0x82, 0x83, 0x5e, 0x1d, 0x1a, 0xea,
		0x26, 0x8d, 0x9e, 0xc0, 0xc3, 0xa4, 0x38, 0xc2, 0x92, 0xd1, 0x8f, 0x35, 0x1c, 0xa7, 0x4e, 0xa1,
		0xa7, 0xf0, 0x28, 0x81, 0x2a, 0xde, 0x3c, 0xc0, 0x9d, 0x46, 0xfb, 0xf0, 0x38, 0x31, 0x7a, 0xb5,
		0x84, 0x0b, 0xd6, 0xb1, 0x82, 0x5f, 0xf5, 0x92, 0x67, 0x90, 0x0e, 0x5a, 0xd2, 0x8b, 0x85, 0xba,
		0x59, 0x43, 0x74, 0x21, 0xe6, 0xea, 0xda, 0x04, 0x59, 0x64, 0x86, 0x04, 0x37, 0xb3, 0xe8, 0x05,
		0xa8, 0x93, 0xa5, 0x62, 0xbc, 0xa3, 0x39, 0xf4, 0x06, 0xcc, 0x8f, 0x3b, 0x55, 0xed, 0x8d, 0xa9,
		0x61, 0x43, 0x49, 0xf2, 0x0c, 0xe8, 0x19, 0x3c, 0x49, 0x4c, 0x5a, 0xaf, 0xfe, 0xc4, 0xe8, 0x19,
		0xf4, 0x18, 0xbe, 0x1e, 0x43, 0x8f, 0xd7, 0x48, 0xf7, 0x56, 0xa0, 0x17, 0xe4, 0x79, 0xf4, 0x10,
		0xf6, 0xc6, 0x10, 0x79, 0x17, 0x5a, 0x15, 0x53, 0x57, 0x5f, 0x9d, 0x86, 0xcb, 0x45, 0xbd, 0x62,
		0xca, 0x0b, 0xe8, 0x47, 0xf0, 0x83, 0x31, 0xb4, 0xce, 0x66, 0xd9, 0x1f, 0x1a, 0x8e, 0xb5, 0x18,
		0x83, 0x55, 0xb1, 0x26, 0x2f, 0x4e, 0x70, 0x26, 0x15, 0xfd, 0x45, 0x72, 0xe6, 0x96, 0x90, 0x0a,
		0xcf, 0x27, 0x6a, 0x11, 0xf5, 0x48, 0x2f, 0x16, 0x86, 0x3b, 0x91, 0xd1, 0xd7, 0xb0, 0x3b, 0xc6,
		0xc9, 0x61, 0x09, 0xab, 0x9a, 0x98, 0x58, 0x1d, 0x91, 0x58, 0x46, 0x8f, 0xe0, 0xc1, 0x38, 0x92,
		0xa2, 0x17, 0x4b, 0xaf, 0x35, 0xdc, 0xcf, 0x43, 0x6c, 0x8c, 0x4e, 0xb6, 0x75, 0xdd, 0x28, 0x57,
		0x4d, 0xab, 0xa2, 0x7f, 0xa7, 0xc9, 0x2b, 0x6c, 0x8c, 0x26, 0x9e, 0x54, 0x94, 0x2b, 0x79, 0x75,
		0x50, 0x8c, 0x07, 0x5e, 0x72, 0xa0, 0x1b, 0x0a, 0x3e, 0x95, 0xd7, 0x12, 0x6a, 0x6f, 0x50, 0xe8,
		0x7a, 0x4a, 0xe8, 0xfa, 0x24, 0xdb, 0xd1, 0x14, 0xac, 0x1e, 0xc5, 0x33, 0xbe, 0xce, 0xa6, 0xce,
		0x1d, 0xfe, 0x83, 0xcb, 0xc0, 0xbd, 0x2a, 0x2e, 0xf1, 0x7b, 0xb0, 0x1d, 0x9e, 0xdb, 0x90, 0x2a,
		0x18, 0xa1, 0xf6, 0x07, 0xf0, 0xc3, 0xc9, 0x28, 0x9d, 0x75, 0xa5, 0x88, 0x35, 0xa5, 0x70, 0xda,
		0xb9, 0x92, 0x4a, 0xf9, 0xbf, 0x49, 0x90, 0x57, 0x6d, 0xb7, 0x46, 0x1a, 0xd1, 0xef, 0xb8, 0x63,
		0xa3, 0xdc, 0x87, 0xc7, 0x13, 0xf4, 0xfb, 0x88, 0x78, 0x4f, 0xa0, 0xf2, 0xb1, 0xe4, 0xaa, 0xf1,
		0xca, 0x28, 0x9d, 0x18, 0xe3, 0x08, 0x62, 0x13, 0x15, 0xe7, 0x82, 0xff, 0x08, 0x3d, 0xd9, 0x26,
		0x44, 0xd9, 0xfd, 0x67, 0x9b, 0xf8, 0x58, 0xf2, 0x44, 0x9b, 0x38, 0xf8, 0x29, 0xac, 0xd7, 0xbc,
		0xe6, 0xb0, 0xaf, 0xf8, 0x83, 0x85, 0x68, 0x3b, 0x65, 0xf6, 0x19, 0x5b, 0x96, 0xbe, 0xdb, 0xbb,
		0x70, 0xe8, 0x65, 0xfb, 0x6c, 0xa7, 0xe6, 0x35, 0x77, 0xe3, 0xff, 0xc1, 0xdd, 0x76, 0xea, 0x8d,
		0xdd, 0x0b, 0x2f, 0xfc, 0x8f, 0xb0, 0xf8, 0x77, 0xee, 0xbe, 0xdd, 0x72, 0xde, 0xed, 0x9d, 0xcd,
		0x70, 0xdb, 0xd7, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x86, 0x60, 0x70, 0x2f, 0x8e, 0x1e, 0x00,
		0x00,
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
	// uber/cadence/api/v1/common.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x51, 0x6f, 0xdb, 0x36,
		0x17, 0xfd, 0x14, 0xc7, 0x4e, 0x7b, 0x9d, 0x26, 0xfa, 0x98, 0x35, 0x71, 0xd2, 0x75, 0x4b, 0x05,
		0x0c, 0xf5, 0x8a, 0x4d, 0x46, 0xdc, 0x97, 0x62, 0x45, 0x37, 0x38, 0xb6, 0x93, 0xa8, 0xcd, 0x6c,
		0x43, 0xf6, 0x1a, 0x74, 0x03, 0x26, 0xd0, 0x12, 0xe5, 0x72, 0x96, 0x48, 0x81, 0xa2, 0x9c, 0xf8,
		0x65, 0xd8, 0x2f, 0xd9, 0xc3, 0xfe, 0xd2, 0xfe, 0xd0, 0x20, 0x89, 0x8a, 0xed, 0xce, 0x41, 0xf7,
		0x30, 0xec, 0x8d, 0xbc, 0xe7, 0xdc, 0xc3, 0x43, 0xe2, 0xde, 0x2b, 0xc1, 0x71, 0x32, 0x26, 0xa2,
		0xe1, 0x62, 0x8f, 0x30, 0x97, 0x34, 0x70, 0x44, 0x1b, 0xb3, 0x93, 0x86, 0xcb, 0xc3, 0x90, 0x33,
		0x33, 0x12, 0x5c, 0x72, 0xb4, 0x97, 0x32, 0x4c, 0xc5, 0x30, 0x71, 0x44, 0xcd, 0xd9, 0xc9, 0xd1,
		0x67, 0x13, 0xce, 0x27, 0x01, 0x69, 0x64, 0x94, 0x71, 0xe2, 0x37, 0xbc, 0x44, 0x60, 0x49, 0x8b,
		0x24, 0xe3, 0x0d, 0xfc, 0xff, 0x8a, 0x8b, 0xa9, 0x1f, 0xf0, 0xeb, 0xee, 0x0d, 0x71, 0x93, 0x14,
		0x42, 0x9f, 0x43, 0xf5, 0x5a, 0x05, 0x1d, 0xea, 0xd5, 0xb4, 0x63, 0xad, 0x7e, 0xdf, 0x86, 0x22,
		0x64, 0x79, 0xe8, 0x21, 0x54, 0x44, 0xc2, 0x52, 0x6c, 0x23, 0xc3, 0xca, 0x22, 0x61, 0x96, 0x67,
		0x18, 0xb0, 0x5d, 0x88, 0x8d, 0xe6, 0x11, 0x41, 0x08, 0x36, 0x19, 0x0e, 0x89, 0x12, 0xc8, 0xd6,
		0x29, 0xa7, 0xe5, 0x4a, 0x3a, 0xa3, 0x72, 0x7e, 0x27, 0xe7, 0x31, 0x6c, 0x0d, 0xf0, 0x3c, 0xe0,
		0xd8, 0x4b, 0x61, 0x0f, 0x4b, 0x9c, 0xc1, 0xdb, 0x76, 0xb6, 0x36, 0x5e, 0xc2, 0xd6, 0x19, 0xa6,
		0x41, 0x22, 0x08, 0xda, 0x87, 0x8a, 0x20, 0x38, 0xe6, 0x4c, 0xe5, 0xab, 0x1d, 0xaa, 0xc1, 0x96,
		0x47, 0x24, 0xa6, 0x41, 0x9c, 0x39, 0xdc, 0xb6, 0x8b, 0xad, 0xf1, 0xbb, 0x06, 0x9b, 0xdf, 0x93,
		0x90, 0xa3, 0x57, 0x50, 0xf1, 0x29, 0x09, 0xbc, 0xb8, 0xa6, 0x1d, 0x97, 0xea, 0xd5, 0xe6, 0x17,
		0xe6, 0x9a, 0xf7, 0x33, 0x53, 0xaa, 0x79, 0x96, 0xf1, 0xba, 0x4c, 0x8a, 0xb9, 0xad, 0x92, 0x8e,
		0xae, 0xa0, 0xba, 0x14, 0x46, 0x3a, 0x94, 0xa6, 0x64, 0xae, 0x5c, 0xa4, 0x4b, 0xd4, 0x84, 0xf2,
		0x0c, 0x07, 0x09, 0xc9, 0x0c, 0x54, 0x9b, 0x9f, 0xae, 0x95, 0x57, 0xd7, 0xb4, 0x73, 0xea, 0x37,
		0x1b, 0x2f, 0x34, 0xe3, 0x0f, 0x0d, 0x2a, 0x17, 0x04, 0x7b, 0x44, 0xa0, 0xef, 0x3e, 0xb0, 0xf8,
		0x74, 0xad, 0x46, 0x4e, 0xfe, 0x6f, 0x4d, 0xfe, 0xa9, 0x81, 0x3e, 0x24, 0x58, 0xb8, 0xef, 0x5b,
		0x52, 0x0a, 0x3a, 0x4e, 0x24, 0x89, 0x91, 0x03, 0x3b, 0x94, 0x79, 0xe4, 0x86, 0x78, 0xce, 0x8a,
		0xed, 0x17, 0x6b, 0x55, 0x3f, 0x4c, 0x37, 0xad, 0x3c, 0x77, 0xf9, 0x1e, 0x0f, 0xe8, 0x72, 0xec,
		0xe8, 0x67, 0x40, 0x7f, 0x27, 0xfd, 0x8b, 0xb7, 0xf2, 0xe1, 0x5e, 0x07, 0x4b, 0x7c, 0x1a, 0xf0,
		0x31, 0x3a, 0x83, 0x07, 0x84, 0xb9, 0xdc, 0xa3, 0x6c, 0xe2, 0xc8, 0x79, 0x94, 0x17, 0xe8, 0x4e,
		0xf3, 0xc9, 0x5a, 0xad, 0xae, 0x62, 0xa6, 0x15, 0x6d, 0x6f, 0x93, 0xa5, 0xdd, 0x6d, 0x01, 0x6f,
		0x2c, 0x15, 0xf0, 0x20, 0x6f, 0x3a, 0x22, 0xde, 0x12, 0x11, 0x53, 0xce, 0x2c, 0xe6, 0xf3, 0x94,
		0x48, 0xc3, 0x28, 0x28, 0x1a, 0x21, 0x5d, 0xa3, 0xa7, 0xb0, 0xeb, 0x13, 0x2c, 0x13, 0x41, 0x9c,
		0x59, 0x4e, 0x55, 0x0d, 0xb7, 0xa3, 0xc2, 0x4a, 0xc0, 0x78, 0x03, 0x07, 0xc3, 0x24, 0x8a, 0xb8,
		0x90, 0xc4, 0x6b, 0x07, 0x94, 0x30, 0xa9, 0x90, 0x38, 0xed, 0xd5, 0x09, 0x77, 0x62, 0x6f, 0xaa,
		0x94, 0xcb, 0x13, 0x3e, 0xf4, 0xa6, 0xe8, 0x10, 0xee, 0xfd, 0x82, 0x67, 0x38, 0x03, 0x72, 0xcd,
		0xad, 0x74, 0x3f, 0xf4, 0xa6, 0xc6, 0x6f, 0x25, 0xa8, 0xda, 0x44, 0x8a, 0xf9, 0x80, 0x07, 0xd4,
		0x9d, 0xa3, 0x0e, 0xe8, 0x94, 0x51, 0x49, 0x71, 0xe0, 0x50, 0x26, 0x89, 0x98, 0xe1, 0xdc, 0x65,
		0xb5, 0x79, 0x68, 0xe6, 0xe3, 0xc5, 0x2c, 0xc6, 0x8b, 0xd9, 0x51, 0xe3, 0xc5, 0xde, 0x55, 0x29,
		0x96, 0xca, 0x40, 0x0d, 0xd8, 0x1b, 0x63, 0x77, 0xca, 0x7d, 0xdf, 0x71, 0x39, 0xf1, 0x7d, 0xea,
		0xa6, 0x36, 0xb3, 0xb3, 0x35, 0x1b, 0x29, 0xa8, 0xbd, 0x40, 0xd2, 0x63, 0x43, 0x7c, 0x43, 0xc3,
		0x24, 0x5c, 0x1c, 0x5b, 0xfa, 0xe8, 0xb1, 0x2a, 0xe5, 0xf6, 0xd8, 0x2f, 0x17, 0x2a, 0x58, 0x4a,
		0x12, 0x46, 0x32, 0xae, 0x6d, 0x1e, 0x6b, 0xf5, 0xf2, 0x2d, 0xb5, 0xa5, 0xc2, 0xe8, 0x15, 0x3c,
		0x62, 0x9c, 0x39, 0x22, 0xbd, 0x3a, 0x1e, 0x07, 0xc4, 0x21, 0x42, 0x70, 0xe1, 0xe4, 0x23, 0x25,
		0xae, 0x95, 0x8f, 0x4b, 0xf5, 0xfb, 0x76, 0x8d, 0x71, 0x66, 0x17, 0x8c, 0x6e, 0x4a, 0xb0, 0x73,
		0x1c, 0xbd, 0x86, 0x3d, 0x72, 0x13, 0xd1, 0xdc, 0xc8, 0xc2, 0x72, 0xe5, 0x63, 0x96, 0xd1, 0x22,
		0xab, 0x70, 0x6d, 0x84, 0x70, 0x60, 0xc5, 0x3c, 0xc8, 0x82, 0xe7, 0x82, 0x27, 0xd1, 0x00, 0x0b,
		0x49, 0xb3, 0xe1, 0xbc, 0x66, 0x60, 0xa2, 0x6f, 0xa1, 0x1c, 0x4b, 0x2c, 0xf3, 0x82, 0xdf, 0x69,
		0xd6, 0xd7, 0x16, 0xe9, 0xaa, 0xe0, 0x30, 0xe5, 0xdb, 0x79, 0x9a, 0x31, 0x83, 0x47, 0xab, 0x68,
		0x9b, 0x33, 0x9f, 0x4e, 0x94, 0x43, 0x74, 0x05, 0x3a, 0x2d, 0x60, 0x67, 0x92, 0xe2, 0x45, 0x6b,
		0x7f, 0xf5, 0x0f, 0x4e, 0xba, 0xb5, 0x6e, 0xef, 0xd2, 0x15, 0x20, 0x7e, 0x76, 0x0d, 0xdb, 0xcb,
		0xad, 0x83, 0x0e, 0xe1, 0x61, 0xb7, 0xd7, 0xee, 0x77, 0xac, 0xde, 0xb9, 0x33, 0x7a, 0x37, 0xe8,
		0x3a, 0x56, 0xef, 0x6d, 0xeb, 0xd2, 0xea, 0xe8, 0xff, 0x43, 0x47, 0xb0, 0xbf, 0x0a, 0x8d, 0x2e,
		0x6c, 0xeb, 0x6c, 0x64, 0x5f, 0xe9, 0x1a, 0xda, 0x07, 0xb4, 0x8a, 0xbd, 0x1e, 0xf6, 0x7b, 0xfa,
		0x06, 0xaa, 0xc1, 0x27, 0xab, 0xf1, 0x81, 0xdd, 0x1f, 0xf5, 0x9f, 0xeb, 0xa5, 0x67, 0xbf, 0xc2,
		0xde, 0x9a, 0xe7, 0x40, 0x4f, 0xe0, 0xb1, 0x35, 0xec, 0x5f, 0xb6, 0x46, 0x56, 0xbf, 0xe7, 0x9c,
		0xdb, 0xfd, 0x1f, 0x06, 0xce, 0x70, 0xd4, 0x1a, 0x2d, 0xfb, 0xb8, 0x93, 0x72, 0xd1, 0x6d, 0x5d,
		0x8e, 0x2e, 0xde, 0xe9, 0xda, 0xdd, 0x94, 0x8e, 0xdd, 0xb2, 0x7a, 0xdd, 0x8e, 0xbe, 0x71, 0xfa,
		0x13, 0x1c, 0xb8, 0x3c, 0x5c, 0xf7, 0x78, 0xa7, 0xd5, 0x76, 0xf6, 0x51, 0x1f, 0xa4, 0x75, 0x32,
		0xd0, 0x7e, 0x3c, 0x99, 0x50, 0xf9, 0x3e, 0x19, 0x9b, 0x2e, 0x0f, 0x1b, 0xcb, 0xbf, 0x00, 0x5f,
		0x53, 0x2f, 0x68, 0x4c, 0x78, 0xfe, 0x61, 0x57, 0xff, 0x03, 0x2f, 0x71, 0x44, 0x67, 0x27, 0xe3,
		0x4a, 0x16, 0x7b, 0xfe, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x0a, 0xaa, 0xd2, 0x33, 0x08,
		0x00, 0x00,
	},
	// uber/cadence/api/v1/tasklist.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6e, 0xdb, 0x36,
		0x14, 0x9e, 0xe2, 0xb4, 0x4b, 0x98, 0x25, 0xd5, 0xb8, 0xb5, 0x8d, 0xdd, 0x75, 0xcb, 0x74, 0x51,
		0x04, 0xc5, 0x26, 0xc1, 0x19, 0x76, 0xb5, 0x8b, 0xc1, 0xb1, 0x82, 0x55, 0xf0, 0x4f, 0x0d, 0x49,
		0x0d, 0xe0, 0x01, 0x03, 0x41, 0x89, 0xac, 0x4d, 0xe8, 0x87, 0x02, 0x49, 0x39, 0x31, 0xb0, 0xe7,
		0xd8, 0xc3, 0xec, 0x89, 0xf6, 0x18, 0x03, 0x29, 0xd9, 0xf3, 0x12, 0x6f, 0x77, 0xe4, 0xf9, 0xce,
		0x77, 0x7e, 0x3e, 0x9e, 0x43, 0xe0, 0xd4, 0x09, 0x15, 0x5e, 0x8a, 0x09, 0x2d, 0x53, 0xea, 0xe1,
		0x8a, 0x79, 0xab, 0xbe, 0xa7, 0xb0, 0xcc, 0x72, 0x26, 0x95, 0x5b, 0x09, 0xae, 0x38, 0xfc, 0x42,
		0xfb, 0xb8, 0xad, 0x8f, 0x8b, 0x2b, 0xe6, 0xae, 0xfa, 0xbd, 0xaf, 0x17, 0x9c, 0x2f, 0x72, 0xea,
		0x19, 0x97, 0xa4, 0xfe, 0xe8, 0x91, 0x5a, 0x60, 0xc5, 0x78, 0xd9, 0x90, 0x7a, 0xdf, 0x3c, 0xc4,
		0x15, 0x2b, 0xa8, 0x54, 0xb8, 0xa8, 0x5a, 0x87, 0x47, 0x01, 0xee, 0x04, 0xae, 0x2a, 0x2a, 0x64,
		0x83, 0x3b, 0x1f, 0xc0, 0x51, 0x8c, 0x65, 0x36, 0x66, 0x52, 0x41, 0x08, 0x0e, 0x4b, 0x5c, 0xd0,
		0x73, 0xeb, 0xc2, 0xba, 0x3c, 0x0e, 0xcd, 0x19, 0xfe, 0x08, 0x0e, 0x33, 0x56, 0x92, 0xf3, 0x83,
		0x0b, 0xeb, 0xf2, 0xec, 0xea, 0x5b, 0x77, 0x4f, 0x91, 0xee, 0x26, 0xc0, 0x88, 0x95, 0x24, 0x34,
		0xee, 0x0e, 0x06, 0xf6, 0xc6, 0x3a, 0xa1, 0x0a, 0x13, 0xac, 0x30, 0x9c, 0x80, 0x2f, 0x0b, 0x7c,
		0x8f, 0x74, 0xdb, 0x12, 0x55, 0x54, 0x20, 0x49, 0x53, 0x5e, 0x12, 0x93, 0xee, 0xe4, 0xea, 0x2b,
		0xb7, 0xa9, 0xd4, 0xdd, 0x54, 0xea, 0xfa, 0xbc, 0x4e, 0x72, 0x7a, 0x8b, 0xf3, 0x9a, 0x86, 0x9f,
		0x17, 0xf8, 0x5e, 0x07, 0x94, 0x33, 0x2a, 0x22, 0x43, 0x73, 0x3e, 0x80, 0xee, 0x26, 0xc5, 0x0c,
		0x0b, 0xc5, 0xb4, 0x2a, 0xdb, 0x5c, 0x36, 0xe8, 0x64, 0x74, 0xdd, 0x76, 0xa2, 0x8f, 0xf0, 0x0d,
		0x78, 0xc6, 0xef, 0x4a, 0x2a, 0xd0, 0x92, 0x4b, 0x85, 0x4c, 0x9f, 0x07, 0x06, 0x3d, 0x35, 0xe6,
		0x77, 0x5c, 0xaa, 0x29, 0x2e, 0xa8, 0xf3, 0x97, 0x05, 0xce, 0x36, 0x71, 0x23, 0x85, 0x55, 0x2d,
		0xe1, 0x77, 0x00, 0x26, 0x38, 0xcd, 0x72, 0xbe, 0x40, 0x29, 0xaf, 0x4b, 0x85, 0x96, 0xac, 0x54,
		0x26, 0x76, 0x27, 0xb4, 0x5b, 0x64, 0xa8, 0x81, 0x77, 0xac, 0x54, 0xf0, 0x35, 0x00, 0x82, 0x62,
		0x82, 0x72, 0xba, 0xa2, 0xb9, 0xc9, 0xd1, 0x09, 0x8f, 0xb5, 0x65, 0xac, 0x0d, 0xf0, 0x15, 0x38,
		0xc6, 0x69, 0xd6, 0xa2, 0x1d, 0x83, 0x1e, 0xe1, 0x34, 0x6b, 0xc0, 0x37, 0xe0, 0x99, 0xc0, 0x8a,
		0xee, 0xaa, 0x73, 0x78, 0x61, 0x5d, 0x5a, 0xe1, 0xa9, 0x36, 0x6f, 0x7b, 0x87, 0x3e, 0x38, 0xd5,
		0x32, 0x22, 0x46, 0x50, 0x92, 0xf3, 0x34, 0x3b, 0x7f, 0x62, 0x34, 0xbc, 0xf8, 0xcf, 0xe7, 0x09,
		0xfc, 0x6b, 0xed, 0x17, 0x9e, 0x68, 0x5a, 0x40, 0xcc, 0xc5, 0xf9, 0x19, 0x9c, 0xec, 0x60, 0xb0,
		0x0b, 0x8e, 0xa4, 0xc2, 0x42, 0x21, 0x46, 0xda, 0xe6, 0x3e, 0x35, 0xf7, 0x80, 0xc0, 0xe7, 0xe0,
		0x29, 0x2d, 0x89, 0x06, 0x9a, 0x7e, 0x9e, 0xd0, 0x92, 0x04, 0xc4, 0xf9, 0xc3, 0x02, 0x60, 0xc6,
		0xf3, 0x9c, 0x8a, 0xa0, 0xfc, 0xc8, 0xa1, 0x0f, 0xec, 0x1c, 0x4b, 0x85, 0x70, 0x9a, 0x52, 0x29,
		0x91, 0x1e, 0xc5, 0xf6, 0x71, 0x7b, 0x8f, 0x1e, 0x37, 0xde, 0xcc, 0x69, 0x78, 0xa6, 0x39, 0x03,
		0x43, 0xd1, 0x46, 0xd8, 0x03, 0x47, 0x8c, 0xd0, 0x52, 0x31, 0xb5, 0x6e, 0x5f, 0x68, 0x7b, 0xdf,
		0xa7, 0x4f, 0x67, 0x8f, 0x3e, 0xce, 0x9f, 0x16, 0xe8, 0x46, 0x8a, 0xa5, 0xd9, 0xfa, 0xe6, 0x9e,
		0xa6, 0xb5, 0x1e, 0x8d, 0x81, 0x52, 0x82, 0x25, 0xb5, 0xa2, 0x12, 0xfe, 0x02, 0xec, 0x3b, 0x2e,
		0x32, 0x2a, 0xcc, 0x2c, 0x22, 0xbd, 0x83, 0x6d, 0x9d, 0xaf, 0xff, 0x77, 0xbe, 0xc3, 0xb3, 0x86,
		0xb6, 0x5d, 0x98, 0x18, 0x74, 0x65, 0xba, 0xa4, 0xa4, 0xce, 0x29, 0x52, 0x1c, 0x35, 0xea, 0xe9,
		0xb6, 0x79, 0xad, 0x4c, 0xed, 0x27, 0x57, 0xdd, 0xc7, 0x63, 0xdd, 0x6e, 0x70, 0xf8, 0x62, 0xc3,
		0x8d, 0x79, 0xa4, 0x99, 0x71, 0x43, 0x7c, 0xfb, 0x3b, 0xf8, 0x6c, 0x77, 0xa3, 0x60, 0x0f, 0xbc,
		0x88, 0x07, 0xd1, 0x08, 0x8d, 0x83, 0x28, 0x46, 0xa3, 0x60, 0xea, 0xa3, 0x60, 0x7a, 0x3b, 0x18,
		0x07, 0xbe, 0xfd, 0x09, 0xec, 0x82, 0xe7, 0x0f, 0xb0, 0xe9, 0xfb, 0x70, 0x32, 0x18, 0xdb, 0xd6,
		0x1e, 0x28, 0x8a, 0x83, 0xe1, 0x68, 0x6e, 0x1f, 0xec, 0x89, 0xe8, 0xcf, 0xa7, 0x83, 0x49, 0x30,
		0xb4, 0x3b, 0x6f, 0xc9, 0x3f, 0xd9, 0xe3, 0x75, 0x45, 0xff, 0xed, 0x1b, 0xcf, 0x67, 0x37, 0x3b,
		0xd9, 0x5f, 0x81, 0x97, 0x0f, 0x30, 0xff, 0x66, 0x18, 0x44, 0xc1, 0xfb, 0xa9, 0x6d, 0xed, 0x01,
		0x07, 0xc3, 0x38, 0xb8, 0x0d, 0xe2, 0xb9, 0x7d, 0x70, 0xfd, 0x1b, 0x78, 0x99, 0xf2, 0x62, 0x9f,
		0xda, 0xd7, 0xa7, 0xdb, 0xad, 0xd6, 0x8a, 0xcd, 0xac, 0x5f, 0xfb, 0x0b, 0xa6, 0x96, 0x75, 0xe2,
		0xa6, 0xbc, 0xf0, 0x76, 0xff, 0xd1, 0xef, 0x19, 0xc9, 0xbd, 0x05, 0x6f, 0xbe, 0xb6, 0xf6, 0x53,
		0xfd, 0x09, 0x57, 0x6c, 0xd5, 0x4f, 0x9e, 0x1a, 0xdb, 0x0f, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff,
		0x4e, 0xf8, 0xfc, 0xcc, 0x78, 0x05, 0x00, 0x00,
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
