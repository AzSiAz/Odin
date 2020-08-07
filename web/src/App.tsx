import React from 'react'
import { useQuery } from 'react-query'
import { useRecoilState } from 'recoil'

import { meState } from './atoms'
import { getMyself } from './api'
import { Loading } from './components/loading'
import { LoginScene } from './scenes/login'
import { HomeScene } from './scenes/home'

function App() {
	const [_, setMe] = useRecoilState(meState)
	const { isError, isLoading, isSuccess } = useQuery('me', getMyself, {
		retry: false,
		onSuccess: ({ data }) => setMe(data),
	})

	return (
		<>
			{isLoading && <Loading />}
			{isError && <LoginScene />}
			{isSuccess && <HomeScene />}
		</>
	)
}

export default App
