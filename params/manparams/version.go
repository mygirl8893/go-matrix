package manparams

import (
	"bytes"

	"github.com/MatrixAINetwork/go-matrix/common"
	"github.com/MatrixAINetwork/go-matrix/core/types"
)

const (
	VersionAlpha = "1.0.0.0"
	//增加版本号示例
	VersionBeta  = "1.0.0.1"
	VersionGamma = "1.0.0.2"
	//todo:上线后需要修改
	VersionSignatureGamma = "0x69bd3f6dbbca1012d7f68b5263900c9561da66b675088bc613460701c59b056e7b2695e1c3f84de28afd8f6797f1244bef1652a96d6cb58de151969cdc0956f700" //jerry
	VersionNumGamma        = uint64(330003)
	newP2PVersionTimeStamp = 1558346400
)

var VersionList [][]byte
var VersionSignatureMap map[string][]common.Signature

func init() {
	VersionList = [][]byte{[]byte(VersionAlpha), []byte(VersionBeta), []byte(VersionGamma)}
	VersionSignatureMap = make(map[string][]common.Signature)
	VersionSignatureMap[VersionGamma] = []common.Signature{common.BytesToSignature(common.FromHex(VersionSignatureGamma))}
}

// version1 > version2 return 1
// version1 = version2 return 0
// version1 < version2 return -1
func VersionCmp(version1 string, version2 string) int {
	if version1 == version2 {
		return 0
	}
	if version1 > version2 {
		return 1
	} else {
		return -1
	}
}

func IsCorrectVersion(version []byte) bool {
	if len(version) == 0 {
		return false
	}
	for _, item := range VersionList {
		if bytes.Equal(version, item) {
			return true
		}
	}
	return false
}
func GetVersionSignature(parentBlock *types.Block, version []byte) []common.Signature {
	if len(version) == 0 {
		return nil
	}
	if string(version) == string(parentBlock.Version()) {
		return parentBlock.VersionSignature()
	}
	if sig, ok := VersionSignatureMap[string(version)]; ok {
		return sig
	}

	return nil
}
func CanSwitchGammaCanonicalChain(currentTime int64) bool {
	return currentTime > newP2PVersionTimeStamp
}
