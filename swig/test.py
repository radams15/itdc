from PT import *

client = PT()

vers = client.version()


settings = client.get_settings()

print(settings.version)