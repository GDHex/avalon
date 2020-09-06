# Avalon version 0.1

## A toolset for the modern auditor

#### by [Georgios Delkos](georgios.delkos@certik.io) for [CertiK](certik.io) Auditing Team 2020 

## Disclaimer
This is a WIP and could be buggy. For the moment it will work with any file for sigle file signing and with sol files in directory mode.

## Instalation 

``` bash
go get github.com/GDHex/avalon
```

## Usage

#### Install by running make install

```bash
make install
```

#### Create a new key pair with the gen-keys flag and a name for your keypair
```bash
./avalon gen-keys <name>
```

#### Create a signature from a bundle of data

```bash
./avalon sign <private-key file> <file or directory> 
```

#### Verify the signature 

```bash
./avalon verify <public-key file> <file or directory> <signarure>
```

#### Print Locs for file or directory of files

```bash
./avalon loc <file or directory>
```

### TODO
* Create types for files types(Sol, Rs, Go, Js etc)
* Create a API