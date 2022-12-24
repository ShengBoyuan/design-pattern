# Pre-Knowledge

## Introduction to OOP

4 Pillars of OOP
1. Abstraction
2. Polymorphism
3. Encapsulation
4. Inheritance

## Introduction to Patterns

3 Main Groups of Patterns
1. Creational Patterns
2. Structural Patterns
3. Behavioral Patterns

## Software Design Principles

Features of Good Design
1. Code reuse
    3 levels of reuse
2. Extensibility
    Change is the only constant thing in a programmer's life.

Design Principles
1. Encapsulate the varies
    encapsulation on a method level
    encapsulation on a class level
2. Program to an Interface, not an Implementation
3. Favor Composition Over Inheritance
    
Solid Principles
1. Single Responsibility Principle
    a class should have just one reason to change.
2. Open / Closed Principle
    classes should be open for extension but closed for modification.
3. Liskov Substitution Principle
    1. Parameter types in a method of a subclass should match or be more abstract than parameter types in the method of the superclass
    2. The return type in a method of a subclass should match or be a subtype of the return type in the method of the superclass
    3. A method in a subclass shouldn't throw types of exceptions which the base method isn't expected to throw
    4. A subclass shouldn't strengthen pre-conditions
    5. A subclass should't weaken post-conditions
    6. Invariants of a superclass must be preserved
    7. A subclass shouldn't change values of private fields of the superclass
4. Interface Segregation Principle
    Clients shouldn't be forced to depend on methods they do not use.
5. Dependency Inversion Principle
    Low-level classes depend on high-level ones, details depend on abstractions.