# Web Crawler

## 프로그램 목적

청년 일자리 및 복지 정책 정보를 여러 사이트에서 자동으로 수집하여 DB에 저장하는 크롤러

## 주요 기능

- 지역/도시별 크롤링 (현재 수도권, 과천)
- 채용 정보 크롤링 (사람인)
- 복지 정책 크롤링 (청년몽땅정보통)
- 병렬 크롤링으로 빠른 데이터 수집
- MySQL DB 자동 저장
- HTML 덤프 자동 저장 (`dumps/날짜/사이트명_시간.html`)

## 구조

```
web_crawler/
├── main.go                  # 메인 실행 파일
├── crawler/                 # 크롤러 엔진
│   ├── crawler.go          # 공통 크롤링 로직
│   ├── config.go           # 크롤러 설정
│   └── parsers/            # 사이트별 파서
│       ├── saramin/        # 사람인 파서
│       └── youth_seoul/    # 청년몽땅정보통 파서
├── models/                  # 데이터 모델
│   ├── base.go             # 공통 모델
│   ├── job.go              # 채용 모델
│   └── welfare.go          # 복지 모델
├── utils/                   # 유틸리티
│   ├── file.go             # 파일 저장
│   ├── logger.go           # JSON 로깅
│   └── misc.go             # 기타 헬퍼
├── db/                      # DB 연결
├── repository/              # DB 쿼리
└── ddl/                     # 테이블 스키마
```

## 실행

```bash
# 의존성 설치
go mod download

# 실행
go run main.go
```

## 환경 요구사항

- Go 1.25+
- MySQL 8.0 (Docker)
- Colly v2
