import req from "../utils/axios"

export function loginAction(data) {
    return req({
        method: 'POST',
        path: "/login",
        data
    })
}