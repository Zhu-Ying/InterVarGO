package evidence

import (
	"github.com/Zhu-Ying/InterVarGO/variant"
)

// CheckBA1 判定BA1
func CheckBA1(snv variant.Snv) bool {
	// 人群频率>5%
	return snv.AF > float64(variant.Cutoff_AF2)
}
