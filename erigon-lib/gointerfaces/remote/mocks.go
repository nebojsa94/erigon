// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package remote

import (
	context "context"
	types "github.com/nebojsa94/erigon/erigon-lib/gointerfaces/types"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	sync "sync"
)

// Ensure, that KVClientMock does implement KVClient.
// If this is not the case, regenerate this file with moq.
var _ KVClient = &KVClientMock{}

// KVClientMock is a mock implementation of KVClient.
//
//	func TestSomethingThatUsesKVClient(t *testing.T) {
//
//		// make and configure a mocked KVClient
//		mockedKVClient := &KVClientMock{
//			DomainGetFunc: func(ctx context.Context, in *DomainGetReq, opts ...grpc.CallOption) (*DomainGetReply, error) {
//				panic("mock out the DomainGet method")
//			},
//			DomainRangeFunc: func(ctx context.Context, in *DomainRangeReq, opts ...grpc.CallOption) (*Pairs, error) {
//				panic("mock out the DomainRange method")
//			},
//			HistoryGetFunc: func(ctx context.Context, in *HistoryGetReq, opts ...grpc.CallOption) (*HistoryGetReply, error) {
//				panic("mock out the HistoryGet method")
//			},
//			HistoryRangeFunc: func(ctx context.Context, in *HistoryRangeReq, opts ...grpc.CallOption) (*Pairs, error) {
//				panic("mock out the HistoryRange method")
//			},
//			IndexRangeFunc: func(ctx context.Context, in *IndexRangeReq, opts ...grpc.CallOption) (*IndexRangeReply, error) {
//				panic("mock out the IndexRange method")
//			},
//			RangeFunc: func(ctx context.Context, in *RangeReq, opts ...grpc.CallOption) (*Pairs, error) {
//				panic("mock out the Range method")
//			},
//			SnapshotsFunc: func(ctx context.Context, in *SnapshotsRequest, opts ...grpc.CallOption) (*SnapshotsReply, error) {
//				panic("mock out the Snapshots method")
//			},
//			StateChangesFunc: func(ctx context.Context, in *StateChangeRequest, opts ...grpc.CallOption) (KV_StateChangesClient, error) {
//				panic("mock out the StateChanges method")
//			},
//			TxFunc: func(ctx context.Context, opts ...grpc.CallOption) (KV_TxClient, error) {
//				panic("mock out the Tx method")
//			},
//			VersionFunc: func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*types.VersionReply, error) {
//				panic("mock out the Version method")
//			},
//		}
//
//		// use mockedKVClient in code that requires KVClient
//		// and then make assertions.
//
//	}
type KVClientMock struct {
	// DomainGetFunc mocks the DomainGet method.
	DomainGetFunc func(ctx context.Context, in *DomainGetReq, opts ...grpc.CallOption) (*DomainGetReply, error)

	// DomainRangeFunc mocks the DomainRange method.
	DomainRangeFunc func(ctx context.Context, in *DomainRangeReq, opts ...grpc.CallOption) (*Pairs, error)

	// HistoryGetFunc mocks the HistoryGet method.
	HistoryGetFunc func(ctx context.Context, in *HistoryGetReq, opts ...grpc.CallOption) (*HistoryGetReply, error)

	// HistoryRangeFunc mocks the HistoryRange method.
	HistoryRangeFunc func(ctx context.Context, in *HistoryRangeReq, opts ...grpc.CallOption) (*Pairs, error)

	// IndexRangeFunc mocks the IndexRange method.
	IndexRangeFunc func(ctx context.Context, in *IndexRangeReq, opts ...grpc.CallOption) (*IndexRangeReply, error)

	// RangeFunc mocks the Range method.
	RangeFunc func(ctx context.Context, in *RangeReq, opts ...grpc.CallOption) (*Pairs, error)

	// SnapshotsFunc mocks the Snapshots method.
	SnapshotsFunc func(ctx context.Context, in *SnapshotsRequest, opts ...grpc.CallOption) (*SnapshotsReply, error)

	// StateChangesFunc mocks the StateChanges method.
	StateChangesFunc func(ctx context.Context, in *StateChangeRequest, opts ...grpc.CallOption) (KV_StateChangesClient, error)

	// TxFunc mocks the Tx method.
	TxFunc func(ctx context.Context, opts ...grpc.CallOption) (KV_TxClient, error)

	// VersionFunc mocks the Version method.
	VersionFunc func(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*types.VersionReply, error)

	// calls tracks calls to the methods.
	calls struct {
		// DomainGet holds details about calls to the DomainGet method.
		DomainGet []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *DomainGetReq
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// DomainRange holds details about calls to the DomainRange method.
		DomainRange []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *DomainRangeReq
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// HistoryGet holds details about calls to the HistoryGet method.
		HistoryGet []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *HistoryGetReq
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// HistoryRange holds details about calls to the HistoryRange method.
		HistoryRange []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *HistoryRangeReq
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// IndexRange holds details about calls to the IndexRange method.
		IndexRange []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *IndexRangeReq
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// Range holds details about calls to the Range method.
		Range []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *RangeReq
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// Snapshots holds details about calls to the Snapshots method.
		Snapshots []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *SnapshotsRequest
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// StateChanges holds details about calls to the StateChanges method.
		StateChanges []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *StateChangeRequest
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// Tx holds details about calls to the Tx method.
		Tx []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
		// Version holds details about calls to the Version method.
		Version []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// In is the in argument value.
			In *emptypb.Empty
			// Opts is the opts argument value.
			Opts []grpc.CallOption
		}
	}
	lockDomainGet    sync.RWMutex
	lockDomainRange  sync.RWMutex
	lockHistoryGet   sync.RWMutex
	lockHistoryRange sync.RWMutex
	lockIndexRange   sync.RWMutex
	lockRange        sync.RWMutex
	lockSnapshots    sync.RWMutex
	lockStateChanges sync.RWMutex
	lockTx           sync.RWMutex
	lockVersion      sync.RWMutex
}

// DomainGet calls DomainGetFunc.
func (mock *KVClientMock) DomainGet(ctx context.Context, in *DomainGetReq, opts ...grpc.CallOption) (*DomainGetReply, error) {
	callInfo := struct {
		Ctx  context.Context
		In   *DomainGetReq
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockDomainGet.Lock()
	mock.calls.DomainGet = append(mock.calls.DomainGet, callInfo)
	mock.lockDomainGet.Unlock()
	if mock.DomainGetFunc == nil {
		var (
			domainGetReplyOut *DomainGetReply
			errOut            error
		)
		return domainGetReplyOut, errOut
	}
	return mock.DomainGetFunc(ctx, in, opts...)
}

// DomainGetCalls gets all the calls that were made to DomainGet.
// Check the length with:
//
//	len(mockedKVClient.DomainGetCalls())
func (mock *KVClientMock) DomainGetCalls() []struct {
	Ctx  context.Context
	In   *DomainGetReq
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *DomainGetReq
		Opts []grpc.CallOption
	}
	mock.lockDomainGet.RLock()
	calls = mock.calls.DomainGet
	mock.lockDomainGet.RUnlock()
	return calls
}

// DomainRange calls DomainRangeFunc.
func (mock *KVClientMock) DomainRange(ctx context.Context, in *DomainRangeReq, opts ...grpc.CallOption) (*Pairs, error) {
	callInfo := struct {
		Ctx  context.Context
		In   *DomainRangeReq
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockDomainRange.Lock()
	mock.calls.DomainRange = append(mock.calls.DomainRange, callInfo)
	mock.lockDomainRange.Unlock()
	if mock.DomainRangeFunc == nil {
		var (
			pairsOut *Pairs
			errOut   error
		)
		return pairsOut, errOut
	}
	return mock.DomainRangeFunc(ctx, in, opts...)
}

// DomainRangeCalls gets all the calls that were made to DomainRange.
// Check the length with:
//
//	len(mockedKVClient.DomainRangeCalls())
func (mock *KVClientMock) DomainRangeCalls() []struct {
	Ctx  context.Context
	In   *DomainRangeReq
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *DomainRangeReq
		Opts []grpc.CallOption
	}
	mock.lockDomainRange.RLock()
	calls = mock.calls.DomainRange
	mock.lockDomainRange.RUnlock()
	return calls
}

// HistoryGet calls HistoryGetFunc.
func (mock *KVClientMock) HistoryGet(ctx context.Context, in *HistoryGetReq, opts ...grpc.CallOption) (*HistoryGetReply, error) {
	callInfo := struct {
		Ctx  context.Context
		In   *HistoryGetReq
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockHistoryGet.Lock()
	mock.calls.HistoryGet = append(mock.calls.HistoryGet, callInfo)
	mock.lockHistoryGet.Unlock()
	if mock.HistoryGetFunc == nil {
		var (
			historyGetReplyOut *HistoryGetReply
			errOut             error
		)
		return historyGetReplyOut, errOut
	}
	return mock.HistoryGetFunc(ctx, in, opts...)
}

// HistoryGetCalls gets all the calls that were made to HistoryGet.
// Check the length with:
//
//	len(mockedKVClient.HistoryGetCalls())
func (mock *KVClientMock) HistoryGetCalls() []struct {
	Ctx  context.Context
	In   *HistoryGetReq
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *HistoryGetReq
		Opts []grpc.CallOption
	}
	mock.lockHistoryGet.RLock()
	calls = mock.calls.HistoryGet
	mock.lockHistoryGet.RUnlock()
	return calls
}

// HistoryRange calls HistoryRangeFunc.
func (mock *KVClientMock) HistoryRange(ctx context.Context, in *HistoryRangeReq, opts ...grpc.CallOption) (*Pairs, error) {
	callInfo := struct {
		Ctx  context.Context
		In   *HistoryRangeReq
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockHistoryRange.Lock()
	mock.calls.HistoryRange = append(mock.calls.HistoryRange, callInfo)
	mock.lockHistoryRange.Unlock()
	if mock.HistoryRangeFunc == nil {
		var (
			pairsOut *Pairs
			errOut   error
		)
		return pairsOut, errOut
	}
	return mock.HistoryRangeFunc(ctx, in, opts...)
}

// HistoryRangeCalls gets all the calls that were made to HistoryRange.
// Check the length with:
//
//	len(mockedKVClient.HistoryRangeCalls())
func (mock *KVClientMock) HistoryRangeCalls() []struct {
	Ctx  context.Context
	In   *HistoryRangeReq
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *HistoryRangeReq
		Opts []grpc.CallOption
	}
	mock.lockHistoryRange.RLock()
	calls = mock.calls.HistoryRange
	mock.lockHistoryRange.RUnlock()
	return calls
}

// IndexRange calls IndexRangeFunc.
func (mock *KVClientMock) IndexRange(ctx context.Context, in *IndexRangeReq, opts ...grpc.CallOption) (*IndexRangeReply, error) {
	callInfo := struct {
		Ctx  context.Context
		In   *IndexRangeReq
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockIndexRange.Lock()
	mock.calls.IndexRange = append(mock.calls.IndexRange, callInfo)
	mock.lockIndexRange.Unlock()
	if mock.IndexRangeFunc == nil {
		var (
			indexRangeReplyOut *IndexRangeReply
			errOut             error
		)
		return indexRangeReplyOut, errOut
	}
	return mock.IndexRangeFunc(ctx, in, opts...)
}

// IndexRangeCalls gets all the calls that were made to IndexRange.
// Check the length with:
//
//	len(mockedKVClient.IndexRangeCalls())
func (mock *KVClientMock) IndexRangeCalls() []struct {
	Ctx  context.Context
	In   *IndexRangeReq
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *IndexRangeReq
		Opts []grpc.CallOption
	}
	mock.lockIndexRange.RLock()
	calls = mock.calls.IndexRange
	mock.lockIndexRange.RUnlock()
	return calls
}

// Range calls RangeFunc.
func (mock *KVClientMock) Range(ctx context.Context, in *RangeReq, opts ...grpc.CallOption) (*Pairs, error) {
	callInfo := struct {
		Ctx  context.Context
		In   *RangeReq
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockRange.Lock()
	mock.calls.Range = append(mock.calls.Range, callInfo)
	mock.lockRange.Unlock()
	if mock.RangeFunc == nil {
		var (
			pairsOut *Pairs
			errOut   error
		)
		return pairsOut, errOut
	}
	return mock.RangeFunc(ctx, in, opts...)
}

// RangeCalls gets all the calls that were made to Range.
// Check the length with:
//
//	len(mockedKVClient.RangeCalls())
func (mock *KVClientMock) RangeCalls() []struct {
	Ctx  context.Context
	In   *RangeReq
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *RangeReq
		Opts []grpc.CallOption
	}
	mock.lockRange.RLock()
	calls = mock.calls.Range
	mock.lockRange.RUnlock()
	return calls
}

// Snapshots calls SnapshotsFunc.
func (mock *KVClientMock) Snapshots(ctx context.Context, in *SnapshotsRequest, opts ...grpc.CallOption) (*SnapshotsReply, error) {
	callInfo := struct {
		Ctx  context.Context
		In   *SnapshotsRequest
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockSnapshots.Lock()
	mock.calls.Snapshots = append(mock.calls.Snapshots, callInfo)
	mock.lockSnapshots.Unlock()
	if mock.SnapshotsFunc == nil {
		var (
			snapshotsReplyOut *SnapshotsReply
			errOut            error
		)
		return snapshotsReplyOut, errOut
	}
	return mock.SnapshotsFunc(ctx, in, opts...)
}

// SnapshotsCalls gets all the calls that were made to Snapshots.
// Check the length with:
//
//	len(mockedKVClient.SnapshotsCalls())
func (mock *KVClientMock) SnapshotsCalls() []struct {
	Ctx  context.Context
	In   *SnapshotsRequest
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *SnapshotsRequest
		Opts []grpc.CallOption
	}
	mock.lockSnapshots.RLock()
	calls = mock.calls.Snapshots
	mock.lockSnapshots.RUnlock()
	return calls
}

// StateChanges calls StateChangesFunc.
func (mock *KVClientMock) StateChanges(ctx context.Context, in *StateChangeRequest, opts ...grpc.CallOption) (KV_StateChangesClient, error) {
	callInfo := struct {
		Ctx  context.Context
		In   *StateChangeRequest
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockStateChanges.Lock()
	mock.calls.StateChanges = append(mock.calls.StateChanges, callInfo)
	mock.lockStateChanges.Unlock()
	if mock.StateChangesFunc == nil {
		var (
			kV_StateChangesClientOut KV_StateChangesClient
			errOut                   error
		)
		return kV_StateChangesClientOut, errOut
	}
	return mock.StateChangesFunc(ctx, in, opts...)
}

// StateChangesCalls gets all the calls that were made to StateChanges.
// Check the length with:
//
//	len(mockedKVClient.StateChangesCalls())
func (mock *KVClientMock) StateChangesCalls() []struct {
	Ctx  context.Context
	In   *StateChangeRequest
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *StateChangeRequest
		Opts []grpc.CallOption
	}
	mock.lockStateChanges.RLock()
	calls = mock.calls.StateChanges
	mock.lockStateChanges.RUnlock()
	return calls
}

// Tx calls TxFunc.
func (mock *KVClientMock) Tx(ctx context.Context, opts ...grpc.CallOption) (KV_TxClient, error) {
	callInfo := struct {
		Ctx  context.Context
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		Opts: opts,
	}
	mock.lockTx.Lock()
	mock.calls.Tx = append(mock.calls.Tx, callInfo)
	mock.lockTx.Unlock()
	if mock.TxFunc == nil {
		var (
			kV_TxClientOut KV_TxClient
			errOut         error
		)
		return kV_TxClientOut, errOut
	}
	return mock.TxFunc(ctx, opts...)
}

// TxCalls gets all the calls that were made to Tx.
// Check the length with:
//
//	len(mockedKVClient.TxCalls())
func (mock *KVClientMock) TxCalls() []struct {
	Ctx  context.Context
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		Opts []grpc.CallOption
	}
	mock.lockTx.RLock()
	calls = mock.calls.Tx
	mock.lockTx.RUnlock()
	return calls
}

// Version calls VersionFunc.
func (mock *KVClientMock) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*types.VersionReply, error) {
	callInfo := struct {
		Ctx  context.Context
		In   *emptypb.Empty
		Opts []grpc.CallOption
	}{
		Ctx:  ctx,
		In:   in,
		Opts: opts,
	}
	mock.lockVersion.Lock()
	mock.calls.Version = append(mock.calls.Version, callInfo)
	mock.lockVersion.Unlock()
	if mock.VersionFunc == nil {
		var (
			versionReplyOut *types.VersionReply
			errOut          error
		)
		return versionReplyOut, errOut
	}
	return mock.VersionFunc(ctx, in, opts...)
}

// VersionCalls gets all the calls that were made to Version.
// Check the length with:
//
//	len(mockedKVClient.VersionCalls())
func (mock *KVClientMock) VersionCalls() []struct {
	Ctx  context.Context
	In   *emptypb.Empty
	Opts []grpc.CallOption
} {
	var calls []struct {
		Ctx  context.Context
		In   *emptypb.Empty
		Opts []grpc.CallOption
	}
	mock.lockVersion.RLock()
	calls = mock.calls.Version
	mock.lockVersion.RUnlock()
	return calls
}

// Ensure, that KV_StateChangesClientMock does implement KV_StateChangesClient.
// If this is not the case, regenerate this file with moq.
var _ KV_StateChangesClient = &KV_StateChangesClientMock{}

// KV_StateChangesClientMock is a mock implementation of KV_StateChangesClient.
//
//	func TestSomethingThatUsesKV_StateChangesClient(t *testing.T) {
//
//		// make and configure a mocked KV_StateChangesClient
//		mockedKV_StateChangesClient := &KV_StateChangesClientMock{
//			CloseSendFunc: func() error {
//				panic("mock out the CloseSend method")
//			},
//			ContextFunc: func() context.Context {
//				panic("mock out the Context method")
//			},
//			HeaderFunc: func() (metadata.MD, error) {
//				panic("mock out the Header method")
//			},
//			RecvFunc: func() (*StateChangeBatch, error) {
//				panic("mock out the Recv method")
//			},
//			RecvMsgFunc: func(m any) error {
//				panic("mock out the RecvMsg method")
//			},
//			SendMsgFunc: func(m any) error {
//				panic("mock out the SendMsg method")
//			},
//			TrailerFunc: func() metadata.MD {
//				panic("mock out the Trailer method")
//			},
//		}
//
//		// use mockedKV_StateChangesClient in code that requires KV_StateChangesClient
//		// and then make assertions.
//
//	}
type KV_StateChangesClientMock struct {
	// CloseSendFunc mocks the CloseSend method.
	CloseSendFunc func() error

	// ContextFunc mocks the Context method.
	ContextFunc func() context.Context

	// HeaderFunc mocks the Header method.
	HeaderFunc func() (metadata.MD, error)

	// RecvFunc mocks the Recv method.
	RecvFunc func() (*StateChangeBatch, error)

	// RecvMsgFunc mocks the RecvMsg method.
	RecvMsgFunc func(m any) error

	// SendMsgFunc mocks the SendMsg method.
	SendMsgFunc func(m any) error

	// TrailerFunc mocks the Trailer method.
	TrailerFunc func() metadata.MD

	// calls tracks calls to the methods.
	calls struct {
		// CloseSend holds details about calls to the CloseSend method.
		CloseSend []struct {
		}
		// Context holds details about calls to the Context method.
		Context []struct {
		}
		// Header holds details about calls to the Header method.
		Header []struct {
		}
		// Recv holds details about calls to the Recv method.
		Recv []struct {
		}
		// RecvMsg holds details about calls to the RecvMsg method.
		RecvMsg []struct {
			// M is the m argument value.
			M any
		}
		// SendMsg holds details about calls to the SendMsg method.
		SendMsg []struct {
			// M is the m argument value.
			M any
		}
		// Trailer holds details about calls to the Trailer method.
		Trailer []struct {
		}
	}
	lockCloseSend sync.RWMutex
	lockContext   sync.RWMutex
	lockHeader    sync.RWMutex
	lockRecv      sync.RWMutex
	lockRecvMsg   sync.RWMutex
	lockSendMsg   sync.RWMutex
	lockTrailer   sync.RWMutex
}

// CloseSend calls CloseSendFunc.
func (mock *KV_StateChangesClientMock) CloseSend() error {
	callInfo := struct {
	}{}
	mock.lockCloseSend.Lock()
	mock.calls.CloseSend = append(mock.calls.CloseSend, callInfo)
	mock.lockCloseSend.Unlock()
	if mock.CloseSendFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.CloseSendFunc()
}

// CloseSendCalls gets all the calls that were made to CloseSend.
// Check the length with:
//
//	len(mockedKV_StateChangesClient.CloseSendCalls())
func (mock *KV_StateChangesClientMock) CloseSendCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockCloseSend.RLock()
	calls = mock.calls.CloseSend
	mock.lockCloseSend.RUnlock()
	return calls
}

// Context calls ContextFunc.
func (mock *KV_StateChangesClientMock) Context() context.Context {
	callInfo := struct {
	}{}
	mock.lockContext.Lock()
	mock.calls.Context = append(mock.calls.Context, callInfo)
	mock.lockContext.Unlock()
	if mock.ContextFunc == nil {
		var (
			contextOut context.Context
		)
		return contextOut
	}
	return mock.ContextFunc()
}

// ContextCalls gets all the calls that were made to Context.
// Check the length with:
//
//	len(mockedKV_StateChangesClient.ContextCalls())
func (mock *KV_StateChangesClientMock) ContextCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockContext.RLock()
	calls = mock.calls.Context
	mock.lockContext.RUnlock()
	return calls
}

// Header calls HeaderFunc.
func (mock *KV_StateChangesClientMock) Header() (metadata.MD, error) {
	callInfo := struct {
	}{}
	mock.lockHeader.Lock()
	mock.calls.Header = append(mock.calls.Header, callInfo)
	mock.lockHeader.Unlock()
	if mock.HeaderFunc == nil {
		var (
			mDOut  metadata.MD
			errOut error
		)
		return mDOut, errOut
	}
	return mock.HeaderFunc()
}

// HeaderCalls gets all the calls that were made to Header.
// Check the length with:
//
//	len(mockedKV_StateChangesClient.HeaderCalls())
func (mock *KV_StateChangesClientMock) HeaderCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockHeader.RLock()
	calls = mock.calls.Header
	mock.lockHeader.RUnlock()
	return calls
}

// Recv calls RecvFunc.
func (mock *KV_StateChangesClientMock) Recv() (*StateChangeBatch, error) {
	callInfo := struct {
	}{}
	mock.lockRecv.Lock()
	mock.calls.Recv = append(mock.calls.Recv, callInfo)
	mock.lockRecv.Unlock()
	if mock.RecvFunc == nil {
		var (
			stateChangeBatchOut *StateChangeBatch
			errOut              error
		)
		return stateChangeBatchOut, errOut
	}
	return mock.RecvFunc()
}

// RecvCalls gets all the calls that were made to Recv.
// Check the length with:
//
//	len(mockedKV_StateChangesClient.RecvCalls())
func (mock *KV_StateChangesClientMock) RecvCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRecv.RLock()
	calls = mock.calls.Recv
	mock.lockRecv.RUnlock()
	return calls
}

// RecvMsg calls RecvMsgFunc.
func (mock *KV_StateChangesClientMock) RecvMsg(m any) error {
	callInfo := struct {
		M any
	}{
		M: m,
	}
	mock.lockRecvMsg.Lock()
	mock.calls.RecvMsg = append(mock.calls.RecvMsg, callInfo)
	mock.lockRecvMsg.Unlock()
	if mock.RecvMsgFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.RecvMsgFunc(m)
}

// RecvMsgCalls gets all the calls that were made to RecvMsg.
// Check the length with:
//
//	len(mockedKV_StateChangesClient.RecvMsgCalls())
func (mock *KV_StateChangesClientMock) RecvMsgCalls() []struct {
	M any
} {
	var calls []struct {
		M any
	}
	mock.lockRecvMsg.RLock()
	calls = mock.calls.RecvMsg
	mock.lockRecvMsg.RUnlock()
	return calls
}

// SendMsg calls SendMsgFunc.
func (mock *KV_StateChangesClientMock) SendMsg(m any) error {
	callInfo := struct {
		M any
	}{
		M: m,
	}
	mock.lockSendMsg.Lock()
	mock.calls.SendMsg = append(mock.calls.SendMsg, callInfo)
	mock.lockSendMsg.Unlock()
	if mock.SendMsgFunc == nil {
		var (
			errOut error
		)
		return errOut
	}
	return mock.SendMsgFunc(m)
}

// SendMsgCalls gets all the calls that were made to SendMsg.
// Check the length with:
//
//	len(mockedKV_StateChangesClient.SendMsgCalls())
func (mock *KV_StateChangesClientMock) SendMsgCalls() []struct {
	M any
} {
	var calls []struct {
		M any
	}
	mock.lockSendMsg.RLock()
	calls = mock.calls.SendMsg
	mock.lockSendMsg.RUnlock()
	return calls
}

// Trailer calls TrailerFunc.
func (mock *KV_StateChangesClientMock) Trailer() metadata.MD {
	callInfo := struct {
	}{}
	mock.lockTrailer.Lock()
	mock.calls.Trailer = append(mock.calls.Trailer, callInfo)
	mock.lockTrailer.Unlock()
	if mock.TrailerFunc == nil {
		var (
			mDOut metadata.MD
		)
		return mDOut
	}
	return mock.TrailerFunc()
}

// TrailerCalls gets all the calls that were made to Trailer.
// Check the length with:
//
//	len(mockedKV_StateChangesClient.TrailerCalls())
func (mock *KV_StateChangesClientMock) TrailerCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockTrailer.RLock()
	calls = mock.calls.Trailer
	mock.lockTrailer.RUnlock()
	return calls
}
