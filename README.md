<div  align="left">   

<img src="https://github.com/Wangxin555/BioHits/blob/dev/Logo.png" height="80" width="200" alt = "BioHits logo"/>

</div>

A Go toolkit for finding hot topics in biological and medical field.

## Description
Living in an era where findings in science and technology come out every single day, how to catch hot topics in the field of interest from tons of literature becomes an important problem. Also, when stepping into a new research area, we usually need to read some related papers, including classical ones and latest ones, to get a sense of what is it and what contributions can we make to this field. But carefully reading papers one by one requires much time and energy and sometimes we even don't need to know all the information in papers, maaybe some key words are sufficient for us to start.

To address this problem, here I present BioHits, a Go toolkit that can find hot topics for us automatically with just a few lines of code by scraping [PubMed](https://pubmed.ncbi.nlm.nih.gov/) iteratures, processing words in title and abstract, and doing some statistical summary.

It originally is a course project for 02601: Programming for Scientists, but I will continue making improvements and updates after finishing this course.


## Usage (for Professor Kingsford and TAs of 02-601)
### File organization
   
BioHits has two dependencies: colly and wordclouds, which had been zipped into the tar file

```
tar file
|
└───github.com
│   │
│   └───gocolly
|       └───colly
│   │
│   └───psykhi
|       └───wordclouds
│   
└───BioHits
|   |
|   └───fonts
|   └───stopwords
|   │   go.mod
|   │   go.sum
|   |   ...
|
└───test
|   │   go.mod
|   │   go.sum
|   │   main.go
|   │   test.exe
```
### How to use this package successfully
   
Because BioHits is built as a package, it is better if we test it by importing it and its functions in another program.
  * **Important note 1** : this package needs `go module`. In order to turn on `go module`, try run `set GO111MODULE=on` or `export GO111MODULE on` in your terminal.
  * **Important note 2** : because this package needs `go module`, if you run into an error says  `Can't find BioHits in xxx directory`, try only copy the **BioHits** folder to `GOROOT/src` (use `go env` command to find your `GOROOT` dir).
  * **Important note 3** : DO NOT enter in `BioHits` folder, enter in `test` folder instead, and run the program by running `./test.exe` (default parameters).

### Structure and parameters
  * There are 6 main functions implemented in BioHits, which are all used in test file `main.go`.
  * There are 12 parameters in total, whose type, description, default values have been listed in `main.go`. Parameters can be changed using `-flag value`, such as 
  `./test.exe -keyWords gene,cancer,kidney`.

### Output
   
   By default, there are three outputs, `searchOutput.txt`, `wordCloud.png` and  `topwords.txt`.
  * `searchOutput.txt`: it contains the PMID, title and abstract of papers.
  * `wordCloud.png`: this is the word cloud image generated using words extracted from title and abstract of papers.
  * `topwords.txt`: this txt file contains top words returned by BioHits.
   

## Author and Maintainer
Xin Wang, xinwang3@andrew.cmu.edu, Carnegie Mellon University

## Copyright
Copyright © 2021. Xin Wang. All Rights Reserved<br/>
