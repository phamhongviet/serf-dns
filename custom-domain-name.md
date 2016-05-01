# Custom Domain Name

## Problem
Using serf-dns with domain name parsed as tags has some limitation:
- cannot use together with name and status
- cannot use regex
- limited to 127 labels, or 63 tags
- limited to 253 characters in full domain name

## Solution
Use custom pre-registered domain name with pre-configured serf filter

For example: configure domain name `my-custom-dn-1.serf` with name `^web-[0-5][0-9]`, and tags `role=web`

Use a JSON file to configure custom domain name, like this:

```json
{
	"my-custom-dn-1.serf": {
		"name": "^web-[0-5][0-9]",
		"status": "alive",
		"tags": {
			"role": "web"
		}
	},
	"failed.web.serf": {
		"name": "^web-.*",
		"status": "failed",
		"tags": {
			"role": "web"
		}
	},
	"us.dc.serf": {
		"tags": {
			"dc": "us-.*"
		}
	}
}
```
