To build and launch the cluster just type 
```
make build && make up
```

then to get access to each server console do the next: 

```
make cmdserver
```

and type 

```
servcli --help
```

to get help

Example command to launch server:

```
servcli --port=80 --autotls=false
```

... reader:

```
make cmdreader 
```

```
reader --help
```

```
reader --port=80 --schema=http
```

... and file miner:

```
make cmdminer
```

```
miner --help
```

```
miner --port=80 --schema=http --path=/
```