# Bank-statement handler case study

This case study explains the current working copy of:

```text
C:\go_files\neocore\api\apiMobileAppII\handlers-bank-statement.go
```

The source file was inspected read-only. The learning copy is
[`handlers-bank-statement.annotated.go.txt`](handlers-bank-statement.annotated.go.txt).
It contains every one of the 533 original lines plus teaching comments. It is
saved as `.go.txt` so `go test ./...` in this learning repository does not try
to compile neocore imports.

## What this file does

The file contains ten HTTP handler factories concerned with bank statements:

1. List financial institutions available for linking.
2. Choose an appropriate statement connector.
3. List a customer's linked external accounts.
4. Soft-delete a linked statement/account.
5. Translate internal link states into a small public status.
6. Initiate an MBS link request and schedule a refresh task.
7. Confirm the MBS link and possibly remove a loan requirement.
8. Return the public Mono key.
9. Accept a new Mono code.
10. Accept a manually uploaded statement.

Most handlers follow this backend flow:

```text
route calls factory
    -> factory captures services in a closure
    -> HTTP request enters returned handler
    -> parse and validate input
    -> identify customer
    -> call business service
    -> handle error or nil result
    -> write JSON response
    -> return
```

## The change you just made

Your first change replaces an error for a fintech without Mono support:

```go
// Before
apiHelpers.WriteErrorJSONResponse(
    w,
    400,
    "Sorry we do not currently support this institution",
)

// After
cr.ConnectionType = append(cr.ConnectionType, presenters.UploadOnly)
apiHelpers.WriteOKJSONResponse(w, 200, "success", cr)
```

The new behavior gives the client a manual-upload fallback rather than
rejecting that institution.

Your second change replaces an empty-slice check:

```go
// Before: add UploadOnly only when there are no connector choices.
if cr.ConnectionType == nil || len(cr.ConnectionType) == 0 {
    cr.ConnectionType = append(cr.ConnectionType, presenters.UploadOnly)
}

// After: add UploadOnly whenever neither Mono nor UploadOnly is present.
if !funk.ContainsString(cr.ConnectionType, presenters.Mono) &&
    !funk.ContainsString(cr.ConnectionType, presenters.UploadOnly) {
    cr.ConnectionType = append(cr.ConnectionType, presenters.UploadOnly)
}
```

That is a broader behavior change. For example, when the list contains only
`presenters.MbsAppinit`, the old code kept only MBS while the new code adds
`UploadOnly` too. This is correct if the product requirement is "always offer
manual upload when Mono is unavailable." If the requirement is merely "never
return an empty list," the old condition matched that requirement more closely.

## Topics you have learned that appear here

| Topic | Examples in the original file |
| --- | --- |
| Packages/imports | Lines 1-27 |
| Functions | Every top-level `func`, beginning at line 29 |
| Closures | Every factory returns an inline handler function |
| Variables and inference | `var` at 43-45; `:=` at 39, 41, 46, and many others |
| Strings and booleans | Query values at 48-56 and connector logic at 114-127 |
| Pointers and `nil` | Customer/account pointers and nil guards throughout |
| Structs | Response, link-request, and queue-task literals |
| Slices | `ConnectionType`, account lists, and `append` |
| Maps | JSON payloads at 274, 338, and 427 |
| Methods | `nc.Customer()`, `fi.WorksWith(...)`, `c.GetFullName()`, and others |
| Multiple returns | Calls returning `(value, error)` throughout |
| Error handling | The repeated `if err != nil` guard pattern |
| If/else | Input validation and business decisions throughout |
| Switch | Status conversion at lines 260-271 |
| Constants | HTTP methods and entity/presenter status identifiers |
| Blank identifier | `_` discards an unused link ID at line 457 |

## Topics that do not appear here

This file has no loop, recursion, generic, goroutine, built-in channel,
channel direction, `select`, or timeout. That is normal: production code should
not force a language feature into a handler that does not need it.

Despite the name of this import:

```go
"neocore/app/channels"
```

it is an application package for payment/channel services. It is unrelated to
Go values created with `make(chan T)`.

## Review observations

These are separate from your two changes, but worth noticing while learning:

1. Original line 207 passes `eaId` to `fmt.Sprintf` without a `%s` or `%v`
   placeholder. The resulting message can contain a `%!(EXTRA...)` formatting
   diagnostic. It likely intended:

   ```go
   fmt.Sprintf("External Account %q not found", eaId)
   ```

2. Original line 409 checks `errB` but wraps `err`. That likely logs the wrong
   error variable:

   ```go
   log.Println(errors.Wrap(errB, "RemoveBankStatementRequirement"))
   ```

3. Several `c == nil` branches wrap and pass `err` after `err` was confirmed
   nil. `errors.Wrap(nil, ...)` is nil, so those logs may print `<nil>` and the
   response helper receives no real error.

4. Original line 49 says `salect`; the client-facing message likely intended
   `select`.

5. Queue insertion failures at lines 334-336 and 393-395 are printed, but the
   handlers still return success. That may be intentional best-effort behavior,
   but it should be a deliberate product decision because a missing refresh
   task can leave processing incomplete.

No production source was changed by this case study.

## Suggested reading order

Start with these smaller functions before tackling the main connector logic:

1. `GetMonoPublicKey` — environment string, if guard, map, JSON response.
2. `ListBanksForLinking` — handler closure and nested calls.
3. `GetExternalAccountStatus` — pointers, errors, if/else, and switch.
4. `DeleteBankStatement` — authentication, authorization, and struct mutation.
5. `ChooseBankStatementConnector` — the full decision tree and your change.
