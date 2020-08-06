import ky from 'ky'

const client = ky.create({
    prefixUrl: 'http://localhost:8080/api',
    retry: 0,
    headers: {
        authorization: `Bearer ${localStorage.getItem("token") ?? ''}`
    }
})

export const getMyself = () => client.get('user/me').json()