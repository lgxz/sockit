# What

SockIt: forwarding between a normal TCP port and a SOCKS5 proxy port.

client --tcp-- [(local listening port) --SockIt] --socks5-- [(remote socks5 port) proxy] -- DEST

# Why
1. Not all programs are created equal. Some are not friendly to SOCKS5, such as rsync.
2. socat is too complicated to use.
3. It's interesting.

# How
## Command Line
1. `$ ./sockit -socks5 127.0.0.1:9050 127.0.0.1:8873 HIDDEN-RSYNC-SERVICE.onion:873`
2. `rsync rsync://127.0.0.1:8873`

## Docker
1. `docker run -d --rm -p 9050:9050 --name tao lgxz/tao`
2. `docker network create tornet`
3. `docker run -d --rm --name sockit-rsync --network tornet -p 5327:5327 lgxz/sockit HIDDEN-RSYNC-SERVICE.onion:873`
4. `rsync rsync://127.0.0.1:5327`



