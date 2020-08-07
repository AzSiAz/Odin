import { atom } from 'recoil'

export type MeState = {
    id: string
    username: string
    email: string
}

export const meState = atom<MeState | undefined>({
    key: 'me',
    default: undefined
})