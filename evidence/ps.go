package evidence

import (
	"fmt"
	"github.com/Zhu-Ying/InterVarGO/dataset"
	"github.com/Zhu-Ying/InterVarGO/variant"
)

// CheckPS1 判定PS1
func CheckPS1(snv variant.Snv) bool {
	// 错义变异与已知致病性变异具有相同的氨基酸变化
	t1 := snv.IsMissense() // 是否错义变异
	t2 := false
	if len(snv.RefGeneAAChanges) > 0 {
		aa := snv.RefGeneAAChanges[0][4]
		aaLast := aa[len(aa)-1]
		for _, nt := range []string{"A", "C", "G", "T"} { // 遍历所有可能性
			if nt != snv.Ref && nt != snv.Alt { // 与当前变异不是同一个变异
				key := fmt.Sprintf("%s_%d_%d_%s", snv.Chrom, snv.Start, snv.End, nt)
				if aa, ok := dataset.AAChangeMap[key]; ok && aa == string(aaLast) { // 在AAChange数据库中，且氨基酸变化相同
					t2 = true
				}
			}
		}
	}
	if t1 && t2 {
		return !snv.IsSplicingDeletrious()
	}
	return false
}

// CheckPS2 判定PS2
func CheckPS2(snv variant.Snv) bool {
	return false
}

// CheckPS3 判定PS3
func CheckPS3(snv variant.Snv) bool {
	return false
}

// CheckPS4 判定PS4
func CheckPS4(snv variant.Snv) bool {
	//GWAS分析中，该变异在患者群体中明显比对照群体中高。the dataset is from gwasdb jjwanglab.org/gwasdb
	key := fmt.Sprintf("%s_%d_%d_%s_%s", snv.Chrom, snv.Start, snv.End, snv.Ref, snv.Alt)
	_, ok := dataset.PS4SnpMap[key]
	return ok
}
