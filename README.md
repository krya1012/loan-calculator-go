# Loan Calculator (Go)

A command-line loan calculator built in Go as part of the [Hyperskill](https://hyperskill.org/projects/430) course. It supports both annuity and differentiated payment types and can calculate missing loan parameters.

## Usage

### Current (Stage 3) — annuity calculator

Provide `--interest` and any two of the three remaining flags; the calculator solves for the missing one:

```bash
# Calculate monthly payment
go run "Loan Calculator (Go)/task/main.go" --principal=1000000 --periods=60 --interest=10

# Calculate number of months
go run "Loan Calculator (Go)/task/main.go" --principal=1000000 --payment=15000 --interest=10

# Calculate loan principal
go run "Loan Calculator (Go)/task/main.go" --payment=8721.8 --periods=120 --interest=5.6
```

### Stage 4 (full calculator)

```bash
# Differentiated payments
go run "Loan Calculator (Go)/task/main.go" --type=diff --principal=1000000 --periods=10 --interest=10

# Annuity payment amount
go run "Loan Calculator (Go)/task/main.go" --type=annuity --principal=1000000 --periods=60 --interest=10

# Loan principal
go run "Loan Calculator (Go)/task/main.go" --type=annuity --payment=8722 --periods=120 --interest=5.6

# Number of months to repay
go run "Loan Calculator (Go)/task/main.go" --type=annuity --principal=500000 --payment=23000 --interest=7.8
```

### Parameters

| Flag | Description |
|------|-------------|
| `--type` | Payment type: `annuity` or `diff` (Stage 4, required) |
| `--principal` | Loan principal amount |
| `--periods` | Number of monthly payments |
| `--payment` | Monthly payment amount |
| `--interest` | Annual interest rate in percent, e.g. `10` for 10% (required) |

Provide any 3 of the 4 non-type parameters; the calculator solves for the missing one and prints the overpayment.

## Formulas

**Annuity payment:**

```
A = P × (i × (1+i)^n) / ((1+i)^n − 1)
```

**Differentiated payment (month m):**

```
D_m = P/n + i × (P − P×(m−1)/n)
```

Where `i = annualRate / (12 × 100)`.

## Running Tests

Tests use the Python `hstest` framework:

```bash
pip install -r requirements.txt
cd "Loan Calculator (Go)/task"
python tests.py
```

## Project Stages

| Stage | Description |
|-------|-------------|
| 1 — Beginning | Print hardcoded repayment summary |
| 2 — Dreamworld | Interactive 0% interest calculator (stdin) |
| 3 — Annuity payment | CLI flags, annuity formulas, solve for missing parameter |
| 4 — Differentiate payment | Add `--type`, differentiated payments, overpayment, validation |
