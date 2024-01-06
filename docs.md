## Before start development You should focus on

- Testability
- Readability
- Adaptability

<br>

# The Clean Architecture

<img src='https://cdn-images-1.medium.com/max/1600/1*_5WeMzRt5aCVxXNWLlxAJw.png'>

Source: [Robert C. Martin’s blog — Clean Coder Blog](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

<I>**This rule says that source code dependencies can only point inwards. Nothing in an inner circle can know anything at all about something in an outer circle.**</I>

## What is Dependency Injection (DI) ?

DI is the idea that services should receive their dependencies when they are created. It allows us to decouple the creation of a service’s dependencies from the creation of the service itself.
