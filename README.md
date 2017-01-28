# Mediaproxy

Mediaproxy is a web service for aggregating and serving private media content
such as:
* newsfeeds
* IRC conversations
* email
* web links
* images

## Architecture

Mediaproxy consists of following components:
- API (Manages subscriptions, serves collected contents)
- Web UI (UI for accessing content and managing subscriptions)
- Feed (Service for collecting RSS and Atom contents)
- Redis (Stores subscriptions and contents)

## Models

All data is stored to Redis as JSON

### Subscription

Defines settings for service to use.

Redis key: "sub:[service_name]:[subscription_id]" 
Example key: "sub:feed:98e0b063-f9ce-43fb-80f3-6e5587c4b5c7"

Service name should match to regex /[a-z_]+/
Subscription id should be UUID v4

Subscription content may be any JSON document that service can understand.

### Content

Redis key: "content:[service_name]:[subscription_id]:[timestamp]:[content_id]"
Example key: "content:feed:98e0b063-f9ce-43fb-80f3-6e5587c4b5c7:20170110082001050:1"

Timestamp should have millisecond accuracy.
Content id doesn't need to be unique, it's there only to differentiate
contents that have been registered in short timeperiod.

Content must be JSON document representing single content.
For example feed service might represent one item as follows:
```{
  "type": { "service": "feed", "version": 1.0 },
  "url": "https://feed/item/20140101145001",
  "timestamp": 201401101145001,
  "title": "Newsflash",
  "body": "Bla bla bla",
  "images": ["http://img.google.com/1.png"]
}```
