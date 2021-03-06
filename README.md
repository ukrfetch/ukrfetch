# UKRfetch

:thumbsup:**Simple fetch tool to show Solidarity with Ukraine**

<img src="https://user-images.githubusercontent.com/100527338/155972141-66054d3c-eb5f-4c4b-9d92-5cb2ee90c3ad.png" width="540px" alt="fetch tool image">

This simple tool is inspired by [neofetch](https://github.com/dylanaraps/neofetch). 

This tool can use only for linux.

*Deep respect to everyone who fight for free at Ukraine.*

## Install(source build)

required: go >= 1.17

1. clone repositry
```bash
git clone https://github.com/ukrfetch/ukrfetch.git
```

2. build project
```bash
go build -o ./bin/ukrfetch fetch.go
```

3. if needed, move binary file and ADD $PATH
```bash
# example
mv ./bin/ /usr/local/ukrfetch/
PATH=$PATH:/usr/local/ukrfetch
```

## Install(download)

1. get release
```
wget https://github.com/ukrfetch/ukrfetch/releases/download/v1.0.0/ukrfetch-v1.0.0.tar.gz
```

2. extract and install file
```bash
tar -C /usr/local/ -xvf ukrfetch-v1.0.0.tar.gz
```

3. edit your .bashrc
```bash
# add the following at the end of your ~/.bashrc
PATH=$PATH:/usr/local/ukrfetch-v1.0.0
# add the following if you want to execute it every time you start terminal.
ukrfetch
```
```bash
# execute .bashrc
. ~/.bashrc
```

## Usage

```bash
# this is easy!
ukrfetch
```
