
make cmdserver
servcli --help
servcli --port=80 --autotls=false

make cmdreader 
reader --help
reader --port=80 --schema=http

make cmdminer
miner --help
miner --port=80 --schema=http --path=/
