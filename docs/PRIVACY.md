# Privacy

Under the current specification, the server knows:

- Your username.
- Your IP address, since you can't make HTTP requests to the server without it knowing, unless you take steps to mask it with a VPN or TOR.
- The username of anyone who shares a task with you.
- The times at which tasks were created.
- The IDs of the tasks you access. It does not have access to the contents of these tasks, but knowledge of the IDs can be used to determine the usernames of the people you collaborate with.
- The times at which you access tasks. This could be used to determine your habits or timezone if you consistently access tasks during certain times of the day (your work hours), or consistently never access them during other times (when you're sleeping).
