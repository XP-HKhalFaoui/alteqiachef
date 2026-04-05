# Bug Analysis

## Status: Approved

## Root Cause
The `tables` variable in `ServerInterface` is used with `.filter()` and `.map()` at lines 162 and 165 without ensuring it's an array. The `useQuery` default value (`= []`) only applies when `data` is `undefined`. If the API returns `null` data (Go nil slice serialization) or a non-array shape, the `response.data || []` fallback in the queryFn should catch it — but doesn't handle cases where `response.data` is truthy but not an array.

The same vulnerability exists for `products` at line 139.

Meanwhile, `activeOrders` already has proper `Array.isArray()` guards (lines 147, 252).

## Fix Plan

### Approach: Add Array.isArray guards (consistent with existing activeOrders pattern)

1. **ServerInterface.tsx line 162**: Wrap `tables` with `Array.isArray()` check before `.filter()`
2. **ServerInterface.tsx line 165**: Wrap `tables` with `Array.isArray()` check before `.map()`
3. **ServerInterface.tsx line 139**: Wrap `products` with `Array.isArray()` check before `.filter()`
4. **Backend fix**: Initialize `tables` as empty slice (`tables := make([]models.DiningTable, 0)`) instead of nil to return `[]` instead of `null`

### Files to Modify
- `frontend/src/components/server/ServerInterface.tsx` (lines 139, 162, 165)
- `backend/internal/handlers/tables.go` (line 65)
