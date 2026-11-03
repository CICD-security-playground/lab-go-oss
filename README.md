# lab-go-oss

`lab-go-oss` is the deliberately small Go service used as software under test in the CI/CD Security Playground. The application stays simple so later exercises can focus on pipeline trust, runner isolation, policy gates, dependency security, and artifact integrity rather than application complexity.

The service uses only the Go standard library.

## Endpoints

| Method | Path | Response |
|---|---|---|
| `GET` | `/health` | `{"status":"healthy"}` |
| `GET` | `/hello` | A JSON greeting |

The service listens on `:8080` by default. Set `ADDRESS` to override the listen address, for example `ADDRESS=:9090`.

## Build, run, and test

```sh
make build
make run
make test
```

After starting the service, try it with:

```sh
curl http://localhost:8080/health
curl http://localhost:8080/hello
```

Build the local container image with:

```sh
make docker-build
```

## Security scanner stubs

The scripts in `scripts/` are deterministic placeholders, not real security scanners. They pass by default:

```sh
make security-stubs
```

Each simulated check can be forced to fail independently:

```sh
FORCE_SAST_FAILURE=1 ./scripts/sast-stub.sh
FORCE_SCA_FAILURE=1 ./scripts/sca-stub.sh
FORCE_SECRETS_FAILURE=1 ./scripts/secrets-stub.sh
```

The scripts will later be replaceable by real SAST, dependency, and secret-scanning tools.

Future playground steps will add CI workflows, real scanners, AWS/ECR publishing, signing and provenance, and runner-security experiments. None of those integrations are configured in this repository yet.
