# SubCommand 구현, 패키지 분리
mync 모듈 내의 cmd 패키지를 구현한다.<br>
cmd 패키지 내부엔 http, grpc 기능이 구현되어 있다.<br>
각 기능은 FlagSet으로 분리되어 있다.