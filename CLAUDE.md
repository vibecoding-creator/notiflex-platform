# Notiflex Platform

B2B 알림 SaaS 플랫폼의 GitAIOps 실습 저장소입니다.

## 프로젝트 개요

- **앱명**: Notiflex — B2B 알림 SaaS 플랫폼
- **언어**: Go (표준 라이브러리만 사용)
- **컨테이너**: scratch 베이스 이미지 (경량)

## GCP 환경

| 항목 | 값 |
|------|-----|
| **프로젝트 ID** | ssw-gitaiops-project |
| **리전** | asia-northeast3 (서울) |
| **존** | asia-northeast3-a |
| **Artifact Registry** | asia-northeast3-docker.pkg.dev/ssw-gitaiops-project/notiflex |

## 행동 규칙

1. **실행 전 확인**: 모든 kubectl, gcloud, docker 명령 실행 전에 현재 상태를 확인
2. **파괴적 작업**: `delete`, `rm -rf` 같은 명령은 반드시 대상을 명시하고 확인
3. **변경 기록**: 주요 변경사항은 git commit으로 기록
4. **컨텍스트 확인**: kubectl 명령 실행 시 `--context gke-sysnet4admin_book_gitaiops` 지정

## 디렉터리 구조

```
notiflex-platform/
├── CLAUDE.md           # 프로젝트 컨텍스트
├── app/                # Go 애플리케이션
├── k8s/
│   └── smb/            # Kubernetes 매니페스트
└── .github/
    └── workflows/      # CI/CD 파이프라인
```

## 참고 자료

- 실습 가이드: `../_Book_GitAIOps/`
- 가드레일: `../_Book_GitAIOps/prompt-guardrails/`
- 의사결정 가이드: `../_Book_GitAIOps/decision-guides/`
