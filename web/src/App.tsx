import React from 'react'

import { LoginScene } from './scenes/login'
import { HomeScene } from './scenes/home'
import { Route, Routes } from 'react-router-dom'

function App() {
	return (
		<Routes>
			<Route path="/" element={<HomeScene />} />
			<Route path="/login" element={<LoginScene />} />
			<Route path="/*" element={<div>Not found </div>} />
		</Routes>
	)
}

export default App
