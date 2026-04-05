# Bug Verification

## Status: Pending

## Verification Steps
1. Confirm `tables.filter` no longer throws when API returns null/non-array
2. Confirm `products.filter` is also protected
3. Confirm backend returns `[]` instead of `null` for empty table list
4. Server Station page loads without errors
5. Table selection and order flow still work correctly
