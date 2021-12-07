package evidence

import (
	"fmt"
	"github.com/Zhu-Ying/InterVarGO/variant"
)

func getEvidences(snv variant.Snv, funcs ...func(variant.Snv) bool) []int {
	evidences := make([]int, 0)
	for _, f := range funcs {
		if f(snv) {
			evidences = append(evidences, 1)
		} else {
			evidences = append(evidences, 0)
		}
	}
	return evidences
}

func sumEvidences(evidences []int) int {
	sum := 0
	for _, evidence := range evidences {
		sum += evidence
	}
	return sum
}

func calPAS(pvs []int, ps []int, pm []int, pp []int) int {
	pas := 0
	sumPVS, sumPS, sumPM, sumPP := sumEvidences(pvs), sumEvidences(ps), sumEvidences(pm), sumEvidences(pp)
	// 2: Likely Pathogenic
	if sumPS == 1 && (sumPM == 1 || sumPM == 2) {
		pas = 2
	}
	if sumPVS == 1 && sumPM == 1 {
		pas = 2
	}
	if sumPS == 1 && sumPP >= 2 {
		pas = 2
	}
	if sumPM >= 3 {
		pas = 2
	}
	if sumPM == 2 && sumPP >= 2 {
		pas = 2
	}
	if sumPM == 1 && sumPP >= 4 {
		pas = 2
	}
	// 1: Pathogenic
	if sumPVS == 1 {
		if sumPS >= 1 {
			pas = 1
		}
		if sumPM >= 2 {
			pas = 1
		}
		if sumPM == 1 && sumPP == 1 {
			pas = 1
		}
		if sumPP >= 2 {
			pas = 1
		}
	}
	if sumPS >= 2 {
		pas = 1
	}
	if sumPS == 1 {
		if sumPM >= 3 {
			pas = 1
		}
		if sumPM == 2 && sumPP >= 2 {
			pas = 1
		}
		if sumPM == 1 && sumPP >= 4 {
			pas = 1
		}
	}
	return pas
}

func calBES(ba []int, bs []int, bp []int) int {
	sumBA, sumBS, sumBP := sumEvidences(ba), sumEvidences(bs), sumEvidences(bp)
	bes := 0
	// 4: Likely benign
	if sumBS == 1 && sumBP == 1 {
		bes = 4
	}
	if sumBP >= 2 {
		bes = 4
	}
	// 5: Benign
	if sumBA == 1 || sumBS >= 2 {
		bes = 5
	}
	return bes
}

func getSummary(evidences ...[]int) string {
	names := []string{"PVS", "PS", "PM", "PP", "BP", "BS", "BA"}
	summary := ""
	for i, name := range names {
		for j, e := range evidences[i] {
			if e == 1 {
				if summary == "" {
					summary = fmt.Sprintf("%s%d", name, j+1)
				} else {
					summary += fmt.Sprintf("/%s%d", name, j+1)
				}
			}
		}

	}
	return summary
}

func Intervar(snv variant.Snv) (int, string, map[string][]int) {
	pvs := getEvidences(snv, CheckPVS1)
	ps := getEvidences(snv, CheckPS1, CheckPS2, CheckPS3, CheckPS4)
	pm := getEvidences(snv, CheckPM1, CheckPM2, CheckPM3, CheckPM4, CheckPM5, CheckPM6)
	pp := getEvidences(snv, CheckPP1, CheckPP2, CheckPP3, CheckPP4, CheckPP5)
	bp := getEvidences(snv, CheckBP1, CheckBP2, CheckBP3, CheckBP4, CheckBP5, CheckBP6, CheckBP7)
	bs := getEvidences(snv, CheckBS1, CheckBS2, CheckBS3, CheckBS4)
	ba := getEvidences(snv, CheckBA1)
	summary := getSummary(pvs, ps, pm, pp, bp, bs, ba)
	pas, bes := calPAS(pvs, ps, pm, pp), calBES(ba, bs, bp)
	acmg := 3

	if pas > 0 && bes == 0 {
		acmg = pas
	}
	if pas == 0 && bes > 0 {
		acmg = bes
	}
	return acmg, summary, map[string][]int{"pvs": pvs, "ps": ps, "pm": pm, "pp": pp, "bp": bp, "bs": bs, "ba": ba}
	//return 0, "", map[string][]int{"a": {1, 23}}
}
