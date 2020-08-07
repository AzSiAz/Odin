import React from 'react'
import { useRecoilState } from 'recoil'
import { meState } from '../../atoms'

export const HomeScene: React.FC = () => {
	const [me] = useRecoilState(meState)

	return <div>HomeScene: {JSON.stringify(me)}</div>
}
