# Bug Report: auth-cors-cloudshell

## Summary
Auth login fails with CORS error in Google Cloud Shell environment. The backend returns a 302 redirect which does not include `Access-Control-Allow-Origin` headers, causing the browser to block the request.

## Environment
- **Affected Environment**: Google Cloud Shell (port 8080 backend, port 3000 frontend)
- **Works In**: Docker local (same origin or properly configured)
- **Frontend Origin**: `https://3000-cs-4e479060-e78c-4cdd-82fc-bbcc8ecafa9b.cs-europe-west1-haha.cloudshell.dev`
- **Backend URL**: `https://8080-cs-4e479060-e78c-4cdd-82fc-bbcc8ecafa9b.cs-europe-west1-haha.cloudshell.dev/api/v1`

## Symptoms
```
Access to XMLHttpRequest at 'https://8080-.../api/v1/auth/login' from origin 'https://3000-...'
has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present on the requested resource.

POST https://8080-.../api/v1/auth/login net::ERR_FAILED 302 (Found)
```

## Root Cause (identified)
1. **302 redirect before CORS headers**: Gin's `RedirectTrailingSlash` (enabled by default) sends a 302 redirect that bypasses the CORS middleware — no `Access-Control-Allow-Origin` is added to the redirect response.
2. **No env-configurable CORS origins**: `VITE_API_URL` is `undefined` in Cloud Shell, and the backend has no `CORS_ORIGINS` env var support for dynamic test environments.

## Steps to Reproduce
1. Run backend on Cloud Shell port 8080
2. Run frontend on Cloud Shell port 3000
3. Attempt to log in
4. Observe 302 + CORS error in browser console

## Fix Plan
- Set `router.RedirectTrailingSlash = false` and `router.RedirectFixedPath = false` in `main.go`
- Add `CORS_ORIGINS` env variable support so Cloud Shell origins can be whitelisted at runtime
