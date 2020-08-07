import ky from 'ky'
import { MeState } from '../atoms'

type ClientResponse<T = {}> = {
    message: string,
    data: T
}

const client = ky.create({
    prefixUrl: 'http://localhost:8080/api',
    retry: 0,
    headers: {
        authorization: `Bearer ${localStorage.getItem("token") ?? ''}`
    }
})

export const getMyself= () : Promise<ClientResponse<MeState>> => client.get('user/me').json()