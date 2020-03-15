# ngb - Next-Generation Benchmarker

## Description

ngb is a benchmarker that can execute script at before and after requests.

## Use cases

### Benchmark an endpoint that requires an one-time-password header from other pages.

Target endpoint: https://example.com/target
OTP page: https://example.com/otp

write pre-request.sh like below:

```
#!/bin/bash

OTP=`curl https://example.com/otp`

JSON_FMT='{"headers":{"Authorization":"Bearer %s"}}\n'
printf "$JSON_FMT" "$OTP"
```

Then, execute ngb:

```
ngb -url "https://example.com/target" -prerequest ./pre-request.sh -c 10 -n 30
```

ngb executes pre-request.sh before requesting target.
After pre-request executing, ngb starts to request using output from pre-request.sh
The executing time in pre-request.sh excludes from target latency measure.

A pre-request script can define any of headers, cookies, and parameters.

Example:

```
JSON_FMT='{"headers":{"Authorization":"Basic val"},"cookies":[{"key":"_session_id", "value":"val"}],"params": {"user[name]":"takutakahashi"}}\n'
```

## Feature Works

I'm plan to imprement some features. See below:

- Result analysis with a specified script.
- Support all http method.
