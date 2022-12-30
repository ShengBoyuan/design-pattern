# No.13 Chain of Responsibility (Chain of Command)
Lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process or pass to the next.

## Structure
The structure of CoR consists of 4 parts:
1. Handler
  Declares the interface for all concrete handlers. It usually contains just a single method for handling requests. Sometimes may also have another method for setting the next handler.
2. Base Handler
  Optinal class where you can put your boilerplate code that's common to all handler classes.
3. Concrete Handlers
  Contain the actual code for processing request. Upon receiving a request, each handler must decide whether to process it or pass it along the chain.
4. Client
  Compose chains just once or compose them dynamically.
![avatar](structure.png)

## When to Use
- When your program is expected to process different kinds of requests in various ways, but the exact types of requests and their sequences are unkown beforehand.
- Execute several handlers in a particular order.
- When the set of handlers and their order are supposed to change at runtime.

## How to Implement
1. Declare the handler interface and describe the signature of a method of handling requests.
2. Eliminate duplicate boilerplate code in concrete handlers, it might be worthing creating an abstract base handler class, derived from the handler interface.
3. Create concrete handler subclasses and implement their handling methods.
4. Client may either assemble chains on its own or receive pre-built chains from other objects. (Use some factory classes to build chains)

## Pros and Cons
Pros
- Control the order of request handling.
- Single Responsibility Principle.
- Open /Closed Principle.
Cons
- Some requests may end up unhandled.

## Relations with Other Patterns
- CoR is often used in conjunction with Composite.
- Handlers in CoR can be implemented as Commands.
- CoR and Decorator have very similar class structures. CoR can stop passing the request, and execute arbitrary operations independently. Decorators extend object's behavior while keeping it consistent with the base interface, and it isn't allowed to break the flow of the request.