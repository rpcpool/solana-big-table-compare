# solana-big-table-compare
Scripts &amp; libraries to compare two Solana BigTable Instances.

At the present time, historical Solana ledger data is contained in Google BigTable instances held by several parties. Since integrity of the data is critical, several archive holders would like the ability to compare their respective archives for equivalence.

## Features
Prepare a script or program that can accept read-only credentials for two different BigTable instances and compare the data for a given range of slots.

The script/program should do:
1. Verify that the two instances contain the same blocks.
2. Verify the same number of transactions within the blocks.
3. Allow us to quickly add more checks (e.g. signature comparisons) in the future.

The script/program should support:
- Configuration for two separate BigTable instances.
- A starting slot with a default of zero (0).
- An ending slot with a default of the current slot according to an RPC call.
- Specify rate limits to be sleep briefly within a loop. Default TBD.
- Specification of a log file location with the default in the current directory.
- Specification of different log error levels [INFO, WARN, ERROR] with default INFO.
- Use of environment variables or command line switches for the above settings.
- Error handling, lots of error handling.

We're expecting production-quality code with developer comments and informative error messages.

## Tooling
Your solution should be written in a popular software language and be easy to run with minimal installation requirements. Our team works in Go, Python, Ruby, Rust & other languages. We're open to ideas as long as they can meet the minimal installation requirement.

## BigTable Access
We can provide read-only credentials for two different BigTable Instances. Please contact `@brianlong` on Twitter or Telegram for details.
