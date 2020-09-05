#!/bin/sh
env
/usr/local/bin/sockit -socks5=${SOCKS5_PROXY}  :${PORT} $@

