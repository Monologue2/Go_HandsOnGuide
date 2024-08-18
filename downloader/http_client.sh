#!/usr/bin/env bash
# Unix System Resource

# 3 is New File Descriptor
# File Descriptor is.. 네트워크 소켓과 같은 파일이나 I/O Resource에 액세스하기 위한 추상 표현
# 시스템으로부터 할당받은 파일 or 소켓을 대표하는 정수, 0 ~ 2번까지 고정되어 있다.
# File Open / Create Socket 시 3 부터 생성한다.

# <> Socket Open: Reading, Writing
# host/port
exec 3<>/dev/tcp/ysap.daveeddy.com/80

# HTTP: Hyper TexT Protocol. Is a Text! Protocol!
# HTTP는 평문을 통해 통신한다.
# Verb: GET, The Location the URI: /ping, Version of HTTP we use: HTTP/1.1
# Header 1: Host: ysap.daveeddy.com
# Header 2: Connection: close  -> 이후 Socker을 즉시 해제, 더 Request를 보내지 않을 것이므로..
lines=(
    'GET /ping HTTP/1.1'
    'Host: ysap.daveeddy.com'
    'Connection: close'
    ''
)

printf '%s\r\n' "${lines[@]}" >&3

while read -r data <&3; do
    echo "got server data: $data"
done

exec 3>&