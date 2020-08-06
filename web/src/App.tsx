import React from 'react'
import { useQuery } from 'react-query'

import { getMyself } from './api'

function App() {
	const { data: me, error, isError, isLoading, isSuccess } = useQuery('me', getMyself)

	return (
		<div>
			{isLoading && <div>Loading</div>}
			{isError && <div>Error: {JSON.stringify(error)}</div>}
			{isSuccess && <div>Me: {JSON.stringify(me)}</div>}
		</div>
	)
}

export default App
