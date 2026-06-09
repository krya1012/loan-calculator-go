# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a [Hyperskill](https://hyperskill.org/projects/430) course project: a Loan Calculator written in Go, progressing through 4 staged tasks. The active implementation file is `Loan Calculator (Go)/task/main.go`.

## Running the Program

```bash
# Stage 1 (Beginning) — no args needed
go run "Loan Calculator (Go)/task/main.go"

# Stage 3+ — command-line flags
go run "Loan Calculator (Go)/task/main.go" --type=annuity --principal=1000000 --periods=60 --interest=10
go run "Loan Calculator (Go)/task/main.go" --type=diff --principal=1000000 --periods=10 --interest=10
```

## Running Tests

Tests use Python's `hstest` framework (install with `pip install -r requirements.txt`):

```bash
cd "Loan Calculator (Go)/task"
python tests.py
```

## Stage Progression

The project has 4 stages, each described in its `task.html`:

| Stage | Folder | Description |
|-------|--------|-------------|
| 1 | `Beginning` | Print hardcoded loan repayment strings in order |
| 2 | `Dreamworld` | Interactive stdin-based 0% interest calculator (calculate months or monthly payment) |
| 3 | `Annuity payment` | CLI flags (`--principal`, `--periods`, `--payment`, `--interest`); calculate whichever value is missing using annuity formulas |
| 4 | `Differentiate payment` | Add `--type=annuity\|diff`, differentiated payment formula, overpayment output, and input validation |

## Key Formulas

**Annuity payment:** `A = P * (i*(1+i)^n) / ((1+i)^n - 1)`

**Differentiated payment (month m):** `D_m = P/n + i*(P - P*(m-1)/n)`

**Number of payments:** `n = log_{1+i}(A / (A - i*P))` — use `math.Log(x)/math.Log(base)` for change-of-base

Where `i = annualRate / (12 * 100)` (nominal monthly rate as decimal).

## Validation Rules (Stage 4)

- `--type` must be `annuity` or `diff`; always required
- `--interest` must always be provided
- `--type=diff` is incompatible with `--payment`
- Fewer than 4 total parameters → `Incorrect parameters`
- Any negative value → `Incorrect parameters`

## Output Conventions

- Round all payments **up** (`math.Ceil`)
- Period output: omit "0 years", use singular "1 year"/"1 month" vs plural
- Overpayment = sum of all payments − principal
