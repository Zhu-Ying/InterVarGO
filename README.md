# InterVarGO

#### 介绍

IntervarGO是使用GO语言重写[InterVar](https://github.com/WGLab/InterVar) 软件的项目。

- InterVar是一款依据ACMG-AMP 2015 guidelines进行SNV自动分类及证据定级的生信软件工具。
- InterVar内部自动调用[ANNOVAR](https://annovar.openbioinformatics.org/en/latest/) 进行注释分析，这对分析流程中以完成ANNOVAR注释项目造成了二次计算的资源浪费。
    - 本项目直接读取ANNOVAR注释结果，并通过配置config文件，允许用户配置自定的表头作为InterVar的必要输入。
- 本项目所有证据定级算法均基于InterVar原有算法重构，数据库也来源于InterVar项目

#### 安装教程

```shell
git clone https://github.com/Zhu-Ying/InterVarGO
cd InterVarGO
go build -o intervar main.go
```

#### 使用说明

1. 下载本项目中intervardb目录下所有文件
2. 根据你的annovar注释结果表头，修改intervardb/annovar.yml文件，**所有yaml中不允许修改key信息，只能修改value，否则可能造成错误或结果不准确**

```shell
intervar -i intervardb -o intervar.out 
openanno prepare transcript -d humandb/ -r ref.fa -g refGene.txt -o humandb/refgene
openanno prepare database -i gnomad.txt -o humandb/gnomad
```

## TODO

兼容其他注释软件结果