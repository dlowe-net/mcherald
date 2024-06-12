# mcherald - Announce minecraft servers to the local network

mcherald is a command-line program to announce minecraft servers to
the local network.

Headless servers do not broadcast their presence on a local network,
since the intention is that they should be added to a client's server
list.  There are occasionally times when one wants to host a server
inside a home network, and mcherald can save a lot of configuration.

# Usage

Invocation is simple: `mcherald "My Minecraft Server:25565"`  More
than one server can be specified on the command line.

This can be incorporated in a shell script to automatically run when
the server runs, and to be terminated when the server exits.

```sh
 #!/bin/bash
 mcherald "My Minecraft Server:25565" &
 HERALD_PID=$!
 java -jar ./server.jar --nogui
 kill $HERALD_PID
```

# Help and Support

If you have any questions or feedback, please feel free to reach out
here on Github.  You can also contact me as `dlowe` on the
[libera.chat](https://libera.chat) IRC network.
