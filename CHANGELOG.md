# Changelog

## [Stage 1] — Beginning

- Print hardcoded loan repayment summary in correct order using `fmt.Println`

## [Stage 2] — Dreamworld

- Interactive stdin-based calculator for 0% interest loans
- Calculate number of months or monthly payment from user input
- Handle non-integer payments by rounding up; show adjusted last payment when it differs
- Singular/plural grammar for month count output

## [Stage 3] — Annuity payment

- Switch to CLI flag parsing via the `flag` package (`--principal`, `--periods`, `--payment`, `--interest`)
- Implement annuity payment formula; solve for whichever of principal, periods, or payment is missing
- Format period output as years and months with correct singular/plural (e.g. "8 years and 2 months")

## [Stage 4] — Differentiate payment

- Add required `--type` flag (`annuity` or `diff`)
- Implement differentiated payment formula with per-month output (`Month N: payment is X`)
- Calculate and print overpayment for both payment types
- Validate all inputs: missing/invalid type, missing interest, `diff`+`payment` combo, fewer than 4 flags, negative values — all print `Incorrect parameters`
