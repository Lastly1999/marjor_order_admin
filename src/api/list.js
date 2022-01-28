import req from "../utils/axios"

export function getOrders() {
    return req({
        method: 'GET',
        path: "/orders"
    })
}