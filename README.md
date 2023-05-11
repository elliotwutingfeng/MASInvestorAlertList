# MAS Investor Alert List

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-2088FF?style=for-the-badge&logo=github-actions&logoColor=white)

[![GitHub license](https://img.shields.io/badge/LICENSE-BSD--3--CLAUSE-GREEN?style=for-the-badge)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/elliotwutingfeng/MASInvestorAlertList?style=for-the-badge)](https://goreportcard.com/report/github.com/elliotwutingfeng/MASInvestorAlertList)

Generate a machine-readable `.txt` blocklist of websites affiliated with [unregulated persons](https://www.mas.gov.sg/news/parliamentary-replies/2012/reply-to-parliamentary-question-on-investor-alert-list) who, based on information received by the [Monetary Authority of Singapore (MAS)](https://www.mas.gov.sg/investor-alert-list), may have been wrongly perceived as being licensed or regulated by MAS. This list is not exhaustive and is based on what was known to MAS at the time of publication.

**Disclaimer:** _This project is not sponsored, endorsed, or otherwise affiliated with the Monetary Authority of Singapore._

## Requirements

- Go >= 1.20

## Setup instructions

`git clone` and `cd` into the project directory, then run the following

```bash
go mod init dummy && go mod tidy
```

## Usage

```bash
go run scraper.go
```

## Libraries/Frameworks used

- [fasttld](https://github.com/elliotwutingfeng/go-fasttld)
- [xurls](https://github.com/mvdan/xurls)

## Monetary Authority of Singapore Terms of Use

<https://www.mas.gov.sg/terms-of-use>
