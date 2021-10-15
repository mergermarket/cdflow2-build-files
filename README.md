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

Params
    path: (Required) file/folder to save in the release

