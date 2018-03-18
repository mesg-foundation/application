## mesg-cli workflow test

Test a workflow

### Synopsis

Test a workflow

```
mesg-cli workflow test FILE [flags]
```

### Examples

```
mesg-cli workflow test workflow.yml
```

### Options

```
  -e, --event string   Path to the event file
  -h, --help           help for test
  -k, --keep-alive     Keep the services alive (re-run without the option to shut it down)
  -l, --live           Use live events
  -t, --task string    Run the test on a specific task of the workflow
```

### SEE ALSO

* [mesg-cli workflow](mesg-cli_workflow.md)	 - Manage the workflow you create

