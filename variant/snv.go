package variant

import (
	"math"
	"strings"
)

type Cutoff float64

const (
	Cutoff_DBSCSNV Cutoff = 0.6
	Cutoff_METASVM Cutoff = 0
	Cutoff_GERP    Cutoff = 2
	Cutoff_AF1     Cutoff = 0.005
	Cutoff_AF2     Cutoff = 0.05
)

type Snv struct {
	Chrom              string
	Start              int
	End                int
	Ref                string
	Alt                string
	RefGenes           []string
	ENSGenes           []string
	RefGeneRegion      string
	RefGeneEvent       string
	RefGeneAAChanges   [][]string
	KnownGeneAAChanges [][]string
	AF                 float64
	InterproDomain     string
	RMSK               string
	DbscSnvADA         float64
	DbscSnvRF          float64
	MetaSVM            float64
	GERP               float64
	ClinvarCLNSIG      string
}

func (snv Snv) IsMissense() bool {
	// 错义变异
	return strings.Index(strings.ToLower(snv.RefGeneEvent), "nonsynony") != -1
}

func (snv Snv) IsSynony() bool {
	// 同义变异
	return strings.Index(strings.ToLower(snv.RefGeneEvent), "synony") != -1 && !snv.IsMissense()
}

func (snv Snv) IsNonFrameshift() bool {
	// 移码变异
	return strings.Index(strings.ToLower(snv.RefGeneEvent), "nonframeshift") != -1
}

func (snv Snv) IsFrameshift() bool {
	// 非移码变异
	return strings.Index(strings.ToLower(snv.RefGeneEvent), "frameshift") != -1 && !snv.IsNonFrameshift()
}

func (snv Snv) IsNonsense() bool {
	// 无义变异
	return strings.Index(strings.ToLower(snv.RefGeneEvent), "stopgain") != -1
}

func (snv Snv) IsStoploss() bool {
	// 终止丢失
	return strings.Index(strings.ToLower(snv.RefGeneEvent), "stoploss") != -1
}

func (snv Snv) IsSplicing() bool {
	// 剪接变异
	return strings.Index(strings.ToLower(snv.RefGeneRegion), "splic") != -1
}

func (snv Snv) IsSplicingDeletrious() bool {
	// 预测剪接破化性，主要参考dbscSNV数据库，score > 0.6
	return math.Max(snv.DbscSnvADA, snv.DbscSnvRF) > float64(Cutoff_DBSCSNV)
}

func (snv Snv) IsSplicingTolerated() bool {
	// 预测可容忍性，主要参考dbscSNV数据库，score <= 0.6
	return math.Max(snv.DbscSnvADA, snv.DbscSnvRF) <= float64(Cutoff_DBSCSNV)
}

func (snv Snv) IsEvolutDeletrious() bool {
	// 预测进化破化性，主要参考MetaSVM数据库，score > 0
	return snv.MetaSVM > float64(Cutoff_METASVM)
}

func (snv Snv) IsEvolutTolerated() bool {
	// 预测进化可容忍性，主要参考MetaSVM数据库，score <= 0
	return snv.MetaSVM != math.Inf(-1) && snv.MetaSVM <= float64(Cutoff_METASVM)
}

func (snv Snv) IsConservDeletrious() bool {
	// 预测保守性破化性，主要参考GERP++RS数据库，score > 2
	return snv.GERP > float64(Cutoff_GERP)
}

func (snv Snv) IsConservTolerated() bool {
	// 预测保守性可容忍性，主要参考GERP++RS数据库，score <= 2
	return snv.GERP <= float64(Cutoff_GERP)
}

func (snv Snv) IsClnsigPathogenic() bool {
	// CLINVAR收录致病性证据
	return strings.Index(strings.ToLower(snv.ClinvarCLNSIG), "pathogenic") != -1 && strings.Index(strings.ToLower(snv.ClinvarCLNSIG), "conflict") == -1
}

func (snv Snv) IsClnsigBenign() bool {
	// CLINVAR收录良性证据
	return strings.Index(strings.ToLower(snv.ClinvarCLNSIG), "benign") != -1
}
