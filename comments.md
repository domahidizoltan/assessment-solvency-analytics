# Comments to my solution

Here I'd like to add some comments to my solution which might not be clear at the first sight.


## Algorithmic task

- Validating the input was not a requirement but I implemented validations as well. They are covered with tests. Probably only the `haystack_is_shorter` scenario might not be clear what is intended to validate when the haystack size is smaller than the needle size. Without this validation the output can't have the same size as the needle, what should always be the case.
- The functions and the error types are not exported because I usually try to narrow the access scope as much as possible. 


## Business task

I wanted to solve this task by using one of the JSON schema libs below:
- https://github.com/xeipuuv/gojsonschema
- https://github.com/qri-io/jsonschema

Unfortunately I don't have much experience with these libs and I also had the following concerns with them:
- The schema is dynamic for us (`$.schema`) what could be a problem for these libs. As I see they support reading a schema from string but this means I have to extract it from the input.
- The `required` field goes to the same level with the `properties` definition (and it's also an array) based on the [doc](https://json-schema.org/learn/getting-started-step-by-step#define-required-properties) what means I have to juggle with this because in our case they are separate fields within each field definition.
- My time is limited and I might run out of time while I try to force things work with our current input format.

Because of these reasons I implemented the validation by myself. I think it's reasonable at this level but for a production code I would rather invest the time to investigate the libs (or change the input format if possible) because in real life the validation rules could get more complicated with time so a proper tool would be better in the long run.

I try to implement the validation in a separate package by using any of the above mentioned libs but it might not be finished within the given time.

There are two more interesting parts:
- The `Validate` function and the errors are exported because this is a scenario of creating a lib where it makes no sense to not export the main access points and the error types (to let the callers check them later)
- "Optional business task 1" introduces the type check but the types could be optional in the first task to not break the existing test fixtures. "Optional business task 2" makes the type check mandatory what would break the previous test fixtures. Instead of changing the existing test fixtures I choosed to add the `forceTypeValidation` parameter to the validator. This makes the test drive the implementation details what is not a good practice but I think for a home assignment it could be acceptable in favor of keeping the test data intact.

P.S.: I managed to make the validation by using `github.com/xeipuuv/gojsonschema` what is under the `businesstask_lib` package. Most of the tests are copied from the `businesstask` package with some minor change (mostly the errors). The only thing what I could not cover is restricting the types in the schema definition (to use only `integer`, `string` and `boolean`).  
I also added benchmarks to both implementation and as I expected using the lib is slower. The benchmark on my machine resulted ~3.5x ns/op gain by using my custom implementation.  
This change was done in one step (in the last commit named "*implement JSON validation by using 3rd party lib*") because I already had most of the tests and I was also experimenting with things so I avoided the "baby steps" approach.


## Usage

- Running the tests with coverage:
```bash
go test -cover ./...
```
```bash
❯ go test -cover ./...
ok      solvencyanalytics/algorithmictask       (cached)        coverage: 100.0% of statements
ok      solvencyanalytics/businesstask  (cached)        coverage: 100.0% of statements
ok      solvencyanalytics/businesstask_lib      (cached)        coverage: 100.0% of statements
```

- Running benchmarks:
```bash
go test -bench=. -benchtime=1s ./...
```
```bash
❯ go test -bench=. -benchtime=1s ./...
PASS
ok      solvencyanalytics/algorithmictask       0.005s
goos: linux
goarch: amd64
pkg: solvencyanalytics/businesstask
cpu: Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz
BenchmarkValidate-4       139280              8634 ns/op
PASS
ok      solvencyanalytics/businesstask  2.243s
goos: linux
goarch: amd64
pkg: solvencyanalytics/businesstask_lib
cpu: Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz
BenchmarkValidate-4        36210             30815 ns/op
PASS
ok      solvencyanalytics/businesstask_lib      1.460s
```

- Running the tests and benchmark by using Docker (build and run):
```bash
docker build -t homework-domahidizoltan .
docker run --rm homework-domahidizoltan
```