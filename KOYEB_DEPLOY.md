# Deploying to Koyeb

This project contains separate `backend` and `frontend` services. Koyeb can build either from a Dockerfile in the repository.

- Backend Dockerfile: `backend/Dockerfile` (or use the new root `Dockerfile.koyeb`)
- Frontend Dockerfile: `frontend/Dockerfile`

Basic guidance to deploy on Koyeb:

1. Backend (API)
  - In the Koyeb Dashboard create a new Service -> choose `Dockerfile` build.
  - Set the Dockerfile path to `Dockerfile.koyeb` or `backend/Dockerfile`.
  - Set the port to `8080` (the app listens on `PORT` env var default `8080`).
  - Configure environment variables for DB connection: `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, and `DB_SSLMODE` as appropriate.
  - If your database is not publicly accessible, use a managed DB or an SSH/VPC solution recommended by Koyeb.

2. Frontend (static)
  - Create a separate Service in Koyeb.
  - Choose `Dockerfile` and set the path to `frontend/Dockerfile`.
  - Set the port to `3000` (nginx in the image listens on `3000`).

Notes & tips
- Multi-service: deploy backend and frontend as separate services in Koyeb for simplicity and scalability.
- Build context: Koyeb uses your repository as build context. The root `Dockerfile.koyeb` above copies `backend/` contents and builds it.
- Entrypoint/Command overrides: Koyeb allows overriding `ENTRYPOINT`/`CMD` in the service settings if you need to customize runtime command.
- Build stage selection: for multi-stage Dockerfiles you can choose the build stage Koyeb should produce/run (if needed).

Example Koyeb CLI (replace placeholders):

```bash
# Create backend service
koyeb service create --name alteqiachef-backend --dockerfile Dockerfile.koyeb --env PORT=8080

# Create frontend service
koyeb service create --name alteqiachef-frontend --dockerfile frontend/Dockerfile --env PORT=3000
```

If you'd like a single image that serves both frontend and backend from one container, tell me and I can create a combined multi-stage Dockerfile that copies `frontend/dist` into the backend and serves it (requires updating `backend` to serve static files or to embed them).
