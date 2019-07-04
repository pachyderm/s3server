# s3server example

This example uses s3server to provide an S3-like API, with objects stored in-memory. It has a few shortcomings:

1) Data will be dropped when the process exits.
2) There are no optimizations, and quite a bit is implemented naively (e.g. the hash of files are recomputed every time it's necessary.)
3) Object listing has been purposefully simplified.

But it should be enough to demonstrate how to use s3server.

## Tests

To run tests, start the server with `make run`. Then, in a separate shell, run `make test`. After the tests complete, stats will be printed out regarding which tests failed. Complete logs are available in `test/runs`.