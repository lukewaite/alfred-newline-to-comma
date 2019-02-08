# alfred-newline-to-comma

This workflow, when invoked by running `nlc` updates your clipboard, reformatting the
data contained within.

The purpose is to quickly take lists of values from Databases, Excel, etc, then use
them to quickly be able to write a `SELECT x FROM y WHERE z IN (<VALUES FROM CLIPBOARD>)`.

The workflow will take a newline-separated list of values, wrapping them in quotes if they
are strings, and remove the newlines, creating comma separated values which can be used
in a `IN` SQL clause.