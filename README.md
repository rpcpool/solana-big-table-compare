# solana-big-table-compare
Scripts &amp; libraries to compare two Solana BigTable Instances.

At the present time, historical Solana ledger data is contained in Google BigTable instances held by several parties. Since integrity of the data is critical, several archive holders would like the ability to compare their respective archives for equivalence.

## Solana Miami Hacker House BigTable Script Prize
We are offering a 5,000 USDC prize for the person or team who completes a working BigTable DB comparison script during the Hacker House event. This is a simple project that can be completed within the week for someone who is experienced with Google BigTable, Solana RPC, and writing data management scripts.

The final solution should include a script plus documentation that demonstrates the use of the script to compare two datasets.

The final solution will be open source (MIT license). It must be based on original work, but you are free to rely upon existing MIT-licensed libraries and code.

The final solution must be secure with no obvious vectors for abuse.

Eligibility for the Prize is determined by our judges: Brian Long & Linus Kendall.

## Features
You will prepare a script or program that can accept read-only credentials for two different BigTable instances and compare the data for a given range of slots.

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
We can provide read-only credentials for two different BigTable Instances to use during the Hackathon. Please contact `@brianlong` on Twitter or Telegram for details.

## Usage

### Requirements

- Git
- Go 1.17
- Basic Linux or macOS system

### Verifying available slot numbers

The `./cmd/export_slots` script will scan a `blocks` table and print out all slot numbers.
Outputted slot numbers are left padded to ensure lexicographic ordering.

```shell
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/bigtable-service-account.json
go run ./cmd/export_slots -project my-project-id > my-bigtable-slots.txt
```

The export will take up to a few hours for large Bigtable instances.

Suppose you have two lists of slots, `slots1.txt` and `slots2.txt`.
Using [`comm`](https://linux.die.net/man/1/comm), we can quickly check for differences.

- Get the first slot number: `head -n1 slots1.txt`
- Get the last slot number:  `tail -n1 slots1.txt`
- List slots only in the first table:  `comm -23 slots1.txt slots2.txt`
- List slots only in the second table: `comm -13 slots1.txt slots2.txt`
- List duplicate slots in both tables: `comm -12 slots1.txt slots2.txt`
