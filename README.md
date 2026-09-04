# Music API

Go, MySQL, Docker Compose로 만든 간단한 음악 조회 API입니다.

## 실행

`.env.example`을 복사해 `.env`를 만든 뒤 DB 정보를 설정하세요.

```bash
cp .env.example .env
```

`MYSQL_DSN`은 `MYSQL_USER`, `MYSQL_PASSWORD`, `MYSQL_DATABASE`에서 Docker Compose가 자동으로 구성합니다. 기존 `.env`에 `MYSQL_DSN`이 있다면 삭제하세요.

기존 MySQL volume을 사용하는 경우에는 최초 초기화 당시의 계정 비밀번호가 유지됩니다. 비밀번호를 변경했다면 root 계정으로 다음 명령을 실행하세요.

```bash
docker compose exec mysql mysql -uroot -p -e "ALTER USER 'musicapp'@'%' IDENTIFIED BY '새 비밀번호'; FLUSH PRIVILEGES;"
```

그 다음 실행합니다.

```bash
docker compose up --build
```

서버가 실행되면 `http://localhost:8080`에서 사용할 수 있습니다. 도메인 연결과 HTTPS 설정은 배포 단계에서 추가하면 됩니다.

Swagger UI에서 모든 API를 직접 호출해보려면 `http://localhost:8080/swagger`를 여세요.
OpenAPI 원본 문서는 `http://localhost:8080/openapi.json`에서 확인할 수 있습니다.

## API

- `GET /health` 데이터베이스 연결 상태
- `GET /api/artists` 아티스트 목록
- `GET /api/albums` 앨범 목록 (아티스트 이름 포함)
- `GET /api/songs` 음악 목록 (아티스트와 앨범 정보 포함)

예시:

```bash
curl http://localhost:8080/api/songs
```

MySQL 데이터는 `mysql_data` Docker volume에 저장됩니다. 초기 샘플 데이터를 다시 넣으려면 `docker compose down -v` 후 다시 실행하세요.
