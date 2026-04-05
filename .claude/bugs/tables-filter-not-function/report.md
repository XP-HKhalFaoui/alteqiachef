# Bug Report

## Bug Summary
`ServerInterface` component crashes with `tables.filter is not a function` at line 162 because the `tables` data from the API may not be an array.

## Bug Details

### Expected Behavior
The Server Station page should load and display available dining tables, allowing servers to select a table and take orders.

### Actual Behavior
The component crashes with `TypeError: tables.filter is not a function` at [ServerInterface.tsx:162](frontend/src/components/server/ServerInterface.tsx#L162), causing the entire page to fail and hit the error boundary.

### Steps to Reproduce
1. Log in as a manager/server user
2. Navigate to the Server Station (home page)
3. Component crashes immediately during render

### Environment
- **Version**: Current main branch
- **Platform**: React/Vite frontend, Go backend
- **Configuration**: Development (localhost:5173 + localhost:8080)

## Impact Assessment

### Severity
- [x] High - Major functionality broken

### Affected Users
All server/manager users trying to access the Server Station interface.

### Affected Features
- Server Station (dine-in order taking)
- Table selection and floor view
- Order creation flow

## Additional Context

### Error Messages
```
ServerInterface.tsx:162  Uncaught TypeError: tables.filter is not a function
    at ServerInterface (ServerInterface.tsx:162:34)
    at renderWithHooks (react-dom.development.js:15486:18)
    at mountIndeterminateComponent (react-dom.development.js:20103:13)
```

### Screenshots/Media
Full console output shows authentication succeeds, user data loads correctly, but the component crashes when attempting to filter tables data.

### Related Issues
The same pattern exists for `activeOrders` which already has an `Array.isArray()` guard at line 147, but `tables` and `products` lack this protection.

## Initial Analysis

### Suspected Root Cause
The Go backend's `GetTables` handler declares `var tables []models.DiningTable` which stays `nil` if no rows are returned. When serialized to JSON, a nil Go slice inside an `interface{}` field becomes `null` in JSON (not `[]`). While the `queryFn` has `response.data || []` as a fallback, there are edge cases where `useQuery`'s `data` destructuring default (`= []`) only applies when `data` is `undefined`, not when the query returns a non-array truthy value. Additionally, the `response.data` could be a non-array object if the API response shape changes or if an error response is structured differently.

### Affected Components
- [frontend/src/components/server/ServerInterface.tsx](frontend/src/components/server/ServerInterface.tsx) - Lines 87-98 (query), 139, 162, 165
- [backend/internal/handlers/tables.go](backend/internal/handlers/tables.go) - Lines 65, 114-118 (nil slice serialization)
