redis-cli --scan --pattern "orders||*" | xargs redis-cli del
redis-cli --scan --pattern "users||*" | xargs redis-cli del
redis-cli --scan --pattern "cart||*" | xargs redis-cli del
