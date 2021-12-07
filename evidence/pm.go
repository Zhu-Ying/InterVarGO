package evidence

import (
	"fmt"
	"github.com/Zhu-Ying/InterVarGO/dataset"
	"github.com/Zhu-Ying/InterVarGO/variant"
)

// CheckPM1 判定PM1
func CheckPM1(snv variant.Snv) bool {
	// 位于突变热点、关键且完善的功能域（例如酶的活性位点），该区域一般无良性变异
	t1 := snv.IsMissense() // 是否错义变异
	t2 := false
	if snv.InterproDomain != "" {
		for _, gene := range snv.RefGenes { // 遍历所有Refgene中的基因
			key := fmt.Sprintf("%s_%s_%s", snv.Chrom, gene, snv.InterproDomain)
			if _, ok := dataset.DomainBenignMap[key]; !ok { // 该基因所在区域不属于良性区域
				t2 = true
				break
			}
		}
	}
	return t1 && t2
}

// CheckPM2 判定PM2
func CheckPM2(snv variant.Snv) bool {
	// 人群平均极低，取所属基因隐性遗传
	cutoffAF := 0.005
	if snv.AF == 0 {
		return true
	}
	if snv.AF < cutoffAF {
		for _, gene := range append(snv.RefGenes, snv.ENSGenes...) { // 遍历所有Refgene和ENSGene中的基因
			if mimnumber, ok := dataset.MIM2GeneMap[gene]; ok {
				if _, ok = dataset.MIMRecessMap[mimnumber]; ok {
					return true
				}
			}
		}
	}
	return false
}

// CheckPM3 判定PM3
func CheckPM3(snv variant.Snv) bool {
	return false
}

// CheckPM4 判定PM4
func CheckPM4(snv variant.Snv) bool {
	// 非重复区域内的非移码indel或stoploss
	if snv.IsStoploss() {
		return true
	}
	if snv.IsNonFrameshift() && snv.RMSK == "" {
		return true
	}
	return false
}

// CheckPM5 判定PM5
func CheckPM5(snv variant.Snv) bool {
	// 错义变异与已知致病性变异处于同一密码子但氨基酸改变不同
	t1 := snv.IsMissense() // 是否错义变异
	t2 := false
	t3 := false
	if len(snv.RefGeneAAChanges) > 0 {
		aa := snv.RefGeneAAChanges[0][4]
		aaLast := aa[len(aa)-1]
		for _, nt := range []string{"A", "C", "G", "T"} { // 遍历所有可能性
			if nt != snv.Ref && nt != snv.Alt { // 与当前变异不是同一个变异
				key := fmt.Sprintf("%s_%d_%d_%s", snv.Chrom, snv.Start, snv.End, nt)
				aa, ok := dataset.AAChangeMap[key]
				if ok {
					t3 = true
				}
				if aa == string(aaLast) { // 在AAChange数据库中，且氨基酸变化不同
					t2 = false
				}
			}
		}
	}
	return t1 && t2 && t3
}

// CheckPM6 判定PM6
func CheckPM6(snv variant.Snv) bool {
	return false
}
