# No.1 Factory

The structure of Factory Method consists of 4 parts:
1. Product 
  Declare the interface, what can be produced and what these products do.
2. Concrete Product
  Different implementation of product.
3. Creator
  Declare factory method, return type must match the product interface.
4. Concrete Creator
  Override the base factory method to return different types of product.
![avatar](structure.png)