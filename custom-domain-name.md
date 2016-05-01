# Custom Domain Name

## Problem
Using serf-dns with domain name parsed into tags has some limitation:
- cannot use together with name and status
- cannot use regex
- limited to 127 labels, or 63 tags
- limited to 253 characters in full domain name

## Solution
Use custom pre-registered domain name with pre-configured serf filters

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
			"role": "web",
			"version": "201605[0-9][0-9].*"
		}
	},
	"us.dc.serf": {
		"tags": {
			"dc": "us-.*"
		}
	}
}
```

The JSON file above will register 3 domain names:        
- my-custom-dn-1.serf: return hosts named range from web-00 to web-59, still alive and tagged role=web
- failed.web.serf: return hosts whose name start with web-, already dead and tagged role=web and version start with 201605
- us.dc.serf: return hosts in US whose tag dc start with us-. Normally, query to this domain name will return hosts with tag dc=us
