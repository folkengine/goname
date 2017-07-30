# Concourse CI Setup

1. Create a file called ~/.concourse/psgr-credentials.yaml
2. Add a field that contains your github private rsa key:

```
github-private-key: |
  -----BEGIN RSA PRIVATE KEY-----
  ...
  -----END RSA PRIVATE KEY-----
```