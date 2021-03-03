// Copyright (c) 2017-2021 Uber Technologies, Inc.
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

namespace java com.uber.cadence.shadower

include "shared.thrift"

const string ShadowerLocalDomainName = "cadence-shadower"
const string ShadowerTaskList = "cadence-shadower-tl"

const string ShadowWorkflowName = "cadence-shadow-workflow"

const string ScanWorkflowActivityName = "scanWorkflowActivity"
const string ReplayWorkflowActivityName = "replayWorkflowActivity"

const string ErrReasonDomainNotExists = "domain not exists"
const string ErrReasonInvalidQuery = "invalid visibility query"
const string ErrReasonWorkflowTypeNotRegistered = "workflow type not registered"

enum ShadowMode {
  Normal,
  Continuous,
}

struct ExitCondition {
  10: optional i32 expirationIntervalInSeconds
  20: optional i32 shadowCount
}

struct ShadowWorkflowParams {
  10: optional string domain 
  20: optional string taskList
  30: optional string workflowQuery
  40: optional binary nextPageToken
  50: optional double samplingRate
  60: optional ShadowMode shadowMode
  70: optional ExitCondition exitCondition
  80: optional i32 concurrency
  90: optional ShadowWorkflowResult lastRunResult
}

struct ShadowWorkflowResult {
  10: optional i32 succeeded
  20: optional i32 skipped
  30: optional i32 failed
}

struct ScanWorkflowActivityParams {
  10: optional string domain
  20: optional string workflowQuery
  30: optional binary nextPageToken
  40: optional i32 pageSize
  50: optional double samplingRate
}

struct ScanWorkflowActivityResult {
  10: optional list<shared.WorkflowExecution> executions
  20: optional binary nextPageToken
}

struct ReplayWorkflowActivityParams {
  10: optional string domain
  20: optional list<shared.WorkflowExecution> executions
}

struct ReplayWorkflowActivityResult {
  10: optional i32 succeeded
  20: optional i32 skipped
  30: optional i32 failed
}
