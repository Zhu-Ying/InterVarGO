package evidence

import (
	"github.com/Zhu-Ying/InterVarGO/dataset"
	"github.com/Zhu-Ying/InterVarGO/variant"
)

// CheckBP1 判定BS1
func CheckBP1(snv variant.Snv) bool {
	// 无义变异且在BP1基因列表中
	if snv.IsMissense() {
		for _, gene := range snv.RefGenes {
			if _, ok := dataset.BP1GeneMap[gene]; ok {
				return true
			}
		}
	}
	return false
}

// CheckBP2 判定BP2
func CheckBP2(snv variant.Snv) bool {
	return false
}

// CheckBP3 判定BP3
func CheckBP3(snv variant.Snv) bool {
	// 重复区域的非移码变异，且重复区域不在Interpro_domain中
	return snv.IsNonFrameshift() && snv.RMSK != "" && snv.InterproDomain == ""
}

// CheckBP4 判定BP4
func CheckBP4(snv variant.Snv) bool {
	// 多个生信预测软件预测为良性
	return snv.IsSplicingTolerated() && snv.IsConservTolerated() && (snv.IsSynony() || snv.IsEvolutTolerated())
}

// CheckBP5 判定BP5
func CheckBP5(snv variant.Snv) bool {
	return false
}

// CheckBP6 判定BP6
func CheckBP6(snv variant.Snv) bool {
	// CLNSIG 预测为良性
	return snv.IsClnsigBenign()
}

// CheckBP7 判定BP7
func CheckBP7(snv variant.Snv) bool {
	// 多种计算预测为良性
	return snv.IsSynony() && snv.IsSplicingTolerated() && snv.IsConservTolerated()
}
