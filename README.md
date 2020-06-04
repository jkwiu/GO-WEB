# GO-WEB

1. Web Handler
   1. 정적 라우팅
      1. 인터페이스를 이용하는 방법
         1. type을 인터페이스로 이용
         2. 함수를 인터페이스로 이용
      2. 직접 선언
   2. mux를 이용한 동적 라우팅
2. JSON Transfer
   1. request, response를 JSON으로
   2. struct에 json annotation
   3. advanced rest clinent, post man같은 프로그램을 사용하여 body에 json type으로 담아서 request
      1. response로 json type을 받고 싶다면 header에 json type임을 명시
3. Test 환경(TDD)
   1. myapp폴더 만들어서 mux 분리
   2. myapp에 ``_test``를 붙여주면 test파일로 인식
   3. 함수명 앞에 ``Test``써주면 test code로 인식
   4. package 설치
      1. 알아서 테스트를 해주어 편하다.
      2. 테스트 서버(goconvey)
         1. ``github.com/smartystreets/goconvey``
         2. ``%GOPATH%/bin/goconvey`` 실행
         3. ``localhost:8080``에 서버시작
      3. 테스트 할 때 편하게 해줌(assert)
         1. ``github.com/stretchr/testify/assert``
         2. data buffer(buffer)
            1. 임시적으로 쓰이는 데이터가 물리적으로 저장되는 영역
      4. TCP
         1. Transmission Control Protocol
         2. 연결 지향 프로토콜(connection Oriented Protocol)
            1. 물리적으로 전용회선이 연결되어 있는 것처럼 가상의 연결통로를 설정하여 통신하는 방식으로 가상회선이라고도 한다.
         3. 신뢰할 수 있는 프로토콜(Reliable Protocol)
         4. TCP provides reliable, ordered, error-checked delivery of a stream of octets(bytes) between applications running on hosts communicating via an IP network.
4. File Upload