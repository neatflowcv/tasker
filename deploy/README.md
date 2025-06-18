# Tasker 애플리케이션 배포 가이드

이 디렉토리는 Tasker 애플리케이션과 PostgreSQL 데이터베이스를 Kubernetes에 배포하기 위한 Kustomize 설정을 포함합니다.

## 사전 요구사항

- Kubernetes 클러스터 (v1.20+)
- kubectl CLI 도구
- kustomize (kubectl에 내장됨)

## 배포 방법

### 1. 모든 리소스 배포

```bash
# 현재 디렉토리에서 실행
kubectl apply -k .
```

### 2. 개별 리소스 확인

```bash
# 네임스페이스 확인
kubectl get namespace tasker

# PostgreSQL 확인
kubectl get pods,svc,pvc -n tasker -l app=postgres

# Tasker 애플리케이션 확인
kubectl get pods,svc -n tasker -l app=tasker
```

### 3. 서비스 접근

```bash
# LoadBalancer 타입으로 노출된 서비스 확인
kubectl get svc tasker-service -n tasker

# 포트포워딩으로 로컬 접근
kubectl port-forward svc/tasker-service 8080:80 -n tasker
```

그 후 브라우저에서 `http://localhost:8080/tasker/v1/swagger/index.html` 로 접근 가능합니다.

## 구성 요소

### PostgreSQL
- **Deployment**: postgres:15-alpine 이미지 사용
- **Service**: ClusterIP 타입, 5432 포트
- **PVC**: 2Gi 저장소
- **Secret**: 데이터베이스 인증 정보

### Tasker 애플리케이션
- **Deployment**: 3개 replica (패치로 2개로 조정 가능)
- **Service**: LoadBalancer 타입, 80 포트
- **ConfigMap**: 애플리케이션 설정
- **Init Container**: PostgreSQL 준비 대기

## 설정 변경

### 리소스 조정
`patches/` 디렉토리의 파일들을 편집하여 CPU/메모리 리소스를 조정할 수 있습니다.

### 이미지 태그 변경
`kustomize.yaml`의 `images` 섹션에서 이미지 태그를 변경할 수 있습니다.

### 데이터베이스 인증 정보 변경
`postgres-secret.yaml`의 base64 인코딩된 값을 변경합니다:

```bash
echo -n "new_password" | base64
```

## 정리

```bash
kubectl delete -k .
```

## 문제 해결

### PostgreSQL 연결 확인
```bash
kubectl exec -it deployment/postgres -n tasker -- psql -U postgres -d taskerdb
```

### 애플리케이션 로그 확인
```bash
kubectl logs deployment/tasker -n tasker
```

### 초기화 컨테이너 로그 확인
```bash
kubectl logs deployment/tasker -c wait-for-postgres -n tasker
``` 