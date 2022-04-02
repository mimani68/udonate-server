
# Find all users have GMAIL account

```javascript
db.getCollection('Users').find({ "connections.title": "email", "connections.value": /gmail/ig })
```