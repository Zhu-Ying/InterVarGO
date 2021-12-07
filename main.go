package main

import (
	"fmt"
	"github.com/Zhu-Ying/InterVarGO/dataset"
	"github.com/Zhu-Ying/InterVarGO/evidence"
	"github.com/Zhu-Ying/InterVarGO/utils"
	"github.com/Zhu-Ying/InterVarGO/variant"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

func initViper(dbDir, format string) {
	viper.AddConfigPath(dbDir)
	viper.SetConfigName(format)
	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		log.Fatal(err) // 读取配置文件失败致命错误
	}
}
func runInterver(snvs []variant.Snv, outfile string) {
	fo := utils.CreateFile(outfile)
	defer utils.CloseFile(fo)
	utils.WriteLine(fo, "Chr\tStart\tEnd\tRef\tAlt\tIntervar\tIntervar Summary\tIntervar Evidence\n")
	for _, snv := range snvs {
		intervar, summary, evidenceMap := evidence.Intervar(snv)
		utils.WriteLine(fo, fmt.Sprintf("%s\t%d\t%d\t%s\t%s\t%d\t%s\t%s\n",
			snv.Chrom, snv.Start, snv.End, snv.Ref, snv.Alt, intervar, summary, utils.ToJSON(evidenceMap)))
	}
}

var RootCmd *cobra.Command

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "intervar",
		Short: "AutoACMG by InterVar",
		Run: func(cmd *cobra.Command, args []string) {
			inputFile, _ := cmd.Flags().GetString("input")
			outputFile, _ := cmd.Flags().GetString("output")
			dbDir, _ := cmd.Flags().GetString("dbpath")
			genomeVer, _ := cmd.Flags().GetString("genome_version")
			format, _ := cmd.Flags().GetString("format")
			if inputFile == "" || outputFile == "" || dbDir == "" || genomeVer == "" || format == "" {
				err := cmd.Help()
				if err != nil {
					log.Panic(err)
				}
				os.Exit(-1)
			}
			if format != "annovar" {
				log.Panic("the format must be one of 'annovar'")
			}
			initViper(dbDir, format)
			variant.CheckTSV(inputFile)
			snvs := variant.ReadTSV(inputFile)
			dataset.InitIntervarDataset(dbDir, genomeVer)
			runInterver(snvs, outputFile)
		},
	}
	cmd.Flags().StringP("input", "i", "", "Input annotated file")
	cmd.Flags().StringP("output", "o", "", "Output file")
	cmd.Flags().StringP("dbpath", "d", "", "InterVarDB path")
	cmd.Flags().StringP("genome_version", "g", "hg19", "Genome version")
	cmd.Flags().StringP("format", "f", "annovar", "Input format")
	return cmd
}

func init() {
	RootCmd = NewRootCmd()
}

func main() {
	err := RootCmd.Execute()
	if err != nil {
		log.Panic(err)
	}
}
