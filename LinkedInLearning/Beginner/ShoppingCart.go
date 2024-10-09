// Write your answer here, and then test your code.
package main

type cartItems struct{
    name string
    price float64
    quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []cartItems) float64 {
    var result float64 = 0
    for _, item := range cart {
        result += (item.price * float64(item.quantity))
    }
    return result
}

