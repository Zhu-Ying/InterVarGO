package dataset

import (
	"path"
)

type Dataset map[string]string

var LOFGeneMap Dataset

var PP2GeneMap Dataset

var BP1GeneMap Dataset

var MIMRecessMap Dataset

var MIMDominMap Dataset

var MIMAdultMap Dataset

var MIMPhenoMap Dataset

var MIMOrphaMap Dataset

var OrphaMap Dataset

var AAChangeMap Dataset

var DomainBenignMap Dataset

var PS4SnpMap Dataset

var ExtSnpMap Dataset

var MIM2GeneMap Dataset

var BS2SnpRecessMap, BS2SnpDominMap Dataset

var KnownGeneCanonMap, KnownGeneCanonSTMap, KnownGeneCanonEDMap Dataset

func InitIntervarDataset(dbDir string, genomeVer string) {
	LOFGeneMap = readCommonFile(path.Join(dbDir, "PVS1.LOF.genes."+genomeVer), 0, -1, "\t")
	PP2GeneMap = readCommonFile(path.Join(dbDir, "PP2.genes."+genomeVer), 0, -1, "\t")
	BP1GeneMap = readCommonFile(path.Join(dbDir, "BP1.genes."+genomeVer), 0, -1, "\t")
	MIMRecessMap = readCommonFile(path.Join(dbDir, "mim_recessive.txt"), 0, -1, "\t")
	MIMDominMap = readCommonFile(path.Join(dbDir, "mim_domin.txt"), 0, -1, "\t")
	MIMAdultMap = readCommonFile(path.Join(dbDir, "mim_adultonset.txt"), 0, -1, "\t")
	MIMPhenoMap = readCommonFile(path.Join(dbDir, "mim_pheno.txt"), 0, 1, " ")
	MIMOrphaMap = readCommonFile(path.Join(dbDir, "mim_orpha.txt"), 0, 1, " ")
	OrphaMap = readCommonFile(path.Join(dbDir, "orpha.txt.utf8"), 0, 1, "\t")
	AAChangeMap = readSnpFile(path.Join(dbDir, "PS1.AA.change.patho."+genomeVer), []int{0, 1, 2, 4}, 6)
	DomainBenignMap = readSnpFile(path.Join(dbDir, "PM1_domains_with_benigns."+genomeVer), []int{0, 1, 2}, -1)
	PS4SnpMap = readSnpFile(path.Join(dbDir, "PS4.variants."+genomeVer), []int{0, 1, 1, 3, 4}, -1)
	ExtSnpMap = readSnpFile(path.Join(dbDir, "ext.variants."+genomeVer), []int{0, 1, 2, 3}, -1)
	MIM2GeneMap = readMIM2GeneFile(path.Join(dbDir, "mim2gene.txt"))
	BS2SnpRecessMap, BS2SnpDominMap = readBS2SnpFile(path.Join(dbDir, "BS2_hom_het."+genomeVer))
	KnownGeneCanonMap, KnownGeneCanonSTMap, KnownGeneCanonEDMap = readKnownGeneCanonFile(path.Join(dbDir, "knownGeneCanonical.txt."+genomeVer))
}
