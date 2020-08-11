import React from 'react'
import { useQuery } from 'react-query'
import { useNavigate } from 'react-router-dom'
import { useRecoilState } from 'recoil'
import { getMyself } from '../../api'
import { meState } from '../../atoms'
import { Loading } from '../../components/loading'
import { LoginScene } from '../login'

export const HomeScene: React.FC = () => {
	const [me, setMe] = useRecoilState(meState)
	const navigate = useNavigate()

	const { isError, isLoading } = useQuery('me', getMyself, {
		retry: false,
		onSuccess: ({ data }) => setMe(data),
		onError: () => navigate('login')
	})

	if (isLoading) return <Loading />
	if (isError) return <LoginScene />

	return <div>HomeScene: {JSON.stringify(me)}</div>
}
