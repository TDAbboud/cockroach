// Code generated by "stringer -type=Method"; DO NOT EDIT

package roachpb

import "fmt"

const _Method_name = "GetPutConditionalPutIncrementDeleteDeleteRangeScanReverseScanBeginTransactionEndTransactionAdminSplitAdminMergeAdminTransferLeaseHeartbeatTxnGCPushTxnQueryTxnRangeLookupResolveIntentResolveIntentRangeNoopMergeTruncateLogRequestLeaseTransferLeaseLeaseInfoComputeChecksumDeprecatedVerifyChecksumCheckConsistencyInitPutChangeFrozenWriteBatchExportImport"

var _Method_index = [...]uint16{0, 3, 6, 20, 29, 35, 46, 50, 61, 77, 91, 101, 111, 129, 141, 143, 150, 158, 169, 182, 200, 204, 209, 220, 232, 245, 254, 269, 293, 309, 316, 328, 338, 344, 350}

func (i Method) String() string {
	if i < 0 || i >= Method(len(_Method_index)-1) {
		return fmt.Sprintf("Method(%d)", i)
	}
	return _Method_name[_Method_index[i]:_Method_index[i+1]]
}
