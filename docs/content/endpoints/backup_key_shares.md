+++
categories = ["Endpoints"]
date = "2020-09-11"
description = "Backup Key Shares (key splitting) endpoints"
tags = ["sso", "backup", "keyshares", "key", "splitting", "api", "endpoints"]
title = "SSO - Backup Key Shares"
+++

## 1. Introduction

Key Splitting consists in splitting a secret key in several (currently, always two) *key shares*.
One share alone is completely useless, but by combining two shares of a key one can recover the secret key.

A key share has another attribute than its value,
it has an `other_share_hash` which is used for the guest frontend to identify which share it wants to retrieve.
Technically speaking, the hash is the SHA-512 hash of the other share.

## 2. Creating a backup key share

### 2.1. request

```bash
POST https://api.misakey.com/backup-key-shares
```

_Cookies:_
- `accesstoken` (opaque token) (ACR >= 2): the identity must be linked to an account and this account must fit the one given in the body
- `tokentype`: must be `bearer`

_Headers:_
- `X-CSRF-Token`: a token to prevent from CSRF attacks. Delivered at the end of the auth flow.

_JSON Body:_
```json
{{% include "include/backup-key-share.json" %}}
```

- `account_id` (string) (uuid): the account for which the shares has been created.
- `share` (string) (base64): one of the shares.
- `other_share_hash` (string) (unpadded url-safe base64): a hash of the other share.
- `salt_base64` (string) (base64): the salt corresponding to the backup encryption.

### 2.2. response

_Code:_
```bash
HTTP 201 CREATED
```

_JSON Body:_
```json
{{% include "include/backup-key-share.json" %}}
```

## 3. Getting a Backup Key Share

### 3.1. request

```bash
GET https://api.misakey.com/backup-key-shares/:other-share-hash
```

_Cookies:_
- `accesstoken` (opaque token) (ACR >= 2): the identity must be linked to an account and this account must fit the one for which the key has been created.
- `tokentype`: must be `bearer`

_Headers:_
- `X-CSRF-Token`: a token to prevent from CSRF attacks. Delivered at the end of the auth flow.

_Path Parameters:_
- `other-share-hash` (string): the hash of the key share.


### 3.2. response

_Code:_
```bash
HTTP 200 OK
```

_JSON Body:_
```json
{{% include "include/backup-key-share.json" %}}
```

- `account_id` (string) (uuid): the account for which the shares has been created.
- `share` (string) (base64): one of the shares.
- `other-share-hash` (string) (unpadded url-safe base64): a hash of the other share.
- `salt_base64` (string) (base64): the salt corresponding to the backup encryption.
