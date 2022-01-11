# What is the Bekci?

Bekci is an application that periodically checks the versions and Kubernetes API Server SSL certificates of the clusters in the inventory list in Google Docs using the Google sheets API and informs them with slack.

Kubernetes master node ip information is kept as a list in a google sheet, and bekci sends a version request to the api server address of the ips in this list, finds the cluster version information, finds ssl certificate validity date by establishing a tls connection and automatically updates the relevant field in the google sheet. If the certificate expiry time is less than 1 month, it generates an alert to the slack channel.

# Using Goolge Sheets Api

Use this [link](https://developers.google.com/workspace/guides/create-credentials) to create google sheet credential token.