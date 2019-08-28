# Check Domain
Play around with Altravia srl Test API for fun but also for profit

## The Problem
We want to have a small, fast program that:
- Checks for availability of a italian internet domain.
- If the domain is available, purchase it.

We want to run this program as a scheduled task, for example as a cronjob on a Linux server,
where it will be run every hour or so.

## The Constraints
- The program should be small, lightweight and require minimal dependencies to run.
- We should use Altravia's API. Altravia is a web registrar whom has given us permission to use their api.
Credentials are available, and also [documentation](https://test.api.altravia.com/docs/).
- We should run it responsibly, without spamming Altravia srl's api.
- The application should be written in Go, with TDD approach.
- Deadline for a working product is 10th of September 2019.

## Some questions.
- Access to Altravia srl's API is available, but what about funds? how can we add some to our account/user?
...