import axios from "axios"

const API_URL = "http://www.fjranling.com:3600"


const req = (config) => {
    return new Promise((resolve, reject) => {
        axios({
            url: API_URL + config.path,
            method: config.method,
            data: config.data
        }).then(res => {
            resolve(res.data)
        }).catch(err => {
            reject(err)
        })
    })
}

export default req