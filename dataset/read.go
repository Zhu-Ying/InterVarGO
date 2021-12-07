package dataset

import (
	"github.com/Zhu-Ying/InterVarGO/utils"
	"log"
	"regexp"
	"strings"
)

func readCommonFile(datFile string, kcol int, vcol int, delimiter string) Dataset {
	log.Printf("read %s", datFile)
	dataset := make(Dataset)
	fi, reader := utils.OpenFile(datFile)
	defer utils.CloseFile(fi)
	for {
		line, isEof := utils.ReadLine(reader, 0)
		if isEof {
			break
		}
		cells := strings.Split(line, delimiter)
		dataset[cells[kcol]] = "true"
		if vcol >= 0 {
			dataset[cells[kcol]] = cells[vcol]
		}
	}
	return dataset
}

func readSnpFile(datFile string, columns []int, vcol int) Dataset {
	log.Printf("read %s", datFile)
	dataset := make(Dataset)
	fi, reader := utils.OpenFile(datFile)
	defer utils.CloseFile(fi)
	for {
		line, isEof := utils.ReadLine(reader, 0)
		if isEof {
			break
		}
		cells := strings.Split(line, "\t")
		keys := make([]string, 0)
		for _, i := range columns {
			keys = append(keys, cells[i])
		}
		reg := regexp.MustCompile("[Cc][Hh][Rr]")
		key := reg.ReplaceAllString(strings.Join(keys, "_"), "")
		dataset[key] = "true"
		if vcol >= 0 {
			dataset[key] = cells[vcol]
		}
	}
	return dataset
}

func readMIM2GeneFile(datFile string) Dataset {
	log.Printf("read %s", datFile)
	mim2geneMap := make(Dataset)
	fi, reader := utils.OpenFile(datFile)
	defer utils.CloseFile(fi)
	for {
		line, isEof := utils.ReadLine(reader, 0)
		if isEof {
			break
		}
		cells := strings.Split(line, "\t")
		if len(cells) > 3 {
			mim2geneMap[strings.ToUpper(cells[3])] = cells[0]
		}
		if len(cells) > 4 {
			mim2geneMap[strings.ToUpper(cells[4])] = cells[0]
		}
	}
	return mim2geneMap
}

func readBS2SnpFile(datFile string) (Dataset, Dataset) {
	log.Printf("read %s", datFile)
	bs2SnpRecessMap, bs2SnpDominMap := make(Dataset), make(Dataset)
	fi, reader := utils.OpenGzipFile(datFile)
	defer utils.CloseFile(reader)
	defer utils.CloseFile(fi)
	lines := utils.ReadGzip(reader)
	for _, line := range lines {
		field := strings.Split(line, " ")
		key := strings.Join([]string{field[0], field[1], field[1], field[2], field[3]}, "_")
		bs2SnpRecessMap[key] = field[4]
		bs2SnpDominMap[key] = field[5]
		key = strings.Join([]string{field[0], field[1], field[1], utils.FlipACGT(field[2]), utils.FlipACGT(field[3])}, "_")
		bs2SnpRecessMap[key] = field[4]
		bs2SnpDominMap[key] = field[5]
	}
	return bs2SnpRecessMap, bs2SnpDominMap
}

func readKnownGeneCanonFile(datFile string) (Dataset, Dataset, Dataset) {
	log.Printf("read %s", datFile)
	knownGeneCanonMap, knownGeneCanonSTMap, knownGeneCanonEDMap := make(Dataset), make(Dataset), make(Dataset)
	fi, reader := utils.OpenFile(datFile)
	defer utils.CloseFile(fi)
	for {
		line, isEof := utils.ReadLine(reader, 0)
		if isEof {
			break
		}
		cells := strings.Split(line, " ")
		knownGeneCanonMap[cells[0]] = cells[1]
		knownGeneCanonSTMap[cells[0]] = cells[2]
		knownGeneCanonEDMap[cells[0]] = cells[3]
	}
	return knownGeneCanonMap, knownGeneCanonSTMap, knownGeneCanonEDMap
}
