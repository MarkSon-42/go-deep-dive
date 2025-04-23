# go-deep-dive
A deep dive into the core concepts and unique characteristics of the Go programming language. This repository explores Go's fundamentals like concurrency, interfaces, error handling, and idiomatic practices beyond the surface level. 


---

단순한 문법 나열을 넘어, Go가 **'Go답게' 작동하는 방식** (예: 동시성 모델, 인터페이스의 철학, 에러 처리 방식, 관용적 표현 등)의 원리와 그 배경에 담긴 **설계 사상**을 파악하는 것을 목표로 합니다

모든 설명은 마크다운(`.md`) 파일로 작성되며, 연습용 코드는 `.go` 또는 마크다운 내 코드 블록으로 작성합니다.

## 구조 (Table of Contents & Structure)

```text
.
├── README.md                     # 레포지토리 소개 및 전체 안내
├── LICENSE                       # 라이선스 정보
│
├── 01-concurrency/               # Go의 핵심: 동시성 (고루틴, 채널, select, sync)
│   ├── README.md                 #   - 동시성 개요
│   ├── goroutines.md             #   - 고루틴 상세
│   ├── channels.md               #   - 채널 상세
│   └── ...                       #   - 기타 관련 문서 및 코드
│
├── 02-interfaces/                # 인터페이스 (암시적 구현, 덕 타이핑, 표준 인터페이스 활용)
│   ├── README.md                 #   - 인터페이스 개요
│   └── ...                       #   - 관련 문서 및 코드
│
├── 03-error-handling/            # Go의 에러 처리 철학 (error 타입, 관용적 처리, defer/panic/recover)
│   ├── README.md                 #   - 에러 처리 개요
│   └── ...                       #   - 관련 문서 및 코드
│
├── 04-data-structures/           # 내장 자료구조 심층 이해 (슬라이스, 맵, 구조체와 임베딩)
│   ├── README.md                 #   - 자료구조 개요
│   └── ...                       #   - 관련 문서 및 코드
│
├── 05-functions-and-methods/     # 함수와 메서드 (값/포인터 수신자, 클로저, defer 심층 분석)
│   ├── README.md                 #   - 함수/메서드 개요
│   └── ...                       #   - 관련 문서 및 코드
│
├── 06-packages-and-visibility/   # 패키지와 가시성 (패키지 스코프, 가시성 규칙, 모듈)
│   ├── README.md                 #   - 패키지 시스템 개요
│   └── ...                       #   - 관련 문서 및 코드
│
├── 07-tooling-and-testing/       # Go 개발 도구 및 테스팅 (빌드/실행, 테스트, 벤치마크)
│   ├── README.md                 #   - 도구/테스팅 개요
│   └── ...                       #   - 관련 문서 및 코드
│
└── 08-advanced-topics/           # 고급 주제 (context, 리플렉션, 제네릭 등)
├── README.md                 #   - 고급 주제 개요
└── ...                       #   - 관련 문서 및 코드
```

---

04.24 defer, Goroutine

## 기회가 된다면 이후, github.io로 전환하여 기술 블로그 작성..
