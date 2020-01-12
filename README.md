# terraform-provider-remotefile

A terraform provider to allow for download of remote files over many protocols

## Installation

### Mac and Linux

Run the install script on the command line, the plugin will be installed to your
user terraform plugins directory.

```shell script
source <(curl -s https://raw.githubusercontent.com/simplecatsoftware/terraform-provider-remotefile/master/scripts/install.sh)`
```

## Usage

The initial use case is to allow for copying files from source to another while 
also checking that hashes are valid.

### Getting a remote file

```hcl-terraform
data "remotefile_read" "http_file" {
  source = "https://github.com/simplecatsoftware/lambda-http-example/archive/master.zip"
}

resource "aws_lambda_function" "http_lambda" {
  function_name = "test_http_lambda"
  handler = "program.Function"
  role = aws_iam_role.http_lambda.arn
  runtime = "ruby2.5"
  source_code_hash = data.remotefile.http_file.actual_sha256
  filename = data.remotefiles.http_file.local_path
}
```