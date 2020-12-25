# Golang Serverless Suite

This package is a suite of helpers that I use regularly for my serverless applications, but their use is not just
limited to serverless applications, they can be used in any kind of go application.

Most of the helpers serve to facilitate the handling of external providers such as Sendgrid (for EMails) and 
AWS DynamoDB (NoSQL Key-Value DB) and to reduce the code coupling to these providers. For example, if one day 
I decide to replace DynamoDB with MongoDB, all I need to do is write a connector that implements the interface [BasetableProvider](/blob/master/itf/basetable.go) . 


## Structure

### `common`

`common` are co