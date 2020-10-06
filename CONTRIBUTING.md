Please test Go code using
```
go run
```
before committing as we do not have a GitHub action to test Go code yet.

Any and all changes must progress through lower branches: `deployment`, `frontend`, `backend` before being pushed to `main`, unless they are administrative commits.

Any deployment code relevant to deployment can be submitted with just
```
go build
```
as a test. Any extensive testing will be done on the administrative end.

Any PRs with minor edits will be rejected.

HTML commits will automatically tested when made on the `frontend` branch.

Contributions are currently limited to collaborators only.
