package variant

import (
	"github.com/Zhu-Ying/InterVarGO/utils"
	"github.com/spf13/viper"
	"log"
	"math"
	"regexp"
	"strings"
)

func CheckTSV(datFile string) {
	fi, reader := utils.OpenFile(datFile)
	defer utils.CloseFile(fi)
	var headers []string
	for {
		line, isEof := utils.ReadLine(reader, 0)
		if isEof {
			break
		}
		headers = strings.Split(line, "\t")
		break
	}
	for _, key := range viper.AllKeys() {
		name := viper.GetString(key)
		pass := false
		for _, header := range headers {
			if strings.Contains(header, name) {
				pass = true
				break
			}
		}
		if !pass {
			log.Panicf("Error: can not find %v in headers", name)
		}
	}
}

func ReadTSV(datFile string) []Snv {
	snvs := make([]Snv, 0)
	fi, reader := utils.OpenFile(datFile)
	defer utils.CloseFile(fi)
	var headers []string
	for {
		line, isEof := utils.ReadLine(reader, 0)
		if isEof {
			break
		}
		fields := strings.Split(line, "\t")
		if len(headers) == 0 {
			headers = fields
			continue
		}
		snv := Snv{}
		for i, cell := range fields {
			header := headers[i]
			if cell == "." {
				cell = ""
			}
			switch header {
			case viper.GetString("chrom"):
				snv.Chrom = cell
			case viper.GetString("start"):
				snv.Start = utils.StrToInt(cell)
			case viper.GetString("end"):
				snv.End = utils.StrToInt(cell)
			case viper.GetString("ref"):
				snv.Ref = cell
			case viper.GetString("alt"):
				snv.Alt = cell
			case viper.GetString("refgene_gene"):
				if cell != "" {
					snv.RefGenes = regexp.MustCompile("[,;]").Split(cell, -1)
				}
			case viper.GetString("ensgene_gene"):
				if cell != "" {
					snv.ENSGenes = regexp.MustCompile("[,;]").Split(cell, -1)
				}
			case viper.GetString("refgene_aachange"):
				if cell != "" {
					for _, ac := range regexp.MustCompile("[,;]").Split(cell, -1) {
						snv.RefGeneAAChanges = append(snv.RefGeneAAChanges, strings.Split(ac, ":"))
					}
				}
			case viper.GetString("knowgene_aachange"):
				if cell != "" {
					for _, ac := range regexp.MustCompile("[,;]").Split(cell, -1) {
						snv.KnownGeneAAChanges = append(snv.KnownGeneAAChanges, strings.Split(ac, ":"))
					}
				}
			case viper.GetString("refgene_region"):
				snv.RefGeneRegion = cell
			case viper.GetString("refgene_event"):
				snv.RefGeneEvent = cell
			case viper.GetString("rmsk"):
				snv.RMSK = cell
			case viper.GetString("af"):
				if cell != "" {
					snv.AF = utils.StrToFloat64(cell)
				}
			case viper.GetString("metasvm"):
				if cell == "" {
					snv.MetaSVM = math.Inf(-1)
				} else {
					snv.MetaSVM = utils.StrToFloat64(cell)
				}
			case viper.GetString("gerp"):
				if cell == "" {
					snv.GERP = math.Inf(-1)
				} else {
					snv.GERP = utils.StrToFloat64(cell)
				}
			case viper.GetString("dbscsnv_ada"):
				if cell == "" {
					snv.DbscSnvADA = math.Inf(-1)
				} else {
					snv.DbscSnvADA = utils.StrToFloat64(cell)
				}
			case viper.GetString("dbscsnv_rf"):
				if cell == "" {
					snv.DbscSnvRF = math.Inf(-1)
				} else {
					snv.DbscSnvRF = utils.StrToFloat64(cell)
				}
			case viper.GetString("clinvar_clnsig"):
				snv.ClinvarCLNSIG = cell
			case viper.GetString("interpro_domain"):
				snv.InterproDomain = cell
			default:
				continue
			}
		}
		snvs = append(snvs, snv)
	}
	return snvs
}
