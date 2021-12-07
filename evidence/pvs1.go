package evidence

import (
	"github.com/Zhu-Ying/InterVarGO/dataset"
	"github.com/Zhu-Ying/InterVarGO/utils"
	"github.com/Zhu-Ying/InterVarGO/variant"
)

func CheckPVS1(snv variant.Snv) bool {
	pvs1 := false
	t1 := snv.IsNonsense() || snv.IsFrameshift() || snv.IsSplicing() // 条件1：无意变异或移码变异或剪接变异
	t2 := false
	for _, gene := range snv.RefGenes { // 基因中是否包含loss of fuction基因列表中的基因
		if _, ok := dataset.LOFGeneMap[gene]; ok {
			t2 = true
			break
		}
	}
	t3 := snv.IsSplicingDeletrious() // 剪接变异是否具有波坏性
	if t1 && t2 {
		pvs1 = true
		if snv.IsSplicing() && !t3 {
			// 如果是剪接变异，且没有破坏性则pvs1为false
			pvs1 = false
		} else {
			for _, aaChanges := range snv.KnownGeneAAChanges {
				transID, exon := aaChanges[1], aaChanges[2]
				if exonOrder, ok := dataset.KnownGeneCanonMap[transID]; ok && exon == "exon"+exonOrder {
					pvs1 = false
				} else {
					if pos, ok := dataset.KnownGeneCanonEDMap[transID]; ok && utils.StrToInt(pos)-snv.Start < 50 {
						pvs1 = false
					}
				}
			}
		}
	}
	return pvs1
}
