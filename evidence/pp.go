package evidence

import (
	"fmt"
	"github.com/Zhu-Ying/InterVarGO/dataset"
	"github.com/Zhu-Ying/InterVarGO/variant"
)

// CheckPP1 判定PP1
func CheckPP1(snv variant.Snv) bool {
	return false
}

// CheckPP2 判定PP2
func CheckPP2(snv variant.Snv) bool {
	// 错义变异，且所属gene发生错义变异时为良性变异的概率较低
	if snv.IsMissense() { // 是否错义变异
		for _, gene := range snv.RefGenes {
			if _, ok := dataset.PP2GeneMap[gene]; ok {
				return true
			}
		}
	}
	return false
}

// CheckPP3 判定PP3
func CheckPP3(snv variant.Snv) bool {
	// 多个算法预测该变异具有破坏性
	val := 0
	hasSynony := snv.IsMissense() || snv.IsSynony()
	if snv.IsConservDeletrious() {
		val++
	}
	if snv.IsSplicingDeletrious() {
		val++
	}
	if snv.IsEvolutDeletrious() || !hasSynony {
		val++
	}
	if snv.Start == 97447487 {
		fmt.Println(snv.IsConservDeletrious(), snv.IsSplicingDeletrious(), snv.IsEvolutDeletrious(), hasSynony)
	}
	return val >= 2
}

// CheckPP4 判定PP4
func CheckPP4(snv variant.Snv) bool {
	return false
}

// CheckPP5 判定PP5
func CheckPP5(snv variant.Snv) bool {
	// 数据库报道为致病性变异
	return snv.IsClnsigPathogenic()
}
