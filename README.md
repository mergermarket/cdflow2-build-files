# cdflow2-build-files

A simple build plugin that saves a file or folder in the release archive for later use by Terraform.

## Example

```yaml
version: 2
builds:
  files:
    image: mergermarket/cdflow2-build-files
    params:
      path: file/or/folder/to/save
config:
  image: mergermarket/cdflow2-config-acuris
terraform:
  image: hashicorp/terraform
```

## Interface

### cdflow.yaml

Add a key to the `builds` map in cdflow.yaml (e.g. `files`). Along with the image
(i.e. `mergermarket/cdflow2-build-files), this takes a single required parameter
containing the file or folder to save (relative to the root folder where the
`cdflow.yaml` is stored):

Params
    path: (Required) file/folder to save in the release

### Terraform variable

With cdflow2 adding a build results in a Terraform variable with the same name -
i.e. if you call the build `files` then a `files` variable will be set. This includes
a single key `path` which contains the path to the saved file or folder (after the
release is unpacked). This is absolute so you can reference it directly in your
Terraform - for example to upload to an S3 bucket:

```terraform
variable "files" {
  type        = map(string)
  description = "Metadata from the 'files' build."
}

resource "aws_s3_bucket_object" "s3_upload" {
  for_each = fileset(var.files["path"], "**/*")

  bucket       = "my-bucket"
  key          = each.value
  source       = "${var.files["path"]}/${each.value}"
  etag         = filemd5("${var.files["path"]}/${each.value}")
  content_type = lookup(local.mime_types, regex("\\.[^.]+$", each.value), null)
}

locals {
  mime_types = {
    ".html" = "text/html"
  }
}
```

