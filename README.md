# Avalon version 0.1

## A toolset for the modern auditor

#### by [Georgios Delkos](georgios.delkos@certik.io) for [CertiK](certik.io) Auditing Team 2020 

## Disclaimer
This is a WIP and could be buggy. For the moment it will work with any file for sigle file signing and with sol files in directory mode.

Files have to be for the moment in the respective folders /keys, /data, /signatures.

## Getting Started

### Install by cloning the repo and running make install

```bash
git clone https://github.com/GDHex/avalon.git
cd avalon
make install
```

### Create a new key pair with the gen-keys flag and a name for your keypair
```bash
./avalon gen-keys <name>
```

### Create a signature from a bundle of data

```bash
./avalon sign <private-key file> <file or directory> 
```

### Verify the signature 

```bash
./avalon verify <public-key file> <file or directory> <signarure>
```

### Print Locs for file or directory of files

```bash
./avalon loc <directory>
```

### Serve starts a service on the given port that can verify signatures against data and public key 

```bash
./avalon serve <port>
```

### Show returns the private and public keys for the username given in a human readable form

```bash
./avalon show <name>
```


### Info

``` bash
Welcome to Avalon, a tool to help auditors certify audits

Usage:
  avalon [command]

Available Commands:
  gen-keys    Gen-keys will return a ed25519 keypair
  help        Help about any command
  loc         Loc will return lines of code of the codebase in directory
  serve       Serve starts a service given a port number
  show        Show will load private and public key from files and show them in a hex format
  sign        Create a signature from a collection of data signed with a private key
  verify      Verify a signature against a public key and data

Flags:
      --config string   config file (default is $HOME/.avalon.yaml)
  -h, --help            help for avalon
  -t, --toggle          Help message for toggle

Use "avalon [command] --help" for more information about a command.
```

### TODO
* Create types for files types(Sol, Rs, Go, Js etc).
* Fix file importing from any location.
* Improve security.