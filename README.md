# s3upload

CLI to Upload files to amazon s3

Usage:

```bash
s3upload -file dbbackup.sql -key folder/example.txt -bucket bucketname
```

# Installation

```bash
go install github.com/abiiranathan/s3upload
```

AUTHENTICATION
Requires environment variables to be set:

```bash
AWS_ACCESS_KEY=
AWS_SECRET_KEY=
```
