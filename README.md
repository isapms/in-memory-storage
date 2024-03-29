# in-memory-storage

> Command line REPL to manage a simple in-memory key/value storage system.
> This system allows nested transactions. A transaction can then be committed or aborted.


## The shell accepts the following commands:

|   Command  |                                                             Description                                                                |
|:----------:|:--------------------------------------------------------------------------------------------------------------------------------------:|
|   `READ`   | Reads and prints to stdout, the value associated with key. If the value is not present an error is printed to stderr.                  |
|   `WRITE`  | Stores val in key.                                                                                                                     |
|   `DELETE` | Removes all key from store. Future READ commands on that key will return an error.                                                     |
|   `START`  | Start a transaction.                                                                                                                   |
|   `COMMIT` | All actions in the current tx are committed to the parent tx or the root store. If there is no current tx an error is output to stderr.|
|   `ABORT`  | Abort a transaction. All actions in the current transaction are discarded.                                                             |
|   `HELP`   | Show you the Help menu.                                                                                                                |
|   `CLEAR`  | Clear the terminal screen.                                                                                                             |
|   `QUITE`  | Exit the REPL cleanly. A message to stderr may be output.                                                                              |

## Example
<img src="https://i.imgur.com/0j6gimw.png">
