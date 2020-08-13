import React from 'react'
import { Route, Routes } from 'react-router-dom'

import { LoginScene } from './scenes/login'
import { SignUpScene } from './scenes/signup'
import { HomeScene } from './scenes/home'
import { routes } from './constants'

function App() {
	return (
		<Routes>
			<Route path={routes.home} element={<HomeScene />} />
			<Route path={routes.login} element={<LoginScene />} />
			<Route path={routes.signUp} element={<SignUpScene />} />
			<Route path={routes.notFound} element={<div>Not found </div>} />
		</Routes>
	)
}

export default App
