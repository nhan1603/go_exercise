# Relationship

### Friend Management Project

Contains all the REST handlers that process the requests related to the relationship between users and return the corresponding responses.

The possible relationships between users that can be created are as below:
- Make friend
> POST /api/add-friend
- Subscribe
> POST /api/subscribe
- Block
> POST /api/block

The REST handler also provides a range of APIs that retrieve information about the relationship between two or more users:
- Retrieve common friend(s)
> POST /api/common-friend
- Retrieve list of users that can receive update from a provided email
> POST /api/update-receiver
