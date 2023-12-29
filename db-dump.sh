#!/bin/bash

mysqldump invoice \
    --single-transaction \
    --quick \
    --skip-lock-tables \
    --no-create-info | gzip > "invoice.sql.gz"