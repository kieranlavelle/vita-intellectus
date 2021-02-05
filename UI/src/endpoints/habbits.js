import { API } from './http-config'

export async function getHabbits() {
    return API.get("habbits")
        .then(res => res.data)
        .catch(err => console.log(err))
}