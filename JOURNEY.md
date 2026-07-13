# Notiflex 여정 기록

이 파일은 독자가 실제로 진행한 내용을 기록한다. AI가 각 챕터 완료 시 자동으로 업데이트한다.

## 진행 현황

| 챕터 | 서브챕터 | 상태 | 완료일 | 비고 |
|------|---------|------|--------|------|
| ch2 | 2.2 설치 확인 | ✅ | 2026-07-05 | Claude Code, gcloud, kubectl 설치 완료 |
| ch2 | 2.3 gcloud 설정 | ✅ | 2026-07-05 | 프로젝트: ssw-gitaiops-project, 리전: asia-northeast3 |
| ch2 | 2.4 GitHub 저장소 | ✅ | 2026-07-05 | notiflex-platform 저장소 생성 |
| ch2 | 2.5 GKE 클러스터 | ✅ | 2026-07-05 | notiflex-cluster (e2-medium, 2 노드, Spot VM) |
| ch2 | 2.6 빌드/배포 | ✅ | 2026-07-06 | Notiflex API v0.1.0 배포, /health, /id 엔드포인트 동작 확인 |
| ch2 | 2.7 첫 커밋 | ✅ | 2026-07-06 | 초기 배포 완료 |
| ch3 | 3.2 GitOps 도구 | ✅ | 2026-07-07 | ArgoCD 선택 및 설치 |
| ch3 | 3.3 기능 추가 | ✅ | 2026-07-07 | 버전 정보 엔드포인트 추가, Rolling Update |
| ch3 | 3.4 CI | ✅ | 2026-07-07 | GitHub Actions 워크플로우 구성 |
| ch3 | 3.5 CI-CD 연결 | ✅ | 2026-07-07 | ArgoCD와 GitHub Actions 연동 |
| ch4 | 4.2 메트릭 모니터링 | ✅ | 2026-07-11 | Prometheus + Grafana 설치, 대시보드 확인 |
| ch4 | 4.3 로그 수집 | ✅ | 2026-07-12 | Loki + Fluent Bit 설치, {job="fluentbit"} 로그 확인 |
| ch4 | 4.4 알림 | ✅ | 2026-07-12 | PrometheusRule (PodRestart, PodCrashLooping, ApiPodDown) 설정 |
| ch5 | 5.2 트래픽 관리 | ✅ | 2026-07-13 | Gateway API 설정, 외부 IP 35.216.54.5로 접근 가능 |
| ch5 | 5.3 무중단 배포 | ⬜ | | |
| ch6 | 6.1 캐시 | ⬜ | | |
| ch6 | 6.2 시크릿 관리 | ⬜ | | |
| ch6 | 6.3 Canary 전환 | ⬜ | | |
| ch7 | 7.2 멀티 노드풀 | ⬜ | | |
| ch7 | 7.3 App of Apps | ⬜ | | |
| ch7 | 7.4 멀티테넌시 | ⬜ | | |
| ch8 | 8.1 메시징 | ⬜ | | |
| ch8 | 8.2 트레이싱 | ⬜ | | |
| ch8 | 8.3 CronJob | ⬜ | | |
| ch9 | 9.1 저장소 분석 | ⬜ | | |
| ch9 | 9.2 회고 | ⬜ | | |
| ch9 | 9.3 온보딩 문서 | ⬜ | | |
| ch9 | 9.4 GitAIOps 분석 | ⬜ | | |
| ch9 | 9.5 마무리 | ⬜ | | |

## 도구 선택 기록

독자가 3-프롬프트 패턴(탐색→비교→실행)에서 실제로 선택한 도구와 이유를 기록한다.

| 영역 | 선택 | 검토한 대안 | 선택 이유 |
|------|------|-----------|----------|
| GitOps 도구 | ArgoCD | Flux v2 | Kubernetes-native, 선언형 배포, UI 제공 |
| CI 도구 | GitHub Actions | Jenkins, GitLab CI | GitHub 통합, 간편한 설정, 빌드 자동화 |
| 메트릭 모니터링 | Prometheus + Grafana | Datadog, New Relic | 오픈소스, 자체 호스팅, Kubernetes 표준 |
| 로그 수집 | Loki + Fluent Bit | ELK, Splunk | 낮은 리소스 사용, 레이블 기반 인덱싱, Grafana 통합 |

## 현재 버전

| 컴포넌트 | 버전 | 변경 이력 |
|---------|------|----------|
| Go | 1.25 | 초기 설정 |
| Notiflex 이미지 | v0.1.1 | /version 엔드포인트 추가 (ch3에서 업그레이드) |
| ArgoCD | 2.13+ | ch3에서 설치 |
| Prometheus | (kube-prometheus-stack) | ch4에서 설치, kube-state-metrics 포함 |
| Grafana | (kube-prometheus-stack) | ch4에서 설치, Loki 데이터소스 추가 |
| Loki | (latest) | ch4에서 설치, SingleBinary 모드, replication_factor: 1 |
| Fluent Bit | (fluent/fluent-bit) | ch4에서 설치, Loki 데이터 송수신 |
| Kafka | (8장에서 설치) | |
| OTel SDK | (8장에서 설치) | |

## 현재 리소스

| 노드풀 | 머신 타입 | 노드 수 | 주요 워크로드 |
|--------|----------|---------|-------------|
| default-pool | e2-medium | 2 | Notiflex API |

## 트러블슈팅 이력

독자가 겪은 문제와 해결 방법을 기록한다. 같은 문제를 다시 겪지 않도록 한다.

| 챕터 | 문제 | 해결 |
|------|------|------|
| 2.5 | GKE 클러스터 생성 시 Kubernetes Engine API 미활성화 | gcloud services enable container.googleapis.com으로 API 활성화 |
| 2.6 | Cloud Build 서비스 계정 권한 부족 (storage, artifactregistry) | roles/storage.objectAdmin, roles/artifactregistry.writer 추가 |
| 4.3 | Loki Pod OOMKilled (메모리 부족) | singleBinary.resources.limits.memory: 512Mi로 증가 |
| 4.3 | Loki "too many unhealthy instances in the ring" 에러 | ingester.lifecycler.ring.replication_factor: 1 설정 추가 |
| 4.3 | Fluent Bit → Loki 연결 실패 | 비추천 grafana/fluent-bit 차트 → fluent/fluent-bit 공식 차트로 변경, Loki gateway 설정 수정 |
| 4.3 | Grafana에서 {namespace="notiflex"} 쿼리로 로그 조회 실패 | {job="fluentbit"} 레이블로 변경하여 해결 |
