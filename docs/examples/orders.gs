#!/usr/bin/env gs

fmt = import("fmt")

orders = [
  {customer: "Ada", total: 120, paid: true},
  {customer: "Linus", total: 75, paid: false},
  {customer: "Grace", total: 210, paid: true},
  {customer: "Ken", total: 95, paid: true},
]

paid_total = orders
  .filter(order => order.paid)
  .map(order => order.total)
  .sum()

vip_customers = orders
  .filter(order => order.total >= 100)
  .map(order => order.customer)

fmt.printf("paid total: %v\n", paid_total)
fmt.printf("vip customers: %v\n", vip_customers)