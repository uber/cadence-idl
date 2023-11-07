// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

namespace java com.uber.cadence.user

struct User {
    10: optional string username
    20: optional bool admin
    30: optional list<string> groups
}

struct AddUserRequest {
    10: optional User user
    20: optional string password
}

struct AddUserResponse {
      10: optional  User user
}

struct DeleteUserRequest {
      10: optional string username
}

struct GetUserRequest {
    10: optional string username
}

struct GetUserResponse {
      10: optional User user
}

struct ListUsersResponse {
    10: optional list<User> users
}

struct UpdateUserRequest {
 10: optional string username
 20: optional User user
}

struct UpdateUserResponse {
  10: optional User user
}

struct ChangeUserPasswordRequest {
  10: optional string username
  20: optional string password
}
