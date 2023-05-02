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
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x59, 0x4b, 0x73, 0xdb, 0xc8,
		0x11, 0x0e, 0x48, 0x49, 0x96, 0x9a, 0x7a, 0x40, 0x23, 0xc9, 0xa2, 0xe5, 0x97, 0xcc, 0x5d, 0xdb,
		0x32, 0xb3, 0x92, 0x56, 0xf6, 0x7a, 0xbd, 0xb6, 0xe2, 0x38, 0x10, 0x08, 0x59, 0xb0, 0x29, 0x90,
		0x19, 0x82, 0x96, 0xb5, 0x95, 0x04, 0x05, 0x91, 0x23, 0x09, 0x31, 0x09, 0xb0, 0x80, 0xa1, 0x6d,
		0xdd, 0x53, 0x95, 0x73, 0x6e, 0xa9, 0x9c, 0xf2, 0x03, 0x92, 0x4a, 0xa5, 0x72, 0x4e, 0xa5, 0x2a,
		0x87, 0xdc, 0x72, 0xcd, 0x31, 0xf7, 0xfc, 0x8b, 0xd4, 0x0c, 0x06, 0x24, 0x28, 0x3e, 0x40, 0x27,
		0x55, 0xbb, 0x37, 0xa1, 0xe7, 0xfb, 0x1a, 0x3d, 0x3d, 0xdd, 0x5f, 0x0f, 0x21, 0xc8, 0xb5, 0x4f,
		0x88, 0xbf, 0x5d, 0xb3, 0xeb, 0xc4, 0xad, 0x91, 0x6d, 0xbb, 0xe5, 0x6c, 0xbf, 0xdf, 0xd9, 0xfe,
		0xe0, 0xf9, 0xef, 0x4e, 0x1b, 0xde, 0x87, 0xad, 0x96, 0xef, 0x51, 0x0f, 0x2d, 0x31, 0xcc, 0x96,
		0xc0, 0x6c, 0xd9, 0x2d, 0x67, 0xeb, 0xfd, 0xce, 0xda, 0xad, 0x33, 0xcf, 0x3b, 0x6b, 0x90, 0x6d,
		0x0e, 0x39, 0x69, 0x9f, 0x6e, 0xd7, 0xdb, 0xbe, 0x4d, 0x1d, 0xcf, 0x0d, 0x49, 0x6b, 0xb7, 0x2f,
		0xaf, 0x53, 0xa7, 0x49, 0x02, 0x6a, 0x37, 0x5b, 0x02, 0xb0, 0x3e, 0xe8, 0xcd, 0x35, 0xaf, 0xd9,
		0xec, 0xb8, 0x18, 0x18, 0x1b, 0xb5, 0x83, 0x77, 0x0d, 0x27, 0xa0, 0x21, 0x26, 0xf7, 0xc7, 0x69,
		0x58, 0x39, 0x12, 0xe1, 0x6a, 0x1f, 0x49, 0xad, 0xcd, 0x42, 0xd0, 0xdd, 0x53, 0x0f, 0x55, 0x01,
		0x45, 0xfb, 0xb0, 0x48, 0xb4, 0x92, 0x95, 0xd6, 0xa5, 0x8d, 0xcc, 0xc3, 0x7b, 0x5b, 0x03, 0xb6,
		0xb4, 0xd5, 0xe7, 0x07, 0x2f, 0x7e, 0xb8, 0x6c, 0x42, 0x8f, 0x61, 0x82, 0x5e, 0xb4, 0x48, 0x36,
		0xc5, 0x1d, 0xdd, 0x19, 0xe9, 0xc8, 0xbc, 0x68, 0x11, 0xcc, 0xe1, 0xe8, 0x29, 0x40, 0x40, 0x6d,
		0x9f, 0x5a, 0x2c, 0x0d, 0xd9, 0x34, 0x27, 0xaf, 0x6d, 0x85, 0x39, 0xda, 0x8a, 0x72, 0xb4, 0x65,
		0x46, 0x39, 0xc2, 0x33, 0x1c, 0xcd, 0x9e, 0x19, 0xb5, 0xd6, 0xf0, 0x02, 0x12, 0x52, 0x27, 0x92,
		0xa9, 0x1c, 0xcd, 0xa9, 0x26, 0xcc, 0x86, 0xd4, 0x80, 0xda, 0xb4, 0x1d, 0x64, 0x27, 0xd7, 0xa5,
		0x8d, 0xf9, 0x87, 0x3b, 0xe3, 0xed, 0x5e, 0x65, 0xcc, 0x0a, 0x27, 0xe2, 0x4c, 0xad, 0xfb, 0x80,
		0xee, 0xc2, 0xfc, 0xb9, 0x13, 0x50, 0xcf, 0xbf, 0xb0, 0x1a, 0xc4, 0x3d, 0xa3, 0xe7, 0xd9, 0xa9,
		0x75, 0x69, 0x23, 0x8d, 0xe7, 0x84, 0xb5, 0xc8, 0x8d, 0xe8, 0x67, 0xb0, 0xd2, 0xb2, 0x7d, 0xe2,
		0xd2, 0x6e, 0xfa, 0x2d, 0xc7, 0x3d, 0xf5, 0xb2, 0x57, 0xf8, 0x16, 0x36, 0x06, 0x46, 0x51, 0xe6,
		0x8c, 0x9e, 0x93, 0xc4, 0x4b, 0xad, 0x7e, 0x23, 0x52, 0x60, 0xbe, 0xeb, 0x96, 0x67, 0x66, 0x3a,
		0x31, 0x33, 0x73, 0x1d, 0x06, 0xcf, 0xce, 0x26, 0x4c, 0x34, 0x49, 0xd3, 0xcb, 0xce, 0x70, 0xe2,
		0xb5, 0x81, 0xf1, 0x1c, 0x92, 0xa6, 0x87, 0x39, 0x0c, 0x61, 0x58, 0x0c, 0x88, 0xed, 0xd7, 0xce,
		0x2d, 0x9b, 0x52, 0xdf, 0x39, 0x69, 0x53, 0x12, 0x64, 0x81, 0x73, 0xef, 0x0e, 0xe4, 0x56, 0x38,
		0x5a, 0xe9, 0x80, 0xb1, 0x1c, 0x5c, 0xb2, 0xa0, 0x22, 0x2c, 0xda, 0x6d, 0xea, 0x59, 0x3e, 0x09,
		0x08, 0xb5, 0x5a, 0x9e, 0xe3, 0xd2, 0x20, 0x9b, 0xe1, 0x3e, 0xd7, 0x07, 0xfa, 0xc4, 0x0c, 0x58,
		0xe6, 0x38, 0xbc, 0xc0, 0xa8, 0x31, 0x03, 0xba, 0x0e, 0x33, 0xac, 0x3d, 0x2c, 0xd6, 0x1f, 0xd9,
		0xd9, 0x75, 0x69, 0x63, 0x06, 0x4f, 0x33, 0x43, 0xd1, 0x09, 0x28, 0x5a, 0x85, 0x2b, 0x4e, 0x60,
		0xd5, 0x7c, 0xcf, 0xcd, 0xce, 0xad, 0x4b, 0x1b, 0xd3, 0x78, 0xca, 0x09, 0x54, 0xdf, 0x73, 0xd1,
		0x2e, 0x64, 0xda, 0xad, 0xba, 0x4d, 0x45, 0x81, 0xcd, 0x27, 0xa6, 0x11, 0x42, 0x38, 0xcf, 0xe1,
		0x2f, 0x41, 0x6e, 0xd9, 0x3e, 0x75, 0xf8, 0x31, 0xd4, 0x3c, 0xf7, 0xd4, 0x39, 0xcb, 0x2e, 0xac,
		0xa7, 0x37, 0x32, 0x0f, 0x5f, 0x8c, 0x57, 0x65, 0xec, 0x30, 0xd9, 0xa9, 0x87, 0x2e, 0x54, 0xee,
		0x41, 0x73, 0xa9, 0x7f, 0x81, 0x17, 0x5a, 0xbd, 0xd6, 0xb5, 0x3d, 0x58, 0x1e, 0x04, 0x44, 0x32,
		0xa4, 0xdf, 0x91, 0x0b, 0xde, 0xda, 0x33, 0x98, 0xfd, 0x89, 0x96, 0x61, 0xf2, 0xbd, 0xdd, 0x68,
		0x87, 0x5d, 0x3a, 0x83, 0xc3, 0x87, 0x67, 0xa9, 0x6f, 0xa4, 0xdc, 0x6f, 0x53, 0x70, 0xab, 0xbf,
		0xd2, 0xb9, 0x33, 0xa1, 0x5f, 0xe8, 0x59, 0x3c, 0x8b, 0xa1, 0x5e, 0xdc, 0x1c, 0xb8, 0x17, 0x53,
		0xa4, 0x36, 0x96, 0x64, 0x1b, 0xd6, 0xbb, 0x55, 0x29, 0x1a, 0xde, 0xb3, 0xba, 0xed, 0xeb, 0xb5,
		0xa9, 0x50, 0x8e, 0x6b, 0x7d, 0x09, 0x2e, 0x88, 0x00, 0xf0, 0x8d, 0x8e, 0x8b, 0x0a, 0x17, 0x01,
		0x4f, 0x8d, 0x1a, 0xda, 0x6b, 0x53, 0x74, 0x04, 0xd7, 0x79, 0x78, 0x43, 0xbc, 0xa7, 0x93, 0xbc,
		0xaf, 0x32, 0xf6, 0x00, 0xc7, 0xb9, 0x7f, 0x4a, 0xb0, 0x34, 0xa0, 0xfd, 0x58, 0x55, 0xd5, 0xbd,
		0xa6, 0xed, 0xb8, 0x96, 0x53, 0x17, 0x49, 0x9e, 0x0e, 0x0d, 0x7a, 0x1d, 0xdd, 0x86, 0x8c, 0x58,
		0x74, 0xed, 0x66, 0x94, 0x6f, 0x08, 0x4d, 0x86, 0xdd, 0x24, 0x43, 0x64, 0x38, 0xfd, 0xff, 0xca,
		0xf0, 0x1d, 0x98, 0x75, 0x5c, 0x87, 0x3a, 0x36, 0x25, 0x75, 0x16, 0xd7, 0x04, 0x57, 0xa0, 0x4c,
		0xc7, 0xa6, 0xd7, 0x73, 0xbf, 0x91, 0x60, 0x45, 0xfb, 0x48, 0x89, 0xef, 0xda, 0x8d, 0xef, 0x64,
		0x34, 0x5c, 0x8e, 0x29, 0xd5, 0x1f, 0xd3, 0xbf, 0x27, 0x61, 0xa9, 0x4c, 0xdc, 0xba, 0xe3, 0x9e,
		0x29, 0x35, 0xea, 0xbc, 0x77, 0xe8, 0x05, 0x8f, 0xe8, 0x36, 0x64, 0x6c, 0xf1, 0xdc, 0xcd, 0x32,
		0x44, 0x26, 0xbd, 0x8e, 0xf6, 0x61, 0xae, 0x03, 0x48, 0x9c, 0x3f, 0x91, 0x6b, 0x3e, 0x7f, 0x66,
		0xed, 0xd8, 0x13, 0x7a, 0x01, 0x93, 0x6c, 0x16, 0x84, 0x23, 0x68, 0xfe, 0xe1, 0x83, 0xc1, 0x22,
		0xdc, 0x1b, 0x21, 0x93, 0x7d, 0x82, 0x43, 0x1e, 0xd2, 0x61, 0xf1, 0x9c, 0xd8, 0x3e, 0x3d, 0x21,
		0x36, 0xb5, 0xea, 0x84, 0xda, 0x4e, 0x23, 0x10, 0x43, 0xe9, 0xc6, 0x10, 0x45, 0xbf, 0x68, 0x78,
		0x76, 0x1d, 0xcb, 0x1d, 0x5a, 0x21, 0x64, 0xa1, 0x57, 0xb0, 0xd4, 0xb0, 0x03, 0x6a, 0x75, 0xfd,
		0x71, 0x01, 0x9a, 0x4c, 0x14, 0xa0, 0x45, 0x46, 0x3b, 0x88, 0x58, 0x5c, 0x87, 0xf6, 0x81, 0x1b,
		0xc3, 0xae, 0x20, 0xf5, 0xd0, 0xd3, 0x54, 0xa2, 0xa7, 0x05, 0x46, 0xaa, 0x84, 0x1c, 0xee, 0x27,
		0x0b, 0x57, 0x6c, 0x4a, 0x49, 0xb3, 0x45, 0xf9, 0x98, 0x9a, 0xc4, 0xd1, 0x23, 0x7a, 0x00, 0x72,
		0xd3, 0xfe, 0xe8, 0x34, 0xdb, 0x4d, 0x4b, 0x98, 0x02, 0x3e, 0x72, 0x26, 0xf1, 0x82, 0xb0, 0x2b,
		0xc2, 0xcc, 0x66, 0x53, 0x50, 0x3b, 0x27, 0xf5, 0x76, 0x23, 0x8a, 0x64, 0x26, 0x79, 0x36, 0x75,
		0x18, 0x3c, 0x0e, 0x15, 0x16, 0xc8, 0xc7, 0x96, 0x13, 0xf6, 0x6c, 0xe8, 0x03, 0x12, 0x7d, 0xcc,
		0x77, 0x29, 0xdc, 0xc9, 0x0b, 0x98, 0xe5, 0x49, 0x39, 0xb5, 0x9d, 0x46, 0xdb, 0x27, 0x62, 0xb0,
		0x0c, 0x3e, 0xa6, 0xfd, 0x10, 0x83, 0x33, 0x8c, 0x21, 0x1e, 0xd0, 0x97, 0xb0, 0xcc, 0x1d, 0xb0,
		0x5a, 0x27, 0xbe, 0xe5, 0xd4, 0x89, 0x4b, 0x1d, 0x7a, 0x21, 0x66, 0x0b, 0x62, 0x6b, 0x47, 0x7c,
		0x49, 0x17, 0x2b, 0xb9, 0xbf, 0xa4, 0xe0, 0x9a, 0x28, 0x1f, 0xf5, 0xdc, 0x69, 0xd4, 0xbf, 0x93,
		0xc6, 0xfb, 0x22, 0xe6, 0x96, 0x35, 0x47, 0x5c, 0x8b, 0xe4, 0x0f, 0xb1, 0xcb, 0x18, 0x57, 0xa4,
		0xcb, 0x6d, 0x9a, 0xee, 0x6b, 0x53, 0xf4, 0x06, 0xc4, 0x9d, 0x43, 0x88, 0x6b, 0xcb, 0x6b, 0x38,
		0xb5, 0x0b, 0x5e, 0xe6, 0xf3, 0x43, 0x02, 0x0d, 0x95, 0x93, 0x0b, 0x6a, 0x99, 0xa3, 0xf1, 0x62,
		0xeb, 0xb2, 0x09, 0x5d, 0x85, 0xa9, 0x50, 0x1a, 0x79, 0x91, 0xcf, 0x60, 0xf1, 0x94, 0xfb, 0x47,
		0xaa, 0x23, 0x0b, 0x05, 0x52, 0x73, 0x82, 0x28, 0x5f, 0x9d, 0x6e, 0x95, 0x92, 0xbb, 0x35, 0x22,
		0xf6, 0x74, 0x6b, 0x7f, 0x25, 0xa6, 0x3e, 0xb5, 0x12, 0x9f, 0xc3, 0x6c, 0x4f, 0x53, 0x25, 0xdf,
		0x5d, 0x33, 0xc1, 0xe0, 0x86, 0x9a, 0xe8, 0x6d, 0x28, 0x0c, 0xab, 0x9e, 0xef, 0x9c, 0x39, 0xae,
		0xdd, 0xb0, 0x2e, 0x05, 0x99, 0x2c, 0x01, 0x2b, 0x11, 0xb5, 0x12, 0x0f, 0x36, 0xf7, 0xd7, 0x14,
		0x5c, 0x8b, 0x64, 0xab, 0xe8, 0xd5, 0xec, 0x46, 0xc1, 0x09, 0x5a, 0x36, 0xad, 0x9d, 0x8f, 0xa7,
		0xb2, 0xdf, 0x7f, 0xba, 0x7e, 0x01, 0xb7, 0x7a, 0x23, 0xb0, 0xbc, 0x53, 0x8b, 0x9e, 0x3b, 0x81,
		0x15, 0xcf, 0xe2, 0x68, 0x87, 0x6b, 0x3d, 0x11, 0x95, 0x4e, 0xcd, 0x73, 0x27, 0x10, 0xda, 0x84,
		0x6e, 0x02, 0xf0, 0xdb, 0x03, 0xf5, 0xde, 0x91, 0xb0, 0x0a, 0x67, 0x31, 0xbf, 0xee, 0x98, 0xcc,
		0x90, 0x7b, 0x05, 0x99, 0xf8, 0x85, 0x72, 0x17, 0xa6, 0xc4, 0x9d, 0x54, 0xe2, 0x77, 0xba, 0xcf,
		0x12, 0xee, 0xa4, 0xfc, 0xba, 0x2e, 0x28, 0xb9, 0x3f, 0xa5, 0x60, 0xbe, 0x77, 0x09, 0xdd, 0x87,
		0x85, 0x13, 0xc7, 0xb5, 0xfd, 0x0b, 0xab, 0x76, 0x4e, 0x6a, 0xef, 0x82, 0x76, 0x53, 0x1c, 0xc2,
		0x7c, 0x68, 0x56, 0x85, 0x15, 0xad, 0xc0, 0x94, 0xdf, 0x76, 0xa3, 0x21, 0x3a, 0x83, 0x27, 0xfd,
		0x36, 0xbb, 0x6d, 0x3c, 0x87, 0xeb, 0xa7, 0x8e, 0x1f, 0xb0, 0xc1, 0x13, 0x16, 0xbb, 0x55, 0xf3,
		0x9a, 0xad, 0x06, 0xe9, 0xe9, 0xe4, 0x2c, 0x87, 0x44, 0xed, 0xa0, 0x46, 0x00, 0x4e, 0x9f, 0xad,
		0xf9, 0xc4, 0xee, 0x9c, 0x4d, 0x72, 0x2a, 0x33, 0x02, 0x2f, 0xe4, 0x74, 0x8e, 0x0b, 0xac, 0xe3,
		0x9e, 0x8d, 0x5b, 0xa6, 0xb3, 0x11, 0x81, 0x3b, 0xb8, 0x05, 0xc0, 0x2f, 0xfa, 0xd4, 0x3e, 0x69,
		0x84, 0xd3, 0x69, 0x1a, 0xc7, 0x2c, 0xf9, 0x3f, 0x4b, 0xb0, 0x3c, 0x68, 0xf6, 0xa2, 0x1c, 0xdc,
		0x2a, 0x6b, 0x46, 0x41, 0x37, 0x5e, 0x5a, 0x8a, 0x6a, 0xea, 0x6f, 0x74, 0xf3, 0xd8, 0xaa, 0x98,
		0x8a, 0xa9, 0x59, 0xba, 0xf1, 0x46, 0x29, 0xea, 0x05, 0xf9, 0x07, 0xe8, 0x73, 0x58, 0x1f, 0x82,
		0xa9, 0xa8, 0x07, 0x5a, 0xa1, 0x5a, 0xd4, 0x0a, 0xb2, 0x34, 0xc2, 0x53, 0xc5, 0x54, 0xb0, 0xa9,
		0x15, 0xe4, 0x14, 0xfa, 0x21, 0xdc, 0x1f, 0x82, 0x51, 0x15, 0x43, 0xd5, 0x8a, 0x16, 0xd6, 0x7e,
		0x5a, 0xd5, 0x2a, 0x0c, 0x9c, 0xce, 0xff, 0xaa, 0x1b, 0x73, 0x8f, 0x02, 0xc5, 0xdf, 0x54, 0xd0,
		0x54, 0xbd, 0xa2, 0x97, 0x8c, 0x51, 0x31, 0x5f, 0xc2, 0x0c, 0x89, 0xf9, 0x32, 0x2a, 0x8a, 0x39,
		0xff, 0xeb, 0x54, 0xf7, 0x3b, 0x80, 0x5e, 0xc7, 0xa4, 0xdd, 0xd1, 0xdc, 0xcf, 0x61, 0xfd, 0xa8,
		0x84, 0x5f, 0xef, 0x17, 0x4b, 0x47, 0x96, 0x5e, 0xb0, 0xb0, 0x56, 0xad, 0x68, 0x56, 0xb9, 0x54,
		0xd4, 0xd5, 0xe3, 0x58, 0x24, 0xdf, 0xc0, 0x57, 0x43, 0x51, 0x4a, 0x91, 0x59, 0x0b, 0xd5, 0x72,
		0x51, 0x57, 0xd9, 0x5b, 0xf7, 0x15, 0xbd, 0xa8, 0x15, 0xac, 0x92, 0x51, 0x3c, 0x96, 0x25, 0xf4,
		0x05, 0x6c, 0x8c, 0xcb, 0x94, 0x53, 0x68, 0x13, 0x1e, 0x0c, 0x45, 0x63, 0xed, 0x95, 0xa6, 0x9a,
		0x31, 0x78, 0x1a, 0xed, 0xc0, 0xe6, 0x50, 0xb8, 0xa9, 0xe1, 0x43, 0xdd, 0xe0, 0x09, 0xdd, 0xb7,
		0x70, 0xd5, 0x30, 0x74, 0xe3, 0xa5, 0x3c, 0x91, 0xff, 0xbd, 0x04, 0x8b, 0x7d, 0xc3, 0x08, 0xdd,
		0x86, 0xeb, 0x65, 0x05, 0x6b, 0x86, 0x69, 0xa9, 0xc5, 0xd2, 0xa0, 0x04, 0x0c, 0x01, 0x28, 0x7b,
		0x8a, 0x51, 0x28, 0x19, 0xb2, 0x84, 0xee, 0x41, 0x6e, 0x10, 0x40, 0xd4, 0x82, 0x28, 0x0d, 0x39,
		0x85, 0xee, 0xc0, 0xcd, 0x41, 0xb8, 0x4e, 0xb4, 0x72, 0x3a, 0xff, 0x9f, 0x14, 0xdc, 0x18, 0xf5,
		0xb9, 0x81, 0x55, 0x60, 0x67, 0xdb, 0xda, 0x5b, 0x4d, 0xad, 0x9a, 0xec, 0xcc, 0x43, 0x7f, 0xec,
		0xe4, 0xab, 0x95, 0x58, 0xe4, 0xf1, 0x94, 0x0e, 0x01, 0xab, 0xa5, 0xc3, 0x72, 0x51, 0x33, 0x79,
		0x35, 0xe5, 0xe1, 0x5e, 0x12, 0x3c, 0x3c, 0x60, 0x39, 0xd5, 0x73, 0xb6, 0xc3, 0x5c, 0xf3, 0x7d,
		0xb3, 0x56, 0x40, 0x5b, 0x90, 0x4f, 0x42, 0x77, 0xb2, 0x50, 0x90, 0x27, 0xd0, 0x57, 0xf0, 0x65,
		0x72, 0xe0, 0x86, 0xa9, 0x1b, 0x55, 0xad, 0x60, 0x29, 0x15, 0xcb, 0xd0, 0x8e, 0xe4, 0xc9, 0x71,
		0xb6, 0x6b, 0xea, 0x87, 0xac, 0x3e, 0xab, 0xa6, 0x3c, 0x95, 0xff, 0x9b, 0x04, 0x57, 0x55, 0xcf,
		0xa5, 0x8e, 0xdb, 0x26, 0x4a, 0x60, 0x90, 0x0f, 0x7a, 0x78, 0xcf, 0xf1, 0x7c, 0x74, 0x17, 0xee,
		0x44, 0xfe, 0x85, 0x7b, 0x4b, 0x37, 0x74, 0x53, 0x57, 0xcc, 0x12, 0x8e, 0xe5, 0x77, 0x24, 0x8c,
		0x35, 0x64, 0x41, 0xc3, 0x61, 0x5e, 0x87, 0xc3, 0xb0, 0x66, 0xe2, 0x63, 0x51, 0x0a, 0xa1, 0xc2,
		0x0c, 0xc7, 0xaa, 0x98, 0xf5, 0xb7, 0xe8, 0x7f, 0x39, 0x9d, 0xff, 0x83, 0x04, 0x19, 0xf1, 0x1b,
		0x95, 0xff, 0x84, 0xc9, 0xc2, 0x32, 0xdb, 0x60, 0xa9, 0x6a, 0x5a, 0xe6, 0x71, 0x59, 0xeb, 0xad,
		0xe1, 0x9e, 0x15, 0x2e, 0x0f, 0x96, 0x59, 0x0a, 0xb3, 0x13, 0x2a, 0x49, 0x2f, 0x40, 0xbc, 0x85,
		0x61, 0x38, 0x58, 0x4e, 0x8d, 0xc4, 0x84, 0x7e, 0xd2, 0x68, 0x0d, 0xae, 0xf6, 0x60, 0x0e, 0x34,
		0x05, 0x9b, 0x7b, 0x9a, 0x62, 0xca, 0x13, 0xf9, 0xdf, 0x49, 0x70, 0x2d, 0x52, 0x42, 0x93, 0x0d,
		0x56, 0xa7, 0x49, 0xea, 0xa5, 0x36, 0x55, 0xed, 0x76, 0x40, 0xd0, 0x03, 0xb8, 0xdb, 0xd1, 0x30,
		0x53, 0xa9, 0xbc, 0xee, 0x9e, 0x95, 0xa5, 0x2a, 0xac, 0xb9, 0xbb, 0xbb, 0x49, 0x84, 0x8a, 0x10,
		0x64, 0x09, 0xdd, 0x87, 0xcf, 0x46, 0x43, 0xb1, 0x56, 0xd1, 0x4c, 0x39, 0x95, 0xff, 0x57, 0x06,
		0x56, 0xe3, 0xc1, 0xb1, 0x8b, 0x3e, 0xa9, 0x87, 0xa1, 0xdd, 0x83, 0x5c, 0xaf, 0x13, 0xa1, 0x73,
		0x97, 0xe3, 0xda, 0x81, 0xcd, 0x11, 0xb8, 0xaa, 0x71, 0xa0, 0x18, 0x05, 0xf6, 0x1c, 0x81, 0x64,
		0x09, 0xbd, 0x80, 0xdd, 0x11, 0x94, 0x3d, 0xa5, 0xd0, 0xcd, 0x72, 0x67, 0xe2, 0x28, 0xa6, 0x89,
		0xf5, 0xbd, 0xaa, 0xa9, 0x55, 0xe4, 0x14, 0xd2, 0x40, 0x49, 0x70, 0xd0, 0xab, 0x43, 0x03, 0xdd,
		0xa4, 0xd1, 0x53, 0x78, 0x9c, 0x14, 0x47, 0x58, 0x32, 0xfa, 0xa1, 0x86, 0xe3, 0xd4, 0x09, 0xf4,
		0x0c, 0xbe, 0x4e, 0xa0, 0x8a, 0x37, 0xf7, 0x71, 0x27, 0xd1, 0x2e, 0x3c, 0x49, 0x8c, 0x5e, 0x2d,
		0xe1, 0x82, 0x75, 0xa8, 0xe0, 0xd7, 0xbd, 0xe4, 0x29, 0xa4, 0x83, 0x96, 0xf4, 0x62, 0xa1, 0x6e,
		0xd6, 0x00, 0x5d, 0x88, 0xb9, 0xba, 0x32, 0x46, 0x16, 0x99, 0x21, 0xc1, 0xcd, 0x34, 0x7a, 0x09,
		0xea, 0x78, 0xa9, 0x18, 0xed, 0x68, 0x06, 0xbd, 0x05, 0xf3, 0xd3, 0x4e, 0x55, 0x7b, 0x6b, 0x6a,
		0xd8, 0x50, 0x92, 0x3c, 0x03, 0x7a, 0x0e, 0x4f, 0x13, 0x93, 0xd6, 0xab, 0x3f, 0x31, 0x7a, 0x06,
		0x3d, 0x81, 0x47, 0x23, 0xe8, 0xf1, 0x1a, 0xe9, 0xde, 0x0a, 0xf4, 0x82, 0x3c, 0x8b, 0x1e, 0xc3,
		0xce, 0x08, 0x22, 0xef, 0x42, 0xab, 0x62, 0xea, 0xea, 0xeb, 0xe3, 0x70, 0xb9, 0xa8, 0x57, 0x4c,
		0x79, 0x0e, 0xfd, 0x04, 0x7e, 0x34, 0x82, 0xd6, 0xd9, 0x2c, 0xfb, 0x43, 0xc3, 0xb1, 0x16, 0x63,
		0xb0, 0x2a, 0xd6, 0xe4, 0xf9, 0x31, 0xce, 0xa4, 0xa2, 0xbf, 0x4c, 0xce, 0xdc, 0x02, 0x52, 0xe1,
		0xc5, 0x58, 0x2d, 0xa2, 0x1e, 0xe8, 0xc5, 0xc2, 0x60, 0x27, 0x32, 0x7a, 0x04, 0xdb, 0x23, 0x9c,
		0xec, 0x97, 0xb0, 0xaa, 0x89, 0x89, 0xd5, 0x11, 0x89, 0x45, 0xf4, 0x35, 0x3c, 0x1c, 0x45, 0x52,
		0xf4, 0x62, 0xe9, 0x8d, 0x86, 0x2f, 0xf3, 0x10, 0x1b, 0xa3, 0xe3, 0x6d, 0x5d, 0x37, 0xca, 0x55,
		0xd3, 0xaa, 0xe8, 0xdf, 0x6a, 0xf2, 0x12, 0x1b, 0xa3, 0x89, 0x27, 0x15, 0xe5, 0x4a, 0x5e, 0xee,
		0x17, 0xe3, 0xbe, 0x97, 0xec, 0xe9, 0x86, 0x82, 0x8f, 0xe5, 0x95, 0x84, 0xda, 0xeb, 0x17, 0xba,
		0x9e, 0x12, 0xba, 0x3a, 0xce, 0x76, 0x34, 0x05, 0xab, 0x07, 0xf1, 0x8c, 0xaf, 0xb2, 0xa9, 0x73,
		0x87, 0x7f, 0x70, 0xe9, 0xbb, 0x57, 0xc5, 0x25, 0x7e, 0x07, 0x36, 0xc3, 0x73, 0x1b, 0x50, 0x05,
		0x43, 0xd4, 0x7e, 0x0f, 0x7e, 0x3c, 0x1e, 0xa5, 0xb3, 0xae, 0x14, 0xb1, 0xa6, 0x14, 0x8e, 0x3b,
		0x57, 0x52, 0x29, 0xff, 0x77, 0x09, 0xf2, 0xaa, 0xed, 0xd6, 0x48, 0x23, 0xfa, 0x1e, 0x3b, 0x32,
		0xca, 0x5d, 0x78, 0x32, 0x46, 0xbf, 0x0f, 0x89, 0xf7, 0x08, 0x2a, 0x9f, 0x4a, 0xae, 0x1a, 0xaf,
		0x8d, 0xd2, 0x91, 0x31, 0x8a, 0x20, 0x36, 0x51, 0x71, 0xce, 0xf8, 0xc7, 0xe4, 0xf1, 0x36, 0x21,
		0xca, 0xee, 0x7f, 0xdb, 0xc4, 0xa7, 0x92, 0xc7, 0xda, 0xc4, 0xde, 0xcf, 0x61, 0xb5, 0xe6, 0x35,
		0x07, 0xfd, 0x8a, 0xdf, 0x9b, 0x8b, 0xb6, 0x53, 0x66, 0x3f, 0x63, 0xcb, 0xd2, 0xb7, 0x3b, 0x67,
		0x0e, 0x3d, 0x6f, 0x9f, 0x6c, 0xd5, 0xbc, 0xe6, 0x76, 0xfc, 0x3f, 0xb1, 0x9b, 0x4e, 0xbd, 0xb1,
		0x7d, 0xe6, 0x85, 0xff, 0xd9, 0x15, 0xff, 0x96, 0xdd, 0xb5, 0x5b, 0xce, 0xfb, 0x9d, 0x93, 0x29,
		0x6e, 0x7b, 0xf4, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xae, 0x0a, 0xa0, 0x56, 0x1e, 0x00,
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
		0x14, 0x9e, 0xe2, 0xb4, 0x4b, 0xe8, 0xd9, 0xd5, 0xb8, 0xb5, 0x8d, 0xdd, 0x75, 0xf3, 0x74, 0x51,
		0x04, 0xc5, 0x26, 0xc1, 0x19, 0x76, 0xb5, 0x8b, 0xc1, 0xb1, 0x83, 0x55, 0xb0, 0xe3, 0x1a, 0x92,
		0x1a, 0x20, 0x03, 0x06, 0x8e, 0x12, 0x59, 0x9b, 0xd0, 0x0f, 0x05, 0x92, 0x72, 0xe2, 0x17, 0xd9,
		0xc3, 0xec, 0x89, 0xf6, 0x18, 0x03, 0x29, 0xd9, 0xf3, 0x12, 0x6f, 0x77, 0xe4, 0xf9, 0xce, 0x77,
		0x7e, 0x3e, 0x9e, 0x43, 0xe0, 0x54, 0x31, 0x15, 0x5e, 0x82, 0x09, 0x2d, 0x12, 0xea, 0xe1, 0x92,
		0x79, 0xeb, 0xa1, 0xa7, 0xb0, 0x4c, 0x33, 0x26, 0x95, 0x5b, 0x0a, 0xae, 0x38, 0xfc, 0x42, 0xfb,
		0xb8, 0x8d, 0x8f, 0x8b, 0x4b, 0xe6, 0xae, 0x87, 0xfd, 0xaf, 0x97, 0x9c, 0x2f, 0x33, 0xea, 0x19,
		0x97, 0xb8, 0xfa, 0xe8, 0x91, 0x4a, 0x60, 0xc5, 0x78, 0x51, 0x93, 0xfa, 0xdf, 0x3c, 0xc4, 0x15,
		0xcb, 0xa9, 0x54, 0x38, 0x2f, 0x1b, 0x87, 0x47, 0x01, 0xee, 0x04, 0x2e, 0x4b, 0x2a, 0x64, 0x8d,
		0x3b, 0x1f, 0xc0, 0x49, 0x84, 0x65, 0x3a, 0x63, 0x52, 0x41, 0x08, 0x8e, 0x0b, 0x9c, 0xd3, 0x33,
		0x6b, 0x60, 0x9d, 0x9f, 0x06, 0xe6, 0x0c, 0x7f, 0x04, 0xc7, 0x29, 0x2b, 0xc8, 0xd9, 0xd1, 0xc0,
		0x3a, 0xef, 0x5e, 0x7c, 0xeb, 0x1e, 0x28, 0xd2, 0xdd, 0x06, 0x98, 0xb2, 0x82, 0x04, 0xc6, 0xdd,
		0xc1, 0xc0, 0xde, 0x5a, 0xaf, 0xa9, 0xc2, 0x04, 0x2b, 0x0c, 0xaf, 0xc1, 0x97, 0x39, 0xbe, 0x47,
		0xba, 0x6d, 0x89, 0x4a, 0x2a, 0x90, 0xa4, 0x09, 0x2f, 0x88, 0x49, 0xd7, 0xbe, 0xf8, 0xca, 0xad,
		0x2b, 0x75, 0xb7, 0x95, 0xba, 0x13, 0x5e, 0xc5, 0x19, 0xbd, 0xc1, 0x59, 0x45, 0x83, 0xcf, 0x73,
		0x7c, 0xaf, 0x03, 0xca, 0x05, 0x15, 0xa1, 0xa1, 0x39, 0x1f, 0x40, 0x6f, 0x9b, 0x62, 0x81, 0x85,
		0x62, 0x5a, 0x95, 0x5d, 0x2e, 0x1b, 0xb4, 0x52, 0xba, 0x69, 0x3a, 0xd1, 0x47, 0xf8, 0x06, 0x3c,
		0xe3, 0x77, 0x05, 0x15, 0x68, 0xc5, 0xa5, 0x42, 0xa6, 0xcf, 0x23, 0x83, 0x76, 0x8c, 0xf9, 0x1d,
		0x97, 0x6a, 0x8e, 0x73, 0xea, 0xfc, 0x65, 0x81, 0xee, 0x36, 0x6e, 0xa8, 0xb0, 0xaa, 0x24, 0xfc,
		0x0e, 0xc0, 0x18, 0x27, 0x69, 0xc6, 0x97, 0x28, 0xe1, 0x55, 0xa1, 0xd0, 0x8a, 0x15, 0xca, 0xc4,
		0x6e, 0x05, 0x76, 0x83, 0x8c, 0x35, 0xf0, 0x8e, 0x15, 0x0a, 0xbe, 0x06, 0x40, 0x50, 0x4c, 0x50,
		0x46, 0xd7, 0x34, 0x33, 0x39, 0x5a, 0xc1, 0xa9, 0xb6, 0xcc, 0xb4, 0x01, 0xbe, 0x02, 0xa7, 0x38,
		0x49, 0x1b, 0xb4, 0x65, 0xd0, 0x13, 0x9c, 0xa4, 0x35, 0xf8, 0x06, 0x3c, 0x13, 0x58, 0xd1, 0x7d,
		0x75, 0x8e, 0x07, 0xd6, 0xb9, 0x15, 0x74, 0xb4, 0x79, 0xd7, 0x3b, 0x9c, 0x80, 0x8e, 0x96, 0x11,
		0x31, 0x82, 0xe2, 0x8c, 0x27, 0xe9, 0xd9, 0x13, 0xa3, 0xe1, 0xe0, 0x3f, 0x9f, 0xc7, 0x9f, 0x5c,
		0x6a, 0xbf, 0xa0, 0xad, 0x69, 0x3e, 0x31, 0x17, 0xe7, 0x67, 0xd0, 0xde, 0xc3, 0x60, 0x0f, 0x9c,
		0x48, 0x85, 0x85, 0x42, 0x8c, 0x34, 0xcd, 0x7d, 0x6a, 0xee, 0x3e, 0x81, 0xcf, 0xc1, 0x53, 0x5a,
		0x10, 0x0d, 0xd4, 0xfd, 0x3c, 0xa1, 0x05, 0xf1, 0x89, 0xf3, 0x87, 0x05, 0xc0, 0x82, 0x67, 0x19,
		0x15, 0x7e, 0xf1, 0x91, 0xc3, 0x09, 0xb0, 0x33, 0x2c, 0x15, 0xc2, 0x49, 0x42, 0xa5, 0x44, 0x7a,
		0x14, 0x9b, 0xc7, 0xed, 0x3f, 0x7a, 0xdc, 0x68, 0x3b, 0xa7, 0x41, 0x57, 0x73, 0x46, 0x86, 0xa2,
		0x8d, 0xb0, 0x0f, 0x4e, 0x18, 0xa1, 0x85, 0x62, 0x6a, 0xd3, 0xbc, 0xd0, 0xee, 0x7e, 0x48, 0x9f,
		0xd6, 0x01, 0x7d, 0x9c, 0x3f, 0x2d, 0xd0, 0x0b, 0x15, 0x4b, 0xd2, 0xcd, 0xd5, 0x3d, 0x4d, 0x2a,
		0x3d, 0x1a, 0x23, 0xa5, 0x04, 0x8b, 0x2b, 0x45, 0x25, 0xfc, 0x05, 0xd8, 0x77, 0x5c, 0xa4, 0x54,
		0x98, 0x59, 0x44, 0x7a, 0x07, 0x9b, 0x3a, 0x5f, 0xff, 0xef, 0x7c, 0x07, 0xdd, 0x9a, 0xb6, 0x5b,
		0x98, 0x08, 0xf4, 0x64, 0xb2, 0xa2, 0xa4, 0xca, 0x28, 0x52, 0x1c, 0xd5, 0xea, 0xe9, 0xb6, 0x79,
		0xa5, 0x4c, 0xed, 0xed, 0x8b, 0xde, 0xe3, 0xb1, 0x6e, 0x36, 0x38, 0x78, 0xb1, 0xe5, 0x46, 0x3c,
		0xd4, 0xcc, 0xa8, 0x26, 0xbe, 0xfd, 0x1d, 0x7c, 0xb6, 0xbf, 0x51, 0xb0, 0x0f, 0x5e, 0x44, 0xa3,
		0x70, 0x8a, 0x66, 0x7e, 0x18, 0xa1, 0xa9, 0x3f, 0x9f, 0x20, 0x7f, 0x7e, 0x33, 0x9a, 0xf9, 0x13,
		0xfb, 0x13, 0xd8, 0x03, 0xcf, 0x1f, 0x60, 0xf3, 0xf7, 0xc1, 0xf5, 0x68, 0x66, 0x5b, 0x07, 0xa0,
		0x30, 0xf2, 0xc7, 0xd3, 0x5b, 0xfb, 0xe8, 0x2d, 0xf9, 0x27, 0x43, 0xb4, 0x29, 0xe9, 0xbf, 0x33,
		0x44, 0xb7, 0x8b, 0xab, 0xbd, 0x0c, 0xaf, 0xc0, 0xcb, 0x07, 0xd8, 0xe4, 0x6a, 0xec, 0x87, 0xfe,
		0xfb, 0xb9, 0x6d, 0x1d, 0x00, 0x47, 0xe3, 0xc8, 0xbf, 0xf1, 0xa3, 0x5b, 0xfb, 0xe8, 0xf2, 0x37,
		0xf0, 0x32, 0xe1, 0xf9, 0x21, 0x45, 0x2f, 0x3b, 0xbb, 0xcd, 0xd5, 0xaa, 0x2c, 0xac, 0x5f, 0x87,
		0x4b, 0xa6, 0x56, 0x55, 0xec, 0x26, 0x3c, 0xf7, 0xf6, 0xff, 0xca, 0xef, 0x19, 0xc9, 0xbc, 0x25,
		0xaf, 0xbf, 0xaf, 0xe6, 0xe3, 0xfc, 0x09, 0x97, 0x6c, 0x3d, 0x8c, 0x9f, 0x1a, 0xdb, 0x0f, 0x7f,
		0x07, 0x00, 0x00, 0xff, 0xff, 0x41, 0xb6, 0x75, 0xa3, 0x5c, 0x05, 0x00, 0x00,
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
