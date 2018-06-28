# dnsme

`dnsme` is a CLI tool that enables **very basic** management of a very small subset of DNSME's APIs. If you are looking for a complete DNSME tool, *this is not the tool you are looking for*.

## Install

`go get github.com/pdube/dnsme`

## Authentication

```bash
export DNSME_API=MY_API_KEY
export DNSME_SECRET=MY_SECRET_KEY
```

## Usage

### List domains

```bash
dnsme domains
```

*N.B.* if you want a pretty json output use :

```bash
dnsme domains | python -m json.tool
```

### List records in domain

```bash
dnsme records -d 12345
```

### Create record in domain

```bash
dnsme records create -d 12345 HOSTNAME IP_ADDRESS
```

### Update record in domain

```bash
dnsme records update -d 12345 -r 67890 HOSTNAME IP_ADDRESS
```
