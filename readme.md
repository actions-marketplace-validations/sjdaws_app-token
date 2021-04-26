# App token

This action copies [bubkoo/use-app-token](https://github.com/bubkoo/use-app-token) but also works when running behind a proxy.

## Usage

This action should be used when `GITHUB_TOKEN` is too restrictive.

```yaml
- name: Get token
  id: get_token
  uses: sjdaws/app-token@v1
  with:
    appId: ${{ secrets.APP_ID }}
    privateKey: ${{ secrets.APP_PEM }}
- name: Use token
  uses: ...
  with:
    token: ${{ steps.get_token.outputs.token }}
```
