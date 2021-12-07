package evidence

import (
	"fmt"
	"github.com/Zhu-Ying/InterVarGO/dataset"
	"github.com/Zhu-Ying/InterVarGO/variant"
)

// CheckBS1 判定BS1
func CheckBS1(snv variant.Snv) bool {
	// 人群频率在0.5%~5%
	return float64(variant.Cutoff_AF1) <= snv.AF && snv.AF <= float64(variant.Cutoff_AF2)
}

// CheckBS2 判定BS2
func CheckBS2(snv variant.Snv) bool {
	// 在健康成人个体中观察到隐性（纯合）、显性（杂合）或 X 连锁（半合）疾病
	key := fmt.Sprintf("%s_%d_%d_%s_%s", snv.Chrom, snv.Start, snv.End, snv.Ref, snv.Alt)
	for _, gene := range append(snv.RefGenes, snv.ENSGenes...) {
		if mimnumber, ok := dataset.MIM2GeneMap[gene]; ok {
			if _, ok = dataset.MIMAdultMap[mimnumber]; ok {
				return true
			}
			if _, ok = dataset.MIMRecessMap[mimnumber]; ok {
				if val, ok := dataset.BS2SnpRecessMap[key]; ok && val == "1" {
					return true
				}
			}
			if _, ok = dataset.MIMDominMap[mimnumber]; ok {
				if val, ok := dataset.BS2SnpDominMap[key]; ok && val == "1" {
					return true
				}
			}
		}
	}
	return false
}

// CheckBS3 判定BS3
func CheckBS3(snv variant.Snv) bool {
	return false
}

// CheckBS4 判定BS4
func CheckBS4(snv variant.Snv) bool {
	return false
}
