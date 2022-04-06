from PT import PT

client = PT()

vers = client.version()

client.notify("Version", f"Version: {vers}")