# csp-parser

A simple utility to make `Content-Security-Policy`s readable at the command line

### Usage

Send it the Content-Security-Policy from `stdin`.

```
$ curl --silent -I https://maxchadwick.xyz | grep -i content-sec | ./csp-parser 
content-security-policy-report-only:

  default-src
    'self'
    c.disquscdn.com
    disqus.com

  script-src
    'self'
    'report-sample'
    maxchadwickblog.disqus.com
    c.disquscdn.com
    www.google-analytics.com
    cdn.syndication.twimg.com
    platform.twitter.com

  style-src
    'self'
    'unsafe-inline'
    ton.twimg.com
    platform.twitter.com
    c.disquscdn.com

  img-src
    'self'
    data:
    ton.twimg.com
    abs.twimg.com
    www.google-analytics.com
    www.gstatic.com
    c.disquscdn.com
    referrer.disqus.com
    syndication.twitter.com
    github.githubassets.com
    pbs.twimg.com
    platform.twitter.com

  frame-src
    'self'
    disqus.com
    platform.twitter.com
    syndication.twitter.com
    www.youtube.com

  child-src
    disqus.com
    platform.twitter.com
    syndication.twitter.com
    www.youtube.com

  connect-src
    'self'
    www.google-analytics.com

  font-src
    'self'
    fonts.gstatic.com

  form-action
    'self'
    platform.twitter.com
    syndication.twitter.com

  media-src
    'self'
    ssl.gstatic.com

  prefetch-src
    'self'
    disqus.com

  object-src
    'none'

  report-uri
    https://5dbedb272e44aa0ed5fad53074b6570a.report-uri.com/r/d/csp/reportonly
```