## Tasks and Reminder Manager

## Routes

### Tasks Routes

```
"/tasks" - POST - Creates a new task.
"/tasks" - GET - Get all the tasks.
"/tasks/{taskID}" - GET - Get a task by task ID
"/tasks/{taskID}" - PUT - Update a task by task ID
"/tasks/{taskID}" - DELETE - Delete a task by task ID 
```

### Reminder Routes

```
"/reminders" - POST - Creates a new reminder.
"/reminders" - GET - Get all the reminders.
"/reminders/{reminderID}" - GET - Get a reminder by reminder ID
"/reminders/{reminderID}" - PUT - Update a reminder by reminder ID
"/reminders/{reminderID}" - DELETE - Delete a reminder by reminder ID 
```

### Additional Features

```
Logging Enabled to trace the requests.
Adding authentication middleware as well for all the routes.
Attaching a postman collection as well in the respository for reference.
```