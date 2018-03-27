## mesg-cli account import

Import an account from a backup file

### Synopsis

This method imports a previously exported backup file of your account created with the [export method](mesg-cli_account_export.md).

```
mesg-cli account import ./PATH_TO_BACKUP_FILE [flags]
```

### Examples

```
mesg-cli account import ./PATH_TO_BACKUP_FILE
mesg-cli account import ./PATH_TO_BACKUP_FILE --password PASSWORD --new-password PASSWORD
```

### Options

```
  -h, --help                  help for import
      --new-password string   New password of the imported account
      --password string       Current password of the account to import
```

### SEE ALSO

* [mesg-cli account](mesg-cli_account.md)	 - Manage your accounts

